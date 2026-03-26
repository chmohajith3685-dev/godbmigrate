# godbmigrate

A simple and efficient database migration engine for Go projects.

## Installation

```bash
go get github.com/ESousa97/godbmigrate
```

## Usage

### Create a New Migration
```bash
godbmigrate new <migration_name>
```

This will generate two files in the `migrations/` directory:
- `YYYYMMDDHHMMSS_<name>.up.sql`
- `YYYYMMDDHHMMSS_<name>.down.sql`

### List Migrations
```bash
godbmigrate list
```

## Technologies
- Go (Golang)
- Cobra CLI

## Roadmap

### Phase 1: Discovery and Standardization (CLI + File System) [DONE]
**Goal:** Organize how migration files are named and read from disk.
- [x] Create CLI structure using Cobra.
- [x] Implement `new <name>` command to generate `.up.sql` and `.down.sql` files with timestamps.
- [x] Implement logic to list and sort migration files from the `migrations/` directory.

### Phase 2: State Control (Metadata Table) [PENDING]
**Goal:** Track applied and pending migrations to ensure consistency.
- Connect to PostgreSQL using `database/sql`.
- Automatically create the `schema_migrations` table (columns: `version` INT, `applied_at` TIMESTAMP).
- Implement a function to retrieve the last applied version.

### Phase 3: The Executor (Transactions and Pure SQL) [PENDING]
**Goal:** Apply migrations safely with full atomicity.
- Implement the `up` command.
- Identify pending migrations and execute them in order.
- Use `db.Begin()` for each migration file to ensure transactional integrity (rollback on failure).

### Phase 4: Reversal and Integrity (Down Migrations) [PENDING]
**Goal:** Provide controlled rollback capabilities.
- Implement the `down` command to revert the last applied migration.
- Remove the reverted version from the `schema_migrations` table.
- Add a `-all` flag to rollback all applied migrations.

### Phase 5: Production Security (Advisory Locks) [PENDING]
**Goal:** Ensure safe concurrent execution in production environments.
- Use `pg_advisory_lock` to prevent multiple instances from running migrations simultaneously.
- Implement professional logging using `slog`.
- Create a `Makefile` with Docker support for end-to-end testing.
