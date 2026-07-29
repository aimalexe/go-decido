# go-decido

A terminal decision matrix in Go — weigh criteria, score alternatives, and compare options.

This project is a hands-on way to learn Go: packages, structs, methods, interfaces, errors, CLI input, files, JSON, menus, and table-driven tests. Expect more concepts as the app grows.

## Run

```bash
go run .
```

## Test

```bash
go test ./...
```

## Layout

| Package | Role |
|---------|------|
| `internal/decision` | Pure domain (no I/O) |
| `internal/service` | Orchestration + in-memory and JSON `Store` implementations |
| `internal/ui` | Terminal tables and prompts |
| `internal/menu` | Main and workspace menus |
| `internal/input` | Shared stdin helpers |
| `internal/app` | CLI loop wiring |
