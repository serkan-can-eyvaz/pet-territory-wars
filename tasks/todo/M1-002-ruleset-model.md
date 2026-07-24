\# Task: M1-002 — RuleSet Model



\## Goal



Implement the immutable RuleSet domain model used by the Calculator Engine.



The RuleSet defines the complete set of configurable game rules consumed by the engine.

No rule evaluation or business logic should be implemented in this task.



\---



\## Read



Read only the following documents before implementation:



\- docs/architecture/Calculator Rules v1.0.pdf

\- docs/architecture/Engine Architecture v1.0.pdf

\- docs/architecture/MVP Domain Model v1.0.pdf

\- docs/architecture/Project Conventions v1.1.pdf



Do not read unrelated documents.



\---



\## Implement



Implement the immutable RuleSet model.



Requirements:



\- immutable

\- strongly typed

\- no mutable state

\- no methods containing business logic

\- use previously created Domain IDs where required

\- support JSON serialization if required by the architecture

\- organize related configuration into clear nested structures when appropriate



This task only models configuration.



\---



\## Do Not



Do not implement:



\- Rule validation

\- Engine

\- ResolveWalk

\- Engine Input

\- Engine Result

\- Territory calculation

\- Score calculation

\- Events

\- Database

\- API

\- Simulator

\- H3 logic



\---



\## Validation



Run:



```text

gofmt ./...

go test ./...

make backend-check

git diff --check

```



\---



\## Pass Criteria



\- RuleSet compiles.

\- RuleSet is immutable.

\- No business logic exists.

\- No validation logic exists.

\- Tests pass.

\- No unrelated files modified.



\---



\## Expected Changed Files



Only RuleSet model files and corresponding tests.



\---



\## Completion



Update PROJECT\_STATUS.md only after all validations pass.

