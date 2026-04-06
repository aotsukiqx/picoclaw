# PicoClaw 安卓界面与后台交互技术文档

## 一、架构概述

PicoClaw 的安卓界面并非独立的应用商店应用，而是基于 Web 技术（React）打包的渐进式 Web 应用（PWA），通过系统浏览器访问Launcher后端服务。

### 1.1 系统架构图

```
┌─────────────────────────────────────────────────────────────────────────┐
│                           安卓设备 / 浏览器                               │
│  ┌─────────────┐    ┌──────────────────────────────────────────────┐   │
│  │  Android    │    │           React Web UI (PWA)                 │   │
│  │  Chrome     │    │  ┌─────────────┐  ┌────────────────────────┐  │   │
│  │  Browser    │    │  │  Chat UI   │  │   Dashboard Config    │  │   │
│  └──────┬──────┘    │  └──────┬──────┘  └──────────┬─────────────┘  │   │
│         │ WebSocket │         │                      │ HTTP REST     │   │
│         │           │         │                      │               │   │
│         │           │  ┌──────┴──────────────────────┴──────────┐   │   │
│         │           │  │      websocket.ts + controller.ts        │   │   │
│         │           │  └──────────────────┬───────────────────────┘   │   │
│         │           └─────────────────────┼───────────────────────────┘   │
└─────────┼────────────────────────────────┼─────────────────────────────┘
          │                                │
          │      ┌─────────────────────────┴─────────────────────────┐
          │      │              Web Backend (Go)                       │
          │      │              Port: 18800                           │
          │      │  ┌─────────────────────────────────────────────┐    │
          │      │  │  API Router (router.go)                    │    │
          │      │  │  ├── GET  /api/pico/token                  │    │
          │      │  │  ├── POST /api/pico/token                  │    │
          │      │  │  ├── POST /api/pico/setup                  │    │
          │      │  │  ├── GET  /pico/ws  (WebSocket代理)         │    │
          │      │  │  ├── GET  /api/version                     │    │
          │      │  │  ├── GET  /api/models                       │    │
          │      │  │  └── ...                                    │    │
          │      │  └──────────────────┬────────────────────────────┘    │
          │      │                   │                                 │
          │      │  ┌───────────────┴────────────────────────────┐    │
          │      │  │  Pico WebSocket Proxy (pico.go)            │    │
          │      │  │  将 WebSocket 连接反向代理到 Gateway         │    │
          │      │  └──────────────────┬───────────────────────────┘    │
          │      └─────────────────────┼────────────────────────────────┘
          │                            │
          │      ┌────────────────────┴────────────────────┐
          │      │              picoclaw Gateway            │
          │      │         (Pico Channel - WebSocket)         │
          │      │  ┌────────────────────────────────────┐  │
          │      │  │  pkg/channels/pico/pico.go         │  │
          │      │  │  ├── handleWebSocket()             │  │
          │      │  │  ├── readLoop()                    │  │
          │      │  │  ├── handleMessage()               │  │
          │      │  │  └── broadcastToSession()          │  │
          │      │  └──────────────────┬─────────────────────┘  │
          │      └─────────────────────┼──────────────────────────┘
                                      │
                              ┌───────┴───────┐
                              │  MessageBus    │
                              │  (消息总线)    │
                              └───────┬───────┘
                                      │
                              ┌───────┴───────┐
                              │  Agent Loop   │
                              │  (AI 推理)    │
                              └───────────────┘
```

### 1.2 通信方式总结

| 通信类型 | 协议 | 端口 | 用途 |
|---------|------|------|------|
| Web UI ↔ Backend | HTTP REST | 18800 | 配置管理、模型列表、认证 |
| Web UI ↔ Gateway | WebSocket | 18800 (代理) | 实时聊天、输入状态同步 |
| Gateway ↔ Agent | Internal | In-Process | 消息处理、AI 推理 |

---

## 二、Pico Protocol 协议详解

### 2.1 消息格式

所有消息采用 JSON 格式，定义在 `pkg/channels/pico/protocol.go`：

```go
type PicoMessage struct {
    Type      string         `json:"type"`       // 消息类型
    ID        string         `json:"id,omitempty"`        // 消息 ID
    SessionID string         `json:"session_id,omitempty"` // 会话 ID
    Timestamp int64          `json:"timestamp,omitempty"`  // 时间戳
    Payload   map[string]any `json:"payload,omitempty"`   // 载荷
}
```

### 2.2 消息类型

#### 客户端 → 服务端

| 消息类型 | 说明 | Payload |
|---------|------|---------|
| `message.send` | 发送文本消息 | `content: string` |
| `media.send` | 发送媒体（Base64 图片） | `content: string, media: string[]` |
| `ping` | 心跳检测 | - |

#### 服务端 → 客户端

| 消息类型 | 说明 | Payload |
|---------|------|---------|
| `message.create` | 新消息创建 | `content: string, message_id: string` |
| `message.update` | 消息更新（流式输出） | `content: string, message_id: string` |
| `typing.start` | 开始输入状态 | - |
| `typing.stop` | 停止输入状态 | - |
| `pong` | 心跳响应 | - |
| `error` | 错误信息 | `code: string, message: string` |

