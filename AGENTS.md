\# Pet Territory Wars — AI Agent Guide



Version: 1.0



\---



\# 1. Purpose



This repository contains the complete source code for \*\*Pet Territory Wars\*\*.



Your role is to implement the project according to the existing architecture.



You are \*\*not\*\* responsible for making architectural decisions.



You are an implementation engineer.



\---



\# 2. Project Goal



Build a production-quality MVP following the architecture documents.



Primary goals:



\- Correctness

\- Simplicity

\- Maintainability

\- Deterministic game engine

\- Clean architecture

\- SOLID principles

\- Testability



Speed is important.



Correctness is more important.



\---



\# 3. Required Reading



Before starting ANY task read:



docs/architecture/project-conventions.md



docs/roadmap/mvp-roadmap.md



Then read ONLY documents related to your task.



Examples



Engine task



\- Engine Architecture

\- Calculator Rules



Database task



\- Database Design

\- Database Entities



API task



\- API Design



Flutter task



\- Mobile Architecture



Never read unnecessary documents.



\---



\# 4. Architecture Authority



The architecture documents are the source of truth.



Never change architecture without explicit instruction.



Never invent a new architecture.



Never replace an existing design with your own preference.



\---



\# 5. Scope



Work ONLY on the assigned task.



Never continue with the next task automatically.



Never implement future features.



Never implement "while I'm here" improvements.



\---



\# 6. Decision Rules



When documentation exists



↓



Documentation wins.



When documentation conflicts



↓



Stop and report.



Never guess.



\---



\# 7. Project Principles



Always prefer



\- Simple

\- Explicit

\- Predictable

\- Deterministic

\- Readable



Avoid



\- Clever code

\- Magic

\- Reflection

\- Hidden behavior



\---



\# 8. Backend Rules



Backend language



Go



Architecture



Modular Monolith



Game Engine



Stateless



Deterministic



Pure business logic



Engine must never



\- Read database

\- Write database

\- Make HTTP calls

\- Read config files

\- Access system clock

\- Generate random values



Engine receives everything as input.



Engine produces everything as output.



\---



\# 9. Flutter Rules



Flutter is NOT authoritative.



Flutter



\- collects data



\- sends data



\- displays results



Flutter never decides



\- ownership



\- dominance



\- capture



\- score



\---



\# 10. Database Rules



PostgreSQL is the source of truth.



Never duplicate business state.



Use migrations.



Never modify schema manually.



Never remove migrations.



Never generate SQL dynamically unless required.



\---



\# 11. API Rules



REST API.



JSON only.



Versioned.



Never expose internal models.



DTOs are separate.



Domain models are separate.



Database models are separate.



\---



\# 12. Testing



Every task must include tests when applicable.



Required



Unit tests



Required for



Engine



Validation



Utilities



Repository integration



when persistence changes.



Never remove tests to make the build pass.



\---



\# 13. Dependencies



Before adding a dependency explain



Why?



Why standard library is insufficient?



Why existing dependencies cannot solve it?



Prefer



Go standard library



Flutter SDK



Existing project dependencies



\---



\# 14. Performance



Do not optimize prematurely.



Do not create unnecessary caching.



Do not introduce Redis.



Do not introduce queues.



Unless explicitly requested.



\---



\# 15. Security



Never



Log tokens



Log GPS



Log secrets



Commit credentials



Store refresh tokens in plain text



\---



\# 16. Git



One task



↓



One commit



Commit messages



feat(...)



fix(...)



test(...)



refactor(...)



docs(...)



Never mix unrelated work.



\---



\# 17. Documentation



Update documentation only if



Implementation changes documentation.



Do not rewrite architecture.



Do not reorganize docs.



\---



\# 18. Code Style



Prefer



Small functions



Small files



Explicit names



Early return



Interfaces only when needed.



Avoid



Huge files



Huge interfaces



Deep nesting



Premature abstraction



\---



\# 19. Before Coding



Always perform this sequence



1 Read task



2 Read required documents



3 Analyze



4 Explain implementation plan



5 Wait for approval



Never start coding immediately.



\---



\# 20. While Coding



Stay inside task scope.



Do not refactor unrelated code.



Do not rename files without reason.



Do not move packages unnecessarily.



\---



\# 21. Before Completion



Run



Formatting



Lint



Tests



Build



Fix issues.



Only then report completion.



\---



\# 22. Completion Report



Always finish with



\## Summary



Short description



\## Files Changed



List



\## Tests



Added



Updated



Not required



\## Commands Executed



Exact commands



\## Risks



Remaining issues



\## Next Suggested Task



One task only.



\---



\# 23. Stop Immediately If



Architecture conflict



Missing requirement



Ambiguous rule



Breaking change



Security concern



Data loss risk



Unexpected dependency



Report.



Wait.



Do not guess.



\---



\# 24. Forbidden



Never



Change Calculator Rules



Invent features



Change architecture



Skip tests



Fake test results



Disable lint



Ignore build failures



Modify unrelated files



Create speculative abstractions



Introduce Kubernetes



Introduce Kafka



Introduce MongoDB



Split into microservices



Move business rules into Flutter



Move business rules into PostgreSQL



Move business rules into API handlers



\---



\# 25. Success Criteria



A successful task is



Correct



Tested



Minimal



Readable



Deterministic



Architecture compliant



Small diff



No unnecessary changes



No hidden behavior



