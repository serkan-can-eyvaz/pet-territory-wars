# Pet Territory Wars

Pet Territory Wars is an MVP for a location-based territory game. The repository is a monorepo with a Go backend, a Flutter mobile application, PostgreSQL with PostGIS, and Docker Compose for local development.

## Repository structure

- `backend/` - Go API, workers, configuration, infrastructure, and migrations.
- `mobile/` - Flutter mobile application.
- `backend/migrations/` - SQL migrations managed independently from application startup.
- `compose.yaml` - Local API, walk worker, outbox worker, and PostgreSQL/PostGIS services.
- `Makefile` - Standard development, validation, Compose, and migration commands.

## Prerequisites

- Docker Desktop with Docker Compose v2
- GNU Make
- Go 1.26.5 (defined in `backend/go.mod`)
- Flutter stable SDK
- Standalone `golang-migrate` CLI for migration commands

## Environment setup

Copy the tracked template before running local services:

```sh
cp .env.example .env
```

Keep `.env` local and do not commit it. Before running Docker Compose, define `POSTGRES_DB`, `POSTGRES_USER`, and `POSTGRES_PASSWORD` in `.env`; Compose requires these values. Migration commands also require `DATABASE_URL` to be available in the shell environment.

## Development commands

Show the available commands:

```sh
make help
```

Validate the Compose configuration, start local services, stop them, or view their logs:

```sh
make compose-config
make compose-up
make compose-down
make compose-logs
```

Apply, inspect, or roll back migrations. `migrate-down` asks for confirmation before rolling back all migrations:

```sh
make migrate-up
make migrate-version
make migrate-down
```

## Validation commands

Run the standard backend and Flutter checks from the repository root:

```sh
make backend-check
make flutter-check
```