### 2.3 消息流示例

**用户发送消息：**
```json
{
  "type": "message.send",
  "id": "msg-1-1234567890",
  "session_id": "abc123",
  "payload": {
    "content": "你好，请介绍一下 PicoClaw"
  }
}
```

**AI 逐步响应（流式）：**
```json
// 开始
{ "type": "typing.start", "session_id": "abc123" }

// 响应片段 1
{ "type": "message.create", "id": "msg-2-1234567891", "session_id": "abc123", "payload": { "content": "Pico", "message_id": "msg-2" } }

// 响应片段 2
{ "type": "message.update", "session_id": "abc123", "payload": { "content": "PicoClaw 是", "message_id": "msg-2" } }

// 完成
{ "type": "typing.stop", "session_id": "abc123" }
```

---

## 三、连接建立流程

### 3.1 初始化流程（Web UI）

```
┌─────────────┐                              ┌─────────────┐
│   Browser   │                              │   Backend   │
│  (React UI) │                              │  (Go :18800)│
└──────┬──────┘                              └──────┬──────┘
       │                                            │
       │  1. GET /api/pico/token                   │
       │───────────────────────────────────────────►│
       │                                            │
       │  { token: "xxx", ws_url: "...", enabled }│
       │◄───────────────────────────────────────────│
       │                                            │
       │  2. WebSocket /pico/ws?session_id=xxx      │
       │     (SubProtocol: token.xxx)               │
       │───────────────────────────────────────────►│
       │                                            │
       │  3. [反向代理到 Gateway Pico Channel]      │
       │                                            │
       │◄══════════ WebSocket 双向通信 ════════════►│
```

### 3.2 关键代码位置

**前端连接逻辑：** `web/frontend/src/features/chat/controller.ts`

```typescript
// 获取 Token
const { token, ws_url } = await getPicoToken()

// 建立 WebSocket 连接
const socket = new WebSocket(url, [`token.${token}`])

// 发送消息
socket.send(JSON.stringify({
  type: "message.send",
  id,
  payload: { content, media: [...] }
}))
```

**后端 Token 获取：** `web/backend/api/pico.go`

```go
// GET /api/pico/token
func (h *Handler) handleGetPicoToken(w http.ResponseWriter, r *http.Request) {
    cfg, _ := config.LoadConfig(h.configPath)
    json.NewEncoder(w).Encode(map[string]any{
        "token":   cfg.Channels.Pico.Token.String(),
        "ws_url":  h.buildWsURL(r),
        "enabled": cfg.Channels.Pico.Enabled,
    })
}
```

### 3.3 WebSocket 代理机制

Launcher 后端将 `/pico/ws` 请求反向代理到 Gateway：

```
web/backend/api/pico.go:
  handleWebSocketProxy() 
    → createWsProxy()
      → httputil.ReverseProxy
        → Gateway Pico Channel (pkg/channels/pico/pico.go)
```

---

## 四、认证机制

### 4.1 Token 认证方式

Pico Channel 支持三种 Token 认证方式（按优先级）：

1. **Authorization Header**
   ```
   Authorization: Bearer <token>
   ```

2. **WebSocket SubProtocol**
   ```
   Sec-WebSocket-Protocol: token.<token_value>
   ```

3. **Query Parameter**（仅当 `AllowTokenQuery=true` 时）
   ```
   /pico/ws?token=<token>
   ```

### 4.2 源码实现

**服务端认证：** `pkg/channels/pico/pico.go`

```go
func (c *PicoChannel) authenticate(r *http.Request) bool {
    token := c.config.Token.String()
    
    // 1. Authorization Header
    if auth := r.Header.Get("Authorization"); strings.HasPrefix(auth, "Bearer ") {
        if strings.TrimPrefix(auth, "Bearer ") == token {
            return true
        }
    }
    
    // 2. Sec-WebSocket-Protocol
    if c.matchedSubprotocol(r) != "" {
        return true
    }
    
    // 3. Query Parameter
    if c.config.AllowTokenQuery {
        if r.URL.Query().Get("token") == token {
            return true
        }
    }
    
    return false
}
```

---

## 五、前端消息处理

### 5.1 消息处理映射

**文件：** `web/frontend/src/features/chat/protocol.ts`

```typescript
function handlePicoMessage(message: PicoMessage, expectedSessionId: string) {
    switch (message.type) {
        case "message.create":    // 新消息 → 添加到聊天列表
            updateChatStore({ messages: [..., { role: "assistant", content }] })
            break
        case "message.update":    // 消息更新 → 替换对应消息内容（流式）
            updateChatStore({ messages: messages.map(m => 
                m.id === messageId ? { ...m, content } : m
            )})
            break
        case "typing.start":     // 对方正在输入
            updateChatStore({ isTyping: true })
            break
        case "typing.stop":     // 停止输入
            updateChatStore({ isTyping: false })
            break
        case "error":            // 错误处理
            toast.error(errorMessage)
            break
    }
}
```

