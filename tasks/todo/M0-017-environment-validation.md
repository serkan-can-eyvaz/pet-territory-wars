\# M0-017 Environment Validation



STATUS:Completed



\## Goal



Validate that the complete local development environment is ready for MVP development.



\---



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Project Conventions v1.1.pdf

\- docs/architecture/Infrastructure \& Deployment v1.0.pdf

\- docs/roadmap/Milestone 0 — Repository \& Development Environment.pdf



\---



\## Do



\- Verify the required development tools are available.

\- Verify the backend environment.

\- Verify the Flutter environment.

\- Verify Docker and Docker Compose.

\- Verify the standalone migration CLI.

\- Verify the Makefile commands execute correctly.

\- Do not modify the development environment.

\- Update PROJECT\_STATUS.md only after all checks pass.



\---



\## Required validation



Run:



```text

go version

flutter --version

docker --version

docker compose version

migrate -version

make help

make backend-check

make flutter-check

docker compose config --quiet

```



\---



\## Don't



\- Install software.

\- Upgrade dependencies.

\- Modify configuration.

\- Create scripts.

\- Change application code.

\- Change Docker configuration.

\- Change CI workflows.



\---



\## Expected scope



```text

PROJECT\_STATUS.md

```



\---



\## Validation



Run:



```text

git diff --check

```



No repository files except `PROJECT\_STATUS.md` should change.



\---



\## Done checklist



\- Required tools verified.

\- Backend validation passed.

\- Flutter validation passed.

\- Docker validation passed.

\- Migration CLI verified.

\- Makefile verified.

\- No environment modifications made.

\- Required validations passed.

\- PROJECT\_STATUS.md updated.



\---



\## Completion update



```text

Task: M0-018 — Local Installation Verification

Task File: tasks/todo/M0-018-local-installation-verification.md

Next Task: Milestone 1

Status: In Progress

```



\---



\## Output format



```text

Task ID:

Changed files:

Checks run:

Remaining risk:

```

