\# M0-015 Git Ignore Review



\## Goal



Review and complete the repository root `.gitignore` for the current Go, Flutter, Docker and local development setup.



\---



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Project Conventions v1.1.pdf

\- docs/architecture/Infrastructure \& Deployment v1.0.pdf

\- docs/roadmap/Milestone 0 — Repository \& Development Environment.pdf



\---



\## Do



\- Review the existing root `.gitignore`.

\- Preserve currently valid ignore rules.

\- Add only missing rules required by the current repository.

\- Cover local environment files, Go build artifacts, Flutter generated files, IDE files and operating-system artifacts.

\- Keep tracked templates such as `.env.example`.

\- Confirm that required source files and repository placeholders remain trackable.

\- Update PROJECT\_STATUS.md only after all checks pass.



\---



\## Required coverage



Review rules for:



```text

.env

.env.\*

!.env.example



Go build and test artifacts



Flutter and Dart generated files



IDE files



Operating-system files



Local logs and temporary files

