<div align="center">
  <h1>godbmigrate</h1>
  <p>A fast, flexible, and database-agnostic migration tool for Go projects.</p>

  <img src="assets/github-go.png" alt="godbmigrate Banner" width="600px">

  <br>

[![CI](https://img.shields.io/github/actions/workflow/status/ESousa97/godbmigrate/ci.yml?branch=master&label=CI&logo=github)](https://github.com/ESousa97/godbmigrate/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ESousa97/godbmigrate)](https://goreportcard.com/report/github.com/ESousa97/godbmigrate)
[![Go Reference](https://pkg.go.dev/badge/github.com/ESousa97/godbmigrate.svg)](https://pkg.go.dev/github.com/ESousa97/godbmigrate)
[![License: MIT](https://img.shields.io/github/license/ESousa97/godbmigrate?color=blue)](https://github.com/ESousa97/godbmigrate/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ESousa97/godbmigrate)](https://github.com/ESousa97/godbmigrate)
[![Last Commit](https://img.shields.io/github/last-commit/ESousa97/godbmigrate)](https://github.com/ESousa97/godbmigrate/commits/master)

</div>

---

`godbmigrate` is a lightweight CLI tool and Go library designed to handle database migrations with ease. It supports PostgreSQL out-of-the-box and focuses on simplicity, speed, and safety through advisory locks to prevent concurrent execution in distributed environments.

## Demonstration

### CLI Usage

```bash
# Create a new migration
godbmigrate new add_users_table

# Apply all pending migrations
godbmigrate up --dsn "postgres://user:pass@localhost:5432/dbname?sslmode=disable"

# Revert the last migration
godbmigrate down --dsn "postgres://user:pass@localhost:5432/dbname?sslmode=disable"
```

### Library Usage

```go
import "github.com/ESousa97/godbmigrate/internal/db"

// Connect to the database
store, err := db.Connect(dsn)
if err != nil {
    log.Fatal(err)
}
defer func() {
    _ = store.Close()
}()

// Apply pending migrations
if err := store.ApplyMigration(version, sqlContent); err != nil {
    log.Fatal(err)
}
```

## Tech Stack

| Technology | Role |
|---|---|
| Go 1.25 | Core language and concurrent execution |
| Cobra | Framework for creation of powerful CLI applications |
| PostgreSQL | Target database (initial support) |
| Slog | Native structured logging for Go |
| Advisory Locks | Distributed lock mechanism for safe migrations |

## Prerequisites

- Go >= 1.25.0 (defined in `go.mod`)
- PostgreSQL (or Docker for local setup)
- Network tools (curl)

## Installation and Usage

### As a Binary

```bash
go install github.com/ESousa97/godbmigrate@latest
```

### From Source

```bash
git clone https://github.com/ESousa97/godbmigrate.git
cd godbmigrate
make build
# make test-full
```

## Makefile Targets

| Target | Description |
|---|---|
| `make build` | Compiles the binary to `godbmigrate.exe` |
| `make db-up` | Starts a PostgreSQL container via Docker |
| `make db-down` | Stops and removes the PostgreSQL container |
| `make test-full` | Executes a complete test cycle (build, new, up, status, down) |
| `make clean` | Removes binaries and temporary migration files |

## Architecture

The project adopts a modular architecture focused on safety and simplicity:

- **`cmd/`**: CLI interface and application bootstrapping using Cobra.
- **`internal/db/`**: Core logic containing the `MigrationStore` and SQL execution.
- **`Advisory Locks`**: PostgreSQL-native distributed locks ensuring single-process execution.
- **`Schema Tracking`**: Automated management of the `schema_migrations` table.

## API Reference

Detailed documentation of the internal packages and structures can be found at [pkg.go.dev](https://pkg.go.dev/github.com/ESousa97/godbmigrate).

## Configuration

| Variable | Description | Type | Default |
|---|---|---|---|
| `--dsn` | PostgreSQL connection string | string | `postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable` |
| `--debug` | Enables DEBUG level logs | bool | `false` |

## Roadmap

- [x] Phase 1: PostgreSQL Basic Support
- [x] Phase 2: Advisory Locks & Concurrency Control
- [ ] Phase 3: Multiple Database Support (MySQL, SQLite)
- [ ] Phase 4: Programmatic Go Migrations
- [ ] Phase 5: Advanced CI/CD Integration & Linting

## Contributing

Contributions are welcome! See the full guide at [CONTRIBUTING.md](CONTRIBUTING.md).

## License

Distributed under the MIT license. See [LICENSE](LICENSE) for more details.

<div align="center">

## Author

**Enoque Sousa**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/enoque-sousa-bb89aa168/)
[![GitHub](https://img.shields.io/badge/GitHub-100000?style=flat&logo=github&logoColor=white)](https://github.com/ESousa97)
[![Portfolio](https://img.shields.io/badge/Portfolio-FF5722?style=flat&logo=target&logoColor=white)](https://enoquesousa.vercel.app)

**[⬆ Back to top](#godbmigrate)**

Made with ❤️ by [Enoque Sousa](https://github.com/ESousa97)

**Project Status:** Active — Study Project

</div>
