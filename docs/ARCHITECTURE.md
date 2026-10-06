# Architecture Overview

Aurora Orchestrator aims to be a small, extensible framework that demonstrates best practices for agent-oriented systems in Go.

Core concepts
- Agent: a persisted entity that represents an autonomous actor. Contains metadata and a JSON blob for capabilities.
- Manager: the in-memory coordinator that runs agent tasks (and can cancel them). It acts as the single source of truth for active runs.
- Store: persistence layer via Gorm. Models are migrated at startup and are abstracted into internal/store.
- Handlers & Router: Gin-based HTTP surface for interacting with agents.

Design goals
- Separation of concerns: persistence, orchestration, and HTTP surface are decoupled
- Testability: Manager and store are easy to test with in-memory SQLite
- Extensibility: swap mock execution with real AI tools & LLMs; add distributed queueing, authorization, or observability

Possible extensions
- Pluggable backends for execution (Kubernetes jobs, queues)
- Websocket stream of execution logs
- Task graphs and tool-use orchestration
- Secure credentials store for API keys

