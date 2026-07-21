\# M0-013 Backend CI

STATUS:Completed



\## Goal



Create the initial GitHub Actions workflow for the Go backend so every push and pull request runs the standard backend validation pipeline.



\---



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Infrastructure \& Deployment v1.0.pdf

\- docs/architecture/Project Conventions v1.1.pdf

\- docs/roadmap/Milestone 0 — Repository \& Development Environment.pdf



\---



\## Do



\- Create the backend GitHub Actions workflow.

\- Trigger on push and pull\_request.

\- Use the current supported Go version defined by the project.

\- Check out the repository.

\- Set up Go.

\- Execute the existing Makefile target instead of duplicating commands.

\- Fail immediately if any backend validation fails.

\- Keep the workflow minimal and deterministic.

\- Update PROJECT\_STATUS.md only after every validation succeeds.



\---



\## Required workflow



The workflow must execute:



```text

make backend-check

```



\---



\## Don't



\- Add caching.

\- Add matrix builds.

\- Add code coverage.

\- Add lint tools.

\- Add release jobs.

\- Add deployment.

\- Add Docker jobs.

\- Add Flutter jobs.

\- Add notifications.

\- Duplicate backend commands already provided by the Makefile.



\---



\## Expected scope



```text

.github/workflows/backend-ci.yml

PROJECT\_STATUS.md

```



\---



\## Validation



Run:



```text

git diff --check

```



Validate the workflow file using the available GitHub Actions validation method.



\---



\## Done checklist



\- Backend workflow created.

\- Runs on push.

\- Runs on pull\_request.

\- Uses repository Makefile.

\- Executes make backend-check.

\- No duplicated backend commands.

\- Validation completed.

\- PROJECT\_STATUS.md updated.



\---



\## Completion update



Task:

M0-014 Flutter CI



Next:

M0-015 Git Ignore Review



\---



\## Output format



Task ID:

Changed files:

Checks run:

Remaining risk:

