# Picoclaw 多模态前端改进进度

## 📅 更新日期：2026-04-09

---

## 🎯 当前任务：修复 WebSocket 连接问题

### 问题描述

前端多模态功能无法连接到 WebSocket，出现连接失败错误。

### 错误现象

```
WebSocket connection to 'wss://pico.mynameqx.top:18800/pico/ws' failed
```

### 根因分析

1. **问题1**：后端 `/api/pico/token` 返回的 `ws_url` 包含端口 `18800`：
   ```json
   {"enabled":true,"token":"pico-xxx","ws_url":"wss://pico.mynameqx.top:18800/pico/ws"}
   ```

2. **问题2**：前端直接使用了后端返回的 `ws_url`，导致通过 nginx 代理时端口不匹配

3. **问题3**：nginx 将 WebSocket 请求代理到 `http://192.168.2.18:18800`（LAN IP），而 picoclaw-launcher 只监听 `127.0.0.1:18800`

### 已修复

✅ **修复 1**：修改前端代码，从 `window.location` 构造 WebSocket URL，不再使用后端返回的 `ws_url`

**修改文件**：`web/frontend/src/features/chat/controller.ts`

**修改内容**：
```typescript
// 修改前（使用后端返回的 ws_url）
const wsProtocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
const wsUrl = wsProtocol + '//' + window.location.host + '/pico/ws'

// 原来的代码使用了后端返回的 ws_url，包含端口 18800
```

✅ **修复 2**：前端现在正确构造 URL：
```typescript
const wsProtocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
const wsUrl = wsProtocol + '//' + window.location.host + '/pico/ws'
// 不再包含端口号
```

### 待解决

⏳ **问题 A**：nginx upstream 配置使用 hostname `pico` 解析到 LAN IP `192.168.2.18:18800`

- 需要修改 nginx 配置，将 `proxy_pass http://pico;` 改为 `proxy_pass http://127.0.0.1:18800;`

⏳ **问题 B**：WebSocket 认证问题

- 认证 Cookie 名称：`picoclaw_launcher_auth`
- 需要正确的 Token 格式
- Token 验证逻辑：`picoComposedToken()` 函数

### 相关文件位置

| 文件 | 说明 |
|------|------|
| `web/frontend/src/features/chat/controller.ts` | 前端 WebSocket 连接代码（已修改） |
| `web/backend/api/pico.go` | 后端 WebSocket 代理和 Token 验证 |
| `web/backend/api/gateway.go` | Gateway Token 组合逻辑 |
| `web/backend/api/gateway_host.go` | WebSocket URL 构建逻辑 |

### 测试命令

```bash
# 1. 获取 Token
curl -s -b "picoclaw_launcher_auth=<cookie>" http://127.0.0.1:18800/api/pico/token

# 2. 测试 WebSocket（带 Token）
curl -v --no-buffer \
  -H "Connection: Upgrade" \
  -H "Upgrade: websocket" \
  -H "Sec-WebSocket-Version: 13" \
  -H "Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==" \
  -H "Sec-WebSocket-Protocol: token.<token>" \
  -b "picoclaw_launcher_auth=<cookie>" \
  http://127.0.0.1:18800/pico/ws
```

---

## 📋 相关需求文档

| 文档 | 位置 | 说明 |
|------|------|------|
| 需求文档 | 待补充 | - |
| 设计文档 | 待补充 | - |
| API 文档 | `web/backend/api/` | 后端 API 实现 |

---

## 🔜 下一步

1. [ ] 修改 nginx 配置，将 WebSocket upstream 改为 127.0.0.1:18800
2. [ ] 解决 WebSocket Token 验证问题
3. [ ] 端到端测试 WebSocket 连接
4. [ ] 验证多模态功能正常工作

---

## 📝 备注

- picoclaw-launcher 监听 `127.0.0.1:18800`
- picoclaw-gateway 监听 `127.0.0.1:18790`（仅本地）
- WebSocket 通过 launcher 代理到 gateway
- 认证 Cookie：`picoclaw_launcher_auth`（HttpOnly）

---

*太枢记录 | 2026-04-09*
