# Aurora Orchestrator

Aurora Orchestrator is a modular Go-based framework for building, persisting, and orchestrating autonomous AI agents. It uses Gin for HTTP APIs and Gorm for persistence, and is designed to be a clear starting point for production-grade agent systems.

Key features:
- Agent lifecycle management (create, read, execute)
- Persistent storage via Gorm (SQLite by default)
- HTTP API with Gin and clean routing/handler separation
- Lightweight manager for coordinating agent runs (mock execution)
- Unit tests and easy extension points for AI integrations

Architecture
------------

The repository follows a modular layout for maintainability and testing:

- cmd/server: Application entrypoint
- internal/
  - agent: Agent domain logic and manager
  - store: Database initialization and migrations
  - handlers: HTTP handlers (Gin)
  - router: Router wiring and middleware
  - config: Configuration utilities
- docs: Design & architecture docs
- tests: Unit tests colocated with packages

Sequence (high-level):
1. cmd/server/main.go initializes configuration and DB
2. Gorm migrates models (Agent)
3. AgentManager is created and passed to HTTP handlers
4. Clients interact via RESTful endpoints to create, list, and execute agents

Installation
------------

Prerequisites:
- Go 1.20+
- git

Clone and build:

  git clone https://github.com/example/aurora-orchestrator.git
  cd aurora-orchestrator
  go build ./cmd/server

Run (default SQLite file-based DB):

  ./server

Or with environment override (e.g., to use in-memory SQLite for testing):

  export AURORA_DB_DSN="file::memory:?cache=shared"
  go run ./cmd/server

Usage
-----

Start server (default port 8080). Example requests:

Create an agent:

  curl -X POST localhost:8080/agents -H 'Content-Type: application/json' -d '{"name":"assistant-1","capabilities": {"can_search":true}}'

List agents:

  curl localhost:8080/agents

Get agent by id:

  curl localhost:8080/agents/1

Execute an agent (mock run):

  curl -X POST localhost:8080/agents/1/execute

The execute endpoint starts a background mock-run that updates the agent's status from "idle" -> "running" -> "completed".

Extending
---------

- Swap SQLite for PostgreSQL by changing the DSN and Gorm driver in internal/store/store.go
- Replace the mock execution in internal/agent/manager.go with a real AI call (LLM or tools) and richer task orchestration
- Add authentication middleware in internal/router or handlers

Files of interest
-----------------
- cmd/server/main.go — application bootstrap
- internal/agent/* — domain and manager logic
- internal/store/* — DB initialization and migrations
- internal/handlers/* — HTTP handlers
- docs/ARCHITECTURE.md — architecture and design decisions

License & Contributing
----------------------
This repo is a template. Adapt licensing and contribution policies to your project.

