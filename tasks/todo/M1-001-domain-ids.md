\# Task: M1-001 — Domain IDs



\## Goal



Implement immutable Domain ID value objects used by the Calculator Engine.



This task establishes the identifier foundation for the domain model.

No business rules or engine behavior should be implemented.



\---



\## Read



Read only the following documents before implementation:



\- docs/architecture/Engine Architecture v1.0.pdf

\- docs/architecture/MVP Domain Model v1.0.pdf

\- docs/architecture/Project Conventions v1.1.pdf



Do not read unrelated documents.



\---



\## Implement



Create immutable value objects for all domain identifiers required by the MVP domain model.



Requirements:



\- strongly typed IDs

\- constructors

\- validation

\- zero mutable state

\- comparable

\- JSON compatible where required

\- no database code

\- no API code



\---



\## Do Not



Do not implement:



\- RuleSet

\- Calculator Engine

\- ResolveWalk

\- Events

\- Database

\- PostgreSQL

\- API

\- Mobile

\- Simulator

\- H3

\- Score calculation

\- Territory logic



\---



\## Validation



Run:



```text

gofmt ./...

go test ./...

make backend-check

git diff --check

```

