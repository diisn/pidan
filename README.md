<div align="center">

# π pidan

**用 Go 写的终端 AI 编码助手**

[![Go](https://img.shields.io/badge/Go-1.27-00ADD8?logo=go&logoColor=white)](go.mod)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/diisn/pidan/issues)

</div>

## 特性

- **全屏 TUI**，并提供基于行的 REPL 回退模式（`--no-tui`）
- **Headless 模式**，适合管道与脚本 —— 输出纯文本或 `stream-json` 事件流
- **内置工具集** —— 文件读写/编辑、bash、搜索、网页抓取、待办、记忆、目标
- **会话管理** —— 恢复任意历史会话（`-r`）或继续最近一次（`-c`）
- **技能与提示词包** —— 通过 `pidan install/list/uninstall/update` 管理
- **Provider 无关** —— 兼容 OpenAI 的网关（默认 OpenRouter，本地模型可用 Ollama）；支持自定义 `--model`、`--base-url`、`--protocol`
- **自动更新** —— `pidan update` 拉取最新发布版本

## 安装

从源码构建（需要 Go 1.27+）：

```sh
go install github.com/diisn/pidan/cmd/pidan@latest
# 或克隆后本地构建：
go build -o pidan ./cmd/pidan
```

## 快速开始

先为你的 provider 设置 API key，然后运行：

```sh
export OPENROUTER_API_KEY=sk-or-…
pidan                    # PowerShell 里用 $env:OPENROUTER_API_KEY
```

不进入 TUI，直接跑一次性任务：

```sh
pidan -p "读一下 README 并总结"
```

## 用法

| 命令 | 说明 |
| --- | --- |
| `pidan` | 交互式 TUI（TTY 下） |
| `pidan -p "<prompt>"` | Headless：执行一次并输出最终答案 |
| `pidan -p "<prompt>" -o stream-json` | Headless：输出行分隔的 JSON 事件 |
| `pidan -r <session-id>` / `pidan -c` | 恢复会话 / 继续最近一次会话 |
| `pidan -m <model>` | 指定模型（如 `deepseek-chat`） |
| `pidan --provider <name>` | 使用某 provider 的默认配置（`--help` 可查看列表） |
| `pidan install \| list \| uninstall \| update <pkg>` | 管理技能/提示词包 |
| `pidan --version` | 打印构建信息 |
| `pidan session export <id> \| list` | 管理已保存的会话 |

模型、API key 和端点由 `--model`、`--provider` 以及对应的
`<PROVIDER>_API_KEY` 环境变量解析确定。运行 `pidan --help` 查看完整参数列表。

## 配置

可选配置文件 `~/.config/pidan/config.toml`（支持 `$XDG_CONFIG_HOME`）。
优先级：**命令行参数 > 配置文件 > 内置默认值**。敏感信息可以放在
`$PIDAN_HOME/.credentials.yaml` 中，通过名称引用，不必写进配置文件。

```toml
model = "openrouter/free"
provider = "openrouter"
thinking_level = "medium"

[memory]
enabled = true
```

## 项目结构

```
cmd/pidan              CLI 入口 —— 参数解析、配置叠加、模式分发
internal/cli/tui       全屏 TUI
internal/cli/repl      基于行的 REPL
internal/cli/headless  headless 会话与子代理 RPC
internal/agenttool     内置工具
internal/provider      provider 解析与通信驱动的实现
internal/runtime       运行循环、工具装配、提示词
internal/dream         记忆整合（"dream" 流程）
internal/webhook       GitHub ready-for-review webhook
```

## 开发

```sh
go test ./...
```

## 参与贡献

发现 Bug 或想加功能？欢迎[提 issue](https://github.com/diisn/pidan/issues)或提交 PR。

## License

待定 —— 仓库目前还没有 LICENSE 文件。复用前请先与维护者确认。