\# M0-014 Flutter CI



STATUS:Completed



\## Goal




Create the initial GitHub Actions workflow for the Flutter mobile application.



\---



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Mobile Architecture v1.0.pdf

\- docs/architecture/Project Conventions v1.1.pdf

\- docs/architecture/Infrastructure \& Deployment v1.0.pdf

\- docs/roadmap/Milestone 0 — Repository \& Development Environment.pdf



\---



\## Do



\- Create the Flutter GitHub Actions workflow.

\- Trigger on push and pull\_request.

\- Use a single Ubuntu-based job.

\- Check out the repository.

\- Set up Flutter using the project’s existing version source if available.

\- Run the existing Makefile target from the repository root.

\- Keep the workflow minimal.

\- Update PROJECT\_STATUS.md only after all checks pass.



\---



\## Required workflow



Run:



```text

make flutter-check

