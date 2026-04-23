# PicoClaw CI/CD 开发流程

## 概述

本文档定义 picoclaw (Go 后端) + picoclaw_fui (Flutter 前端) 的开发测试流程。

## 环境架构

```
┌─────────────────────────────────────────────────────────────┐
│                    macOS 开发主机                            │
│  ┌──────────────────┐    ┌──────────────────────────────┐  │
│  │  picoclaw Go     │    │  picoclaw_fui Flutter      │  │
│  │  (Darwin/arm64)  │    │  (Android APK Build)       │  │
│  │  后端服务测试     │    │  前端 UI 测试               │  │
│  └────────┬─────────┘    └──────────────┬───────────────┘  │
│           │                               │                  │
│           │  WebSocket :18800             │                  │
│           └───────────────┬───────────────┘                  │
│                           │                                  │
└───────────────────────────┼──────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                    Android 设备                               │
│  ┌──────────────────────────────────────────────────────┐  │
│  │  picoclaw_fui (APK) + picoclaw Go (Android/arm64)  │  │
│  └──────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

## 开发流程

### 1. 后端开发 (picoclaw)

```bash
# 1. 克隆代码
cd projects/picoclaw

# 2. 开发新功能/修复 bug
# ... 编辑代码 ...

# 3. 运行单元测试
make test

# 4. 本地启动后端服务
make run

# 5. 测试 API
curl http://localhost:18800/api/version

# 6. 提交代码
git add .
git commit -m "feat: 新功能描述"
git push
```

### 2. 前端开发 (picoclaw_fui)

```bash
# 1. 克隆代码
cd projects/picoclaw_fui

# 2. 开发新功能/修复 bug
# ... 编辑代码 ...

# 3. 运行 Flutter 测试
flutter test

# 4. 构建 Android APK
flutter build apk --debug

# 5. 本地真机/模拟器测试
flutter run -d <device_id>

# 6. 提交代码
git add .
git commit -m "feat: 新功能描述"
git push
```

### 3. 集成测试 (前后端联调)

```bash
# 1. 启动后端服务 (macOS)
cd projects/picoclaw
./build/picoclaw-darwin-arm64

# 2. 确认后端运行
curl http://localhost:18800/api/version

# 3. 启动 Flutter 前端连接本地后端
cd projects/picoclaw_fui
flutter run -d <device_id>
# 修改 chat_page.dart 中的 _gatewayHost 为 macOS IP

# 4. 测试语音功能
# - 录音测试
# - TTS 播放测试
# - 端到端对话测试
```

## CI/CD 流程

### GitHub Actions 工作流

#### 1. 后端 CI (`picoclaw/.github/workflows/ci.yml`)

```yaml
name: Go CI

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main, develop]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'

      - name: Run tests
        run: make test

      - name: Build
        run: make build

  android-build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'

      - name: Cross-compile for Android
        run: |
          GOOS=linux GOARCH=arm64 make build
          # 或使用 android build 脚本

      - name: Upload binary
        uses: actions/upload-artifact@v4
        with:
          name: picoclaw-android
          path: build/picoclaw-linux-arm64
```

#### 2. 前端 CI (`picoclaw_fui/.github/workflows/ci.yml`)

```yaml
name: Flutter CI

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main, develop]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Flutter
        uses: subosito/flutter-action@v2
        with:
          flutter-version: '3.24.0'
          channel: 'stable'

      - name: Install dependencies
        run: flutter pub get

      - name: Run tests
        run: flutter test

      - name: Analyze
        run: flutter analyze

  android-build:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Flutter
        uses: subosito/flutter-action@v2
        with:
          flutter-version: '3.24.0'
          channel: 'stable'

      - name: Install dependencies
        run: flutter pub get

      - name: Build APK
        run: flutter build apk --debug

      - name: Upload APK
        uses: actions/upload-artifact@v4
        with:
          name: debug-apk
          path: build/app/outputs/apk/debug/app-debug.apk
```

## 本地测试检查清单

### 后端测试

- [ ] `make test` 通过
- [ ] `curl http://localhost:18800/api/version` 返回正确
- [ ] WebSocket 连接正常
- [ ] ASR 功能正常（需要配置 API Key）
- [ ] TTS 功能正常（需要配置 API Key）

### 前端测试

- [ ] `flutter test` 通过
- [ ] `flutter analyze` 无警告
- [ ] `flutter build apk --debug` 构建成功
- [ ] 安装到真机/模拟器
- [ ] 录音按钮显示正常
- [ ] 录音功能正常
- [ ] TTS 播放正常

### 集成测试

- [ ] 前端连接后端成功
- [ ] 文本对话正常
- [ ] 语音对话正常
- [ ] 端到端流程无断连

## Android 真机测试

### 连接同一网络

```bash
# 1. 查看 macOS IP
ifconfig | grep "192.168" | head -1

# 2. 修改 chat_page.dart 中的 _gatewayHost
static const _gatewayHost = '192.168.x.x';  # macOS IP

# 3. 确保后端开启 public mode 或修改 CORS
./picoclaw --public

# 4. 安装 APK 到手机
adb install -r build/app/outputs/apk/debug/app-debug.apk

# 5. 查看日志
adb logcat | grep picoclaw
```

### 常见问题

| 问题 | 解决方案 |
|------|----------|
| 连接失败 | 检查防火墙/同一网络 |
| 录音无响应 | 检查 RECORD_AUDIO 权限 |
| TTS 无声音 | 检查 Audio Focus |

## 发布流程

### 1. 版本 Tag

```bash
# 后端
cd projects/picoclaw
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0

# 前端
cd projects/picoclaw_fui
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

### 2. 构建发布 APK

```bash
# 使用 GitHub Actions 构建
# 或本地构建 release 版本
cd projects/picoclaw_fui
flutter build apk --release

# 重命名 APK
mv build/app/outputs/apk/release/app-release.apk picoclaw-{version}-release.apk
```

### 3. 测试发布版本

```bash
# 安装到设备
adb install -r picoclaw-{version}-release.apk

# 验证功能
# ... 测试 ...
```

## 自动化测试建议

### 单元测试

- Go: `make test` (已有)
- Flutter: `flutter test` (已有)

### 集成测试

- 使用 Firebase Test Lab
- 或本地模拟器脚本

### E2E 测试

- 使用 Flutter Driver / Integration Test
- 或手动测试清单

## 参考资料

- [Flutter CI/CD](https://docs.flutter.dev/ci/cd)
- [Go Testing](https://pkg.go.dev/testing)
- [GitHub Actions](https://docs.github.com/actions)
