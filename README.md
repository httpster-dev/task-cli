AHA! I've pushed to main!


# Task Manager CLI
## Phase 1: In-Memory Domain Logic

**Go concepts**: structs, `iota` enums, interfaces, pointer receivers, slices, sentinel errors, `time.Time`
**Rails analogy**: Task struct = AR model (explicit fields, no magic). TaskStore interface = the contract AR provides. InMemoryTaskStore = testing with a mock adapter.

| Step | What we build | Key test |
|------|--------------|----------|
| 1.1 | `Task` struct, `Status`/`Priority` types with `iota`, `String()` methods | Create a Task, assert fields |
| 1.2 | `TaskStore` interface in `store.go` (Add, List, Complete, Delete) | No tests — just the contract |
| 1.3 | `InMemoryTaskStore`: `Add` and `List` | Add 2 tasks, list them, verify IDs auto-increment |
| 1.4 | `InMemoryTaskStore`: `Complete` | Complete a task, assert Status=Done + CompletedAt set. Complete non-existent → `ErrTaskNotFound` |
| 1.5 | `InMemoryTaskStore`: `Delete` | Delete a task, verify list shrinks. Delete non-existent → `ErrTaskNotFound` |
| 1.6 | Extract `assert_test.go` with generic helpers | `assertEqual[T]`, `assertString`, `assertNoError`, `assertError`, `assertNotNil[T]` |

**Checkpoint**: `go test ./task/...` passes. Pure logic, zero I/O.

---

## Phase 2: CLI Layer

**Go concepts**: `os.Args`, `io.Writer` for testable output, `bytes.Buffer`, `strconv.Atoi`, `fmt.Fprintf`
**Rails analogy**: CLI dispatch = router mapping verbs to actions. `io.Writer` injection = `$stdout` injection in service tests.

| Step | What we build | Key test |
|------|--------------|----------|
| 2.1 | `CLI` struct with `store` + `out io.Writer`, `Run(args []string) error`, `add` command | `add "Buy groceries"` → store has 1 task, buffer contains "Added task 1" |
| 2.2 | `list` command + `FormatTaskTable` in `format.go` | List 2 tasks → output contains IDs, titles, statuses. Empty list → "No tasks" |
| 2.3 | `done` command | `done 1` → task is StatusDone. `done abc` → "invalid task ID" error. `done 999` → ErrTaskNotFound |
| 2.4 | `delete` command | Same error-case pattern as `done` |
| 2.5 | Unknown command + no-args handling | `frobnicate` → "unknown command". No args → usage message |
| 2.6 | `cmd/task/main.go` wiring (InMemoryTaskStore for now) | Manual test: `go run ./cmd/task add "Hello"` |

**Checkpoint**: `go test ./...` passes. CLI is fully testable via buffer injection. Data doesn't persist yet (intentional).

---

## Phase 3: SQLite Persistence

**Go concepts**: `database/sql`, `sql.Open`, `defer db.Close()`, `db.Exec`/`db.Query`/`db.QueryRow`, `rows.Scan`, `t.TempDir()`, `go get`
**Rails analogy**: `db.Exec(CREATE TABLE)` = migration. `rows.Scan` = manual column-to-field mapping (no ORM). `t.TempDir()` = DatabaseCleaner.

| Step | What we build | Key test |
|------|--------------|----------|
| 3.1 | `go get modernc.org/sqlite` | — |
| 3.2 | `NewSQLiteStore(path)` — opens DB, runs `CREATE TABLE IF NOT EXISTS` | Open temp DB, add a task, no error |
| 3.3 | `SQLiteStore.Add` | Add task → verify ID, title, status |
| 3.4 | `SQLiteStore.List` | Add 2 tasks → list returns both, fields scanned correctly |
| 3.5 | `SQLiteStore.Complete` | Complete → status + completed_at updated. Not found → ErrTaskNotFound (via RowsAffected) |
| 3.6 | `SQLiteStore.Delete` | Same RowsAffected pattern |
| 3.7 | `newTestStore(t)` helper + compile-time interface check: `var _ task.TaskStore = (*task.SQLiteStore)(nil)` | — |
| 3.8 | Wire SQLite into `main.go` with `~/.task-cli/tasks.db` | Manual test: add, quit, list — data persists |

**Checkpoint**: `go test ./...` passes. Tasks survive between runs.

---

## Phase 4: Polish

**Go concepts**: `flag.NewFlagSet` per subcommand, `encoding/json` + struct tags, `strings.Split`/`Join`, pointer-for-optional pattern, `fmt.Errorf` with `%w`
**Rails analogy**: `flag.NewFlagSet` = Thor's `method_option`. `json:"title"` struct tags = `as_json(only: [...])`. `ListOptions{Status: &s}` = AR scopes.

| Step | What we build | Key test |
|------|--------------|----------|
| 4.1 | `AddOptions` struct (title, priority, tags) — update interface + both stores | Add with priority=high + tags=["work"] → stored correctly |
| 4.2 | CLI `add --priority high --tag work` via `flag.NewFlagSet` | Parse flags, verify task created with options |
| 4.3 | `ListOptions` struct (Status *Status, Tag string) — filter in both stores | List --status=done → only done tasks. List --tag=work → only tagged tasks |
| 4.4 | `--format=json` on list command + `json:"..."` struct tags | JSON decode output → valid task array |
| 4.5 | Error UX: wrap errors with `%w`, helpful messages for bad input | `done abc` → `"abc" is not a valid task ID` |
| 4.6 | End-to-end integration test in `cmd/task/main_test.go` | Full flow: add → list → done → filter → json → delete |

**Checkpoint**: `go test ./...` passes. Feature-complete CLI.

---