### 5.2 连接状态管理

```typescript
// web/frontend/src/features/chat/controller.ts

type ConnectionState = "disconnected" | "connecting" | "connected" | "error"

// 自动重连机制
function scheduleReconnect(generation, sessionId) {
    const delay = Math.min(1000 * 2 ** reconnectAttempts, 5000)
    reconnectTimer = setTimeout(() => connectChat(), delay)
}
```

---

## 六、Gateway Pico Channel 实现

### 6.1 核心组件

**文件：** `pkg/channels/pico/pico.go`

```go
type PicoChannel struct {
    *channels.BaseChannel           // 继承基础通道能力
    config       config.PicoConfig  // Pico 特有配置
    connections  map[string]*picoConn  // 连接 ID → 连接
    sessionConnections map[string]map[string]*picoConn  // sessionID → 连接集合
}
```

### 6.2 连接管理

- **最大连接数：** 默认 100，可配置 `MaxConnections`
- **心跳检测：** Ping/Pong 机制，默认 30 秒间隔
- **读超时：** 默认 60 秒

### 6.3 消息广播

```go
// 广播消息到同一会话的所有连接
func (c *PicoChannel) broadcastToSession(chatID string, msg PicoMessage) error {
    sessionID := strings.TrimPrefix(chatID, "pico:")
    for _, pc := range c.sessionConnectionsSnapshot(sessionID) {
        pc.writeJSON(msg)  // 线程安全的写操作
    }
}
```

---

## 七、API 端点总结

### 7.1 Pico Channel API

| 方法 | 路径 | 说明 |
|-----|------|------|
| GET | `/api/pico/token` | 获取当前 WebSocket Token 和连接 URL |
| POST | `/api/pico/token` | 重新生成 Token |
| POST | `/api/pico/setup` | 自动配置 Pico Channel（首次设置） |
| GET | `/pico/ws` | WebSocket 连接端点（反向代理到 Gateway） |

### 7.2 其他管理 API

| 方法 | 路径 | 说明 |
|-----|------|------|
| GET | `/api/version` | 获取版本信息 |
| GET | `/api/models` | 获取模型列表 |
| GET | `/api/status` | 获取运行状态 |
| POST | `/api/gateway/start` | 启动 Gateway |
| POST | `/api/gateway/stop` | 停止 Gateway |

---

## 八、数据流总览

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                            完整数据流                                       │
└─────────────────────────────────────────────────────────────────────────────┘

[用户输入] 
    │
    ▼
web/frontend/src/features/chat/controller.ts
    │ sendChatMessage()
    │
    ▼ JSON: { type: "message.send", payload: { content, media } }
    │
WebSocket ──────────────────────────────────────────────────────►
    │                                                             │
pkg/channels/pico/pico.go                                          │
    │ handleWebSocket() → readLoop()                              │
    ▼                                                             │
handleMessageSend()                                               │
    │ 构建 InboundMessage                                          │
    ▼                                                             │
MessageBus.PublishInbound()                                       │
    │                                                              │
    ▼                                                              │
AgentLoop (AI 推理)                                               │
    │                                                              │
    ▼ 流式输出                                                     │
MessageBus.PublishOutbound()                                       │
    │                                                              │
    ▼ broadcastToSession()                                         │
    │                                                              │
WebSocket ◄─────────────────────────────────────────────────────
    │                                                              │
web/frontend/src/features/chat/protocol.ts                         │
    │ handlePicoMessage()                                          │
    ▼                                                              │
React State Update (setState)                                     │
    │                                                              │
    ▼                                                              │
UI Re-render (消息展示)                                           │
```

---

## 九、关键文件索引

| 功能 | 文件路径 |
|-----|---------|
| Pico Protocol 定义 | `pkg/channels/pico/protocol.go` |
| Pico Channel 服务端 | `pkg/channels/pico/pico.go` |
| Pico Client (反向代理) | `pkg/channels/pico/client.go` |
| Web Backend API 路由 | `web/backend/api/router.go` |
| Pico API Handler | `web/backend/api/pico.go` |
| 前端 WebSocket 控制器 | `web/frontend/src/features/chat/controller.ts` |
| 前端消息处理 | `web/frontend/src/features/chat/protocol.ts` |
| 前端 Pico API 客户端 | `web/frontend/src/api/pico.ts` |
| 前端 HTTP 工具 | `web/frontend/src/api/http.ts` |

---

## 十、安卓特殊说明

1. **安卓应用下载：** PicoClaw 安卓应用非源码形式存在，需从 [picoclaw.io](https://picoclaw.io/download/) 下载 APK 安装

2. **运行模式：** 安卓上的 PicoClaw 通过嵌入的 WebView 加载 Launcher 后端

3. **网络要求：** 
   - 安卓设备与运行 Launcher 的主机需在同一网络
   - 或通过公网暴露 Launcher（需配置 `-public` 参数）

4. **APK 构成：** 安卓 APK 本质是一个打包的 WebView + Go 后端，可离线运行
