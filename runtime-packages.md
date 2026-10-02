# pidan 运行时包清单

刨去与运行无关的内容（电子书 `book/`、文档站 `docs/`、根目录说明性文件、CI/发布配置等）后，
**真正参与编译与运行**的目录如下。

## 运行必需（Go 代码）

```
pidan/
├── agent/                  # 公开 SDK 包（可被外部 import）
├── cmd/pidan/               # 程序入口 main（唯一二进制产物）
├── internal/
│   ├── agentcore/          # 核心类型契约：消息/内容/事件/工具/钩子
│   ├── agenttool/          # 内置工具：read/write/edit/grep/bash/websearch 等
│   ├── builtinskills/      # 内置技能（go:embed 内嵌 skills/**）
│   ├── cli/                # 命令行层
│   │   ├── repl/           #   交互式 REPL
│   │   ├── tui/            #   全屏 TUI
│   │   ├── headless/       #   无头脚本模式
│   │   ├── run/            #   无头运行装配 + hooks 驱动
│   │   ├── config/         #   config.toml 加载
│   │   ├── prompts/        #   提示词模板注册
│   │   ├── pkgcmd/         #   pidan install/list/uninstall/update
│   │   ├── sessioncmd/     #   pidan session export/list
│   │   ├── status/         #   /status 实现
│   │   ├── goal/           #   goal 子命令
│   │   ├── memstatus/      #   记忆状态
│   │   ├── btw/            #   btw 相关
│   │   └── ui/ testutil/   #   UI 渲染工具 / 测试辅助
│   ├── clipboard/          # 剪贴板（TUI 用）
│   ├── compaction/         # 上下文自动压缩
│   ├── dream/              # 记忆整合（dream 模式）
│   ├── hooks/              # 生命周期 hooks 引擎
│   ├── jsonrpc/            # 子 Agent JSON-RPC 传输
│   ├── memory/             # 持久化记忆（SQLite）
│   ├── pihost/             # 插件 extension host（go:embed pihost.mjs）
│   ├── pkgmgr/             # 包管理器（extension/skill/prompt/theme 分发）
│   ├── plugin/             # 外部插件机制
│   ├── provider/           # 模型 Provider 注册表 + OpenAI/Anthropic 协议
│   ├── remotecontrol/      # 手机遥控（go:embed web/ SPA）
│   ├── runtime/            # 核心：两层 Agent 循环/子 Agent/telemetry/slash 命令
│   ├── selfupdate/         # 二进制自更新
│   ├── session/            # 会话持久化/导出
│   ├── trust/              # 项目信任
│   └── webhook/            # GitHub PR review webhook
├── go.mod / go.sum         # 模块依赖
```

## 随二进制内嵌的非 Go 资源（属于运行时，不能删）

| 路径 | 用途 | 嵌入点 |
|---|---|---|
| `internal/builtinskills/skills/**` | 内置技能（SKILL.md + 脚本） | `//go:embed all:skills` |
| `internal/pihost/pihost.mjs` | 插件 extension host JS | `//go:embed pihost.mjs` |
| `internal/remotecontrol/web/` | 遥控 SPA（app.js/index.html） | `//go:embed web` |

## 可选保留（代码，但不进 pidan 二进制）

- `examples/sdk/01-07/` —— SDK 用法示例（`go run` 可跑，用于验证 agent 包）
- `examples/hooks/` —— hook 脚本示例

## 刨去的与运行无关内容

- `book/` —— 配套电子书书稿（章节 md、图、tex、pandoc lua 过滤器、build_pdf.sh）
- `docs/` —— 文档站（HTML、issue 归档、sandboxing.md 等）
- 根目录：`README.md`、`CHANGELOG.md`、`LICENSE`、`config.toml.example`、`install.sh`
- 配置：`.github/workflows/`（CI/Release）、`.goreleaser.yaml`
- 工具产物：`.codegraph/`

> 验证方式：`go list ./...` 列出的 45 个 Go 包即上表运行时包（外加 examples/sdk）；
> `go build ./cmd/pidan` 不依赖任何被刨去的目录。
