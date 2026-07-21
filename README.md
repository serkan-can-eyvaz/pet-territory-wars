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

Copy the tracked template before running local services.

PowerShell:

```powershell
Copy-Item .env.example .env
```

Unix shells:

```sh
cp .env.example .env
```

Keep `.env` local and do not commit it. The template includes safe local values for `POSTGRES_DB`, `POSTGRES_USER`, `POSTGRES_PASSWORD`, and `DATABASE_URL`. Docker Compose reads the PostgreSQL values from `.env`; API and worker containers use the Docker-network hostname `postgres` to reach the database.

The example `DATABASE_URL` is for host-side migration commands and uses `localhost:5433`. Start PostgreSQL before running migrations:

```sh
docker compose up -d postgres
```

Before running a migration command, make `DATABASE_URL` from `.env` available in your current shell environment using that shell's normal environment-loading mechanism. Then run:

```sh
make migrate-up
make migrate-version
make migrate-down
```

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

## Validation commands

Run the standard backend and Flutter checks from the repository root:

```sh
make backend-check
make flutter-check
```
