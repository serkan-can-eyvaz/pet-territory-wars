\# Task ID



M0-018



\# Title



Local Installation Verification



\# Goal



Verify that a clean developer can clone the repository and follow the documented setup instructions without repository changes.



\# Read Before Starting



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Pet Territory Wars — Project Conventions v1.1.pdf

\- docs/architecture/Pet Territory Wars — Infrastructure \& Deployment v1.0.pdf

\- README.md



\# Scope



\- Follow the installation steps documented in the repository.

\- Verify that the documented commands are correct.

\- Verify that no required step is missing.

\- Verify that repository structure matches the documentation.

\- Verify that setup instructions are reproducible.

\- Fix documentation only if it is incorrect or incomplete.

\- Do not modify source code, infrastructure, workflows or configuration unless required to correct documentation.



\# Validation



Run:



```text

git diff --check

make help

make backend-check

make flutter-check

docker compose config --quiet

```



\# Expected Changes



\- README.md (only if documentation corrections are required)

\- PROJECT\_STATUS.md



\# Completion



If every validation passes:



```text

Task: Milestone 1

Task File: tasks/todo/M1-001-calculator-engine.md

Next Task: M1-001 — Calculator Engine

Status: In Progress

```



Otherwise:



\- Leave PROJECT\_STATUS.md unchanged.



\# Output Format



Task ID:

Changed files:

Checks run:

Remaining risk:

