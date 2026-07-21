.DEFAULT_GOAL := help

.PHONY: help backend-format backend-format-check backend-vet backend-test backend-build backend-check flutter-format flutter-format-check flutter-analyze flutter-test flutter-check compose-config compose-up compose-down compose-logs migrate-up migrate-down migrate-version

help: ## Display available commands.
	@printf '%s\n' \
		'Available targets:' \
		'  backend-format   Format backend Go source files.' \
		'  backend-format-check  Check backend Go source formatting.' \
		'  backend-vet      Run backend static checks.' \
		'  backend-test     Run backend tests.' \
		'  backend-build    Build backend binaries.' \
		'  backend-check    Run all backend checks.' \
		'  flutter-format   Format Flutter source files.' \
		'  flutter-format-check  Check Flutter source formatting.' \
		'  flutter-analyze  Analyze Flutter source files.' \
		'  flutter-test     Run Flutter tests.' \
		'  flutter-check    Run all Flutter checks.' \
		'  compose-config   Validate Docker Compose configuration.' \
		'  compose-up       Build and start Docker Compose services.' \
		'  compose-down     Stop Docker Compose services.' \
		'  compose-logs     Show Docker Compose service logs.' \
		'  migrate-up       Apply database migrations.' \
		'  migrate-down     Roll back database migrations.' \
		'  migrate-version  Show the current migration version.'

backend-format: ## Format backend Go source files.
	cd backend && gofmt -w $$(find . -type f -name '*.go')

backend-format-check: ## Check backend Go source formatting.
	cd backend && unformatted="$$(gofmt -l $$(find . -type f -name '*.go'))"; if [ -n "$$unformatted" ]; then printf '%s\n' "$$unformatted"; exit 1; fi

backend-vet: ## Run backend static checks.
	cd backend && go vet ./...

backend-test: ## Run backend tests.
	cd backend && go test ./...

backend-build: ## Build backend binaries.
	cd backend && go build ./...

backend-check: ## Run all backend checks.
	@$(MAKE) backend-format-check
	@$(MAKE) backend-vet
	@$(MAKE) backend-test
	@$(MAKE) backend-build

flutter-format: ## Format Flutter source files.
	cd mobile && dart format .

flutter-format-check: ## Check Flutter source formatting.
	cd mobile && dart format --output=none --set-exit-if-changed .

flutter-analyze: ## Analyze Flutter source files.
	cd mobile && flutter analyze

flutter-test: ## Run Flutter tests.
	cd mobile && flutter test

flutter-check: ## Run all Flutter checks.
	cd mobile && flutter pub get
	@$(MAKE) flutter-format-check
	@$(MAKE) flutter-analyze
	@$(MAKE) flutter-test

compose-config: ## Validate Docker Compose configuration.
	@docker compose config --quiet

compose-up: ## Build and start Docker Compose services.
	docker compose up --build

compose-down: ## Stop Docker Compose services.
	docker compose down

compose-logs: ## Show Docker Compose service logs.
	docker compose logs

migrate-up: ## Apply database migrations.
	@migrate -path backend/migrations -database "$$DATABASE_URL" up

migrate-down: ## Roll back database migrations.
	@migrate -path backend/migrations -database "$$DATABASE_URL" down

migrate-version: ## Show the current migration version.
	@migrate -path backend/migrations -database "$$DATABASE_URL" version
