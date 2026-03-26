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
