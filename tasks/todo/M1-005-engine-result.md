\# M1-005 — Engine Result



\## Goal



Calculator Engine'in `ResolveWalk` çağrısı sonunda döndürdüğü immutable sonuç veri sözleşmesini oluşturmak.



Bu görev yalnızca Engine Result modellerini, bunların doğrudan bağımlı value type'larını ve canonical sabitlerini tanımlar.



Hesaplama, validation davranışı, route segment üretimi, scoring logic, event generation veya persistence içermez.



\---



\## Roadmap Position



Bu görev Milestone 1 roadmap'indeki aşağıdaki adımdır:



\- M1-005 — Engine Result



Sonraki görev:



\- M1-006 — Route Segment Calculation



Roadmap numarası veya görev sırası değiştirilmemelidir.



\---



\## Read



Yalnızca:



\- Calculator Rules v1.0

&#x20; - Section 10 — Walk Validity

&#x20; - Section 10.1 — Rejection Reasons

&#x20; - Section 12.1 — VisitedHex

&#x20; - Section 18 — HexChange Contract

&#x20; - Section 19 — ScoreChange Contract

&#x20; - Section 19.1 — Canonical Score Metrics

&#x20; - Section 20 — Domain Event Contract

&#x20; - Section 20.1 — Event Common Fields

&#x20; - Section 21 — ResolveWalkResult and Metadata

\- Project Conventions v1.1

\- Existing files:

&#x20; - backend/internal/gameengine/hex\_change.go

&#x20; - backend/internal/domain/id/ids.go



Do not read unrelated architecture documents.



\---



\## Normative Source and Gap Resolution



Calculator Rules v1.0 is the normative source.



The document completely defines:



\- `ResolveWalkResult` fields

\- `CalculationMetadata`

\- `VisitedHex`

\- `HexChange`

\- `ScoreChange`

\- canonical score metrics

\- canonical event types

\- event common-field semantics

\- result status values

\- walk validity status values

\- rejection reason values



The document references but does not provide complete Go struct declarations for:



\- `ValidationResult`

\- `WalkMetrics`

\- `DomainEvent`

\- `EventPayload`



This task provides the exact MVP contracts for those missing declarations.



These contracts are explicit architecture decisions for M1-005 and must be implemented exactly as written.



Do not infer additional fields.



\---



\## Existing HexChange



`HexChange`, `HexChangeType` and `HexChangeReason` are already defined in:



```text

backend/internal/gameengine/hex\_change.go

```



Reuse those existing types.



Do not redefine, rename or modify them in this task unless compilation proves that the existing implementation contradicts Calculator Rules Section 18.



No `HexChangeReason` constants are added in this task.



\---



\# Required Models



\## 1. WalkResolutionStatus



Define:



```go

type WalkResolutionStatus string

```



Canonical values:



```go

const (

&#x20;   WalkResolutionStatusResolved WalkResolutionStatus = "RESOLVED"

&#x20;   WalkResolutionStatusRejected WalkResolutionStatus = "REJECTED"

)

```



No additional status value may be added.



Meaning:



\- `RESOLVED`: A `VALID` or `PARTIALLY\_VALID` walk completed engine resolution.

\- `REJECTED`: Business validity thresholds were not met; this is not a technical error.



\---



\## 2. WalkValidationStatus



Define:



```go

type WalkValidationStatus string

```



Canonical values:



```go

const (

&#x20;   WalkValidationStatusValid          WalkValidationStatus = "VALID"

&#x20;   WalkValidationStatusPartiallyValid WalkValidationStatus = "PARTIALLY\_VALID"

&#x20;   WalkValidationStatusInvalid        WalkValidationStatus = "INVALID"

)

```



No additional validation status may be added.



\---



\## 3. WalkRejectionReason



Define:



```go

type WalkRejectionReason string

```



Canonical values, in canonical ordering:



```go

const (

&#x20;   WalkRejectionReasonTooShortDuration     WalkRejectionReason = "TOO\_SHORT\_DURATION"

&#x20;   WalkRejectionReasonTooShortDistance     WalkRejectionReason = "TOO\_SHORT\_DISTANCE"

&#x20;   WalkRejectionReasonLowValidRouteRatio   WalkRejectionReason = "LOW\_VALID\_ROUTE\_RATIO"

&#x20;   WalkRejectionReasonNoValidSegments      WalkRejectionReason = "NO\_VALID\_SEGMENTS"

&#x20;   WalkRejectionReasonMockLocationDetected WalkRejectionReason = "MOCK\_LOCATION\_DETECTED"

&#x20;   WalkRejectionReasonOutsideActiveCity    WalkRejectionReason = "OUTSIDE\_ACTIVE\_CITY"

)

```



Do not add parsing, validation or ordering helpers.



Ordering behavior will be implemented in M1-007 Walk Validation, not in this task.



\---



\## 4. ValidationResult



Define exactly:



```go

type ValidationResult struct {

&#x20;   Status           WalkValidationStatus

&#x20;   RejectionReasons \[]WalkRejectionReason

}

```



Semantics:



\- `Status` represents `VALID`, `PARTIALLY\_VALID` or `INVALID`.

\- `RejectionReasons` contains business rejection reasons.

\- A valid or partially valid result may carry an empty reason slice.

\- An invalid result may carry one or more canonical reasons.

\- This task does not enforce these relationships.



Do not add technical errors to `ValidationResult`.



Technical failures are returned through Go `error`, not as business rejection reasons.



\---



\## 5. WalkMetrics



Define exactly:



```go

type WalkMetrics struct {

&#x20;   TotalDistanceMeters float64

&#x20;   ValidDistanceMeters float64

&#x20;   ValidDurationSeconds int

&#x20;   ValidRouteRatio      float64

}

```



Semantics:



\- `TotalDistanceMeters` is the sum of all calculable segment distances.

\- `ValidDistanceMeters` is the sum of valid segment distances.

\- `ValidDurationSeconds` is the sum of valid segment durations in whole seconds.

\- `ValidRouteRatio` is valid distance divided by total distance, or zero when total distance is zero.



This task only defines the fields.



Do not calculate, round, normalize or validate metric values.



\---



\## 6. VisitedHex



Define exactly:



```go

type VisitedHex struct {

&#x20;   HexID           id.HexID

&#x20;   PresenceSeconds int

&#x20;   DistanceMeters  float64

&#x20;   EntryCount      int

&#x20;   FirstEnteredAt  time.Time

&#x20;   LastExitedAt    time.Time

}

```



Do not add:



\- `Qualified`

\- owner information

\- dominance information

\- score information

\- helper methods



Qualification is calculated later by the H3 aggregation task.



\---



\## 7. ScoreMetric



Define:



```go

type ScoreMetric string

```



Canonical values:



```go

const (

&#x20;   ScoreMetricActiveHex       ScoreMetric = "ACTIVE\_HEX"

&#x20;   ScoreMetricCapture         ScoreMetric = "CAPTURE"

&#x20;   ScoreMetricLifetimeCapture ScoreMetric = "LIFETIME\_CAPTURE"

&#x20;   ScoreMetricDefense         ScoreMetric = "DEFENSE"

&#x20;   ScoreMetricSteal           ScoreMetric = "STEAL"

&#x20;   ScoreMetricLifetimeSteal   ScoreMetric = "LIFETIME\_STEAL"

)

```



No additional score metric may be added.



\---



\## 8. ScoreChangeReason



Calculator Rules does not define a closed canonical value set for this type.



Define only:



```go

type ScoreChangeReason string

```



Do not define `ScoreChangeReason` constants.



Do not add validation, parsing, formatting or default values.



Territory outcome codes such as `EMPTY\_CAPTURE` and `OWNERSHIP\_TRANSFER` may be carried later as strongly typed reason values, but this task must not create an undocumented closed enum.



\---



\## 9. ScoreChange



Define exactly:



```go

type ScoreChange struct {

&#x20;   PlayerID id.PlayerID

&#x20;   Metric   ScoreMetric

&#x20;   Delta    int64

&#x20;   Reason   ScoreChangeReason

&#x20;   WalkID   id.WalkID

}

```



This task does not:



\- calculate deltas

\- reject zero deltas

\- merge score changes

\- order score changes



Those behaviors belong to M1-016 — Score Changes.



\---



\## 10. EventType



Define:



```go

type EventType string

```



Canonical values:



```go

const (

&#x20;   EventTypeWalkValidated             EventType = "WALK\_VALIDATED"

&#x20;   EventTypeWalkPartiallyValidated    EventType = "WALK\_PARTIALLY\_VALIDATED"

&#x20;   EventTypeWalkRejected              EventType = "WALK\_REJECTED"

&#x20;   EventTypeEmptyHexCaptured          EventType = "EMPTY\_HEX\_CAPTURED"

&#x20;   EventTypeHexDefended               EventType = "HEX\_DEFENDED"

&#x20;   EventTypeHexAttacked               EventType = "HEX\_ATTACKED"

&#x20;   EventTypeHexUnderThreat            EventType = "HEX\_UNDER\_THREAT"

&#x20;   EventTypeHexOwnershipTransferred   EventType = "HEX\_OWNERSHIP\_TRANSFERRED"

&#x20;   EventTypePlayerScoreChanged        EventType = "PLAYER\_SCORE\_CHANGED"

)

```



No additional event type may be added.



\---



\## 11. EventPayload



Calculator Rules requires event-type-specific immutable value data but does not define concrete payload contracts.



For M1-005 define:



```go

type EventPayload interface{}

```



Architecture decision:



\- `EventPayload` is a domain-level payload contract.

\- Concrete event payload value structs will be defined during M1-017 — Domain Events.

\- This task must not define event-specific payload structs.

\- This task must not use persistence or transport DTO types.

\- This task must not use `map\[string]any`.

\- This task must not add marker methods.



\---



\## 12. DomainEvent



Define exactly:



```go

type DomainEvent struct {

&#x20;   ID             id.EventID

&#x20;   Type           EventType

&#x20;   OccurredAt     time.Time

&#x20;   WalkID         id.WalkID

&#x20;   PlayerID       id.PlayerID

&#x20;   HexID          \*id.HexID

&#x20;   RuleSetVersion id.RuleSetVersion

&#x20;   EngineVersion  id.EngineVersion

&#x20;   Payload        EventPayload

}

```



Architecture decisions:



\- `HexID` is `\*id.HexID` because walk-level events do not belong to one hex.

\- `nil HexID` represents a non-hex event.

\- `Payload` may be `nil` until concrete payload contracts are defined.

\- `ID` is carried as data; this task does not generate it.

\- `OccurredAt` is carried as data; this task does not assign `EvaluatedAt`.

\- Event ordering is not implemented in this task.



\---



\## 13. CalculationMetadata



Define exactly:



```go

type CalculationMetadata struct {

&#x20;   EngineVersion  id.EngineVersion

&#x20;   RuleSetVersion id.RuleSetVersion

&#x20;   InputHash      string

&#x20;   EvaluatedAt    time.Time

}

```



Use the normative name `CalculationMetadata`.



Do not use:



\- `EngineMetadata`

\- `ResultMetadata`

\- `OutputMetadata`



This task does not calculate `InputHash`.



\---



\## 14. ResolveWalkResult



Define exactly:



```go

type ResolveWalkResult struct {

&#x20;   Status        WalkResolutionStatus

&#x20;   Validation    ValidationResult

&#x20;   Metrics       WalkMetrics

&#x20;   VisitedHexes  \[]VisitedHex

&#x20;   HexChanges    \[]HexChange

&#x20;   ScoreChanges  \[]ScoreChange

&#x20;   Events        \[]DomainEvent

&#x20;   Metadata      CalculationMetadata

}

```



Field names and order must remain exactly as shown.



Do not add:



\- `WalkID`

\- `PlayerID`

\- `Accepted`

\- `DistanceMeters`

\- `DurationSeconds`

\- `Error`

\- `ErrorMessage`



Walk and player identities are already represented by nested output records where required.



\---



\# Files



Create:



```text

backend/internal/gameengine/result.go

backend/internal/gameengine/result\_test.go

```



Reuse without modifying unless compilation requires a normative correction:



```text

backend/internal/gameengine/hex\_change.go

backend/internal/gameengine/hex\_change\_test.go

```



Update only after every required check passes:



```text

PROJECT\_STATUS.md

```



Task-management files may be moved or removed only when necessary to replace the invalid previous M1-005 task.



Do not rename roadmap task numbers.



\---



\# Package and Imports



All result types must be declared in:



```go

package gameengine

```



Allowed standard-library import:



```go

time

```



Allowed internal import:



```go

<module-path>/internal/domain/id

```



Use the repository's actual Go module path.



Do not add a third-party dependency.



\---



\# Implementation Requirements



\- All fields must be exported.

\- Domain IDs must use existing strongly typed ID types.

\- Time values must use `time.Time`.

\- Models must only carry data.

\- No production defaults.

\- No constructors.

\- No builders.

\- No getters or setters.

\- No validation methods.

\- No parsing helpers.

\- No `String()` methods.

\- No marshal/unmarshal methods.

\- No JSON tags.

\- No database tags.

\- No ORM tags.

\- No reflection-based production behavior.

\- No defensive-copy behavior.

\- No business calculation.

\- No event generation.

\- No score generation.

\- No input hashing.



\---



\# Required Tests



Create focused data-contract tests in:



```text

backend/internal/gameengine/result\_test.go

```



Tests must verify the following.



\## Status Constants



Verify exact string values:



\- `RESOLVED`

\- `REJECTED`

\- `VALID`

\- `PARTIALLY\_VALID`

\- `INVALID`



\## Rejection Reason Constants



Verify exact canonical values:



\- `TOO\_SHORT\_DURATION`

\- `TOO\_SHORT\_DISTANCE`

\- `LOW\_VALID\_ROUTE\_RATIO`

\- `NO\_VALID\_SEGMENTS`

\- `MOCK\_LOCATION\_DETECTED`

\- `OUTSIDE\_ACTIVE\_CITY`



\## Score Metric Constants



Verify exact canonical values:



\- `ACTIVE\_HEX`

\- `CAPTURE`

\- `LIFETIME\_CAPTURE`

\- `DEFENSE`

\- `STEAL`

\- `LIFETIME\_STEAL`



\## Event Type Constants



Verify exact canonical values:



\- `WALK\_VALIDATED`

\- `WALK\_PARTIALLY\_VALIDATED`

\- `WALK\_REJECTED`

\- `EMPTY\_HEX\_CAPTURED`

\- `HEX\_DEFENDED`

\- `HEX\_ATTACKED`

\- `HEX\_UNDER\_THREAT`

\- `HEX\_OWNERSHIP\_TRANSFERRED`

\- `PLAYER\_SCORE\_CHANGED`



\## ValidationResult



Verify:



\- validation status is preserved

\- an empty rejection list can be carried

\- multiple ordered rejection reasons can be carried without mutation



Do not test validation behavior.



\## WalkMetrics



Verify all four metric values are preserved.



Do not test formulas or rounding.



\## VisitedHex



Verify:



\- strong `HexID`

\- presence seconds

\- distance

\- entry count

\- first-entered timestamp

\- last-exited timestamp



are preserved.



Do not test qualification or ordering.



\## ScoreChange



Verify:



\- strong player and walk IDs

\- metric

\- positive and negative `int64` deltas

\- arbitrary strongly typed reason code



are preserved.



Do not test score calculation.



\## DomainEvent



Verify:



\- common fields are preserved

\- `HexID == nil` supports walk-level events

\- non-nil `\*id.HexID` supports hex events

\- `Payload == nil` can be carried

\- a simple immutable test payload struct can be assigned through `EventPayload`



Do not define production payload structs in the test.



\## CalculationMetadata



Verify all values are preserved.



Do not calculate or validate a hash.



\## ResolveWalkResult



Construct one complete result value and verify:



\- all nested values are preserved

\- slices carry their elements

\- existing `HexChange` values are reused

\- no engine calculation is performed



\---



\# Test Restrictions



Tests must not:



\- use reflection to inspect field order

\- inspect missing struct tags

\- test route calculation

\- test walk validity algorithms

\- test H3 behavior

\- test territory behavior

\- test score generation

\- test event generation

\- test hashing

\- add constructors solely for tests



Compile-time field usage and value-preservation tests are sufficient.



\---



\# Do Not



\- Modify the roadmap.

\- Create a new milestone task number.

\- Implement M1-006 functionality.

\- Implement route segments.

\- Implement segment invalid reasons.

\- Implement input validation.

\- Implement walk validation calculations.

\- Implement interpolation.

\- Implement H3 conversion.

\- Implement territory decay.

\- Implement capture, defense, attack or transfer behavior.

\- Implement score-change generation.

\- Implement domain-event generation.

\- Implement input hashing.

\- Implement `ResolveWalk`.

\- Add persistence behavior.

\- Add API DTOs.

\- Add database models.

\- Add serialization tags.

\- Add third-party packages.

\- Invent additional canonical values.



\---



\# Expected Changed Files



Expected:



```text

backend/internal/gameengine/result.go

backend/internal/gameengine/result\_test.go

PROJECT\_STATUS.md

```



Task-file replacement or cleanup is also allowed:



```text

tasks/todo/M1-005-engine-result.md

```



Existing `hex\_change.go` files should remain unchanged unless a compilation issue reveals a direct contradiction with Calculator Rules Section 18.



No other production file should change.



\---



\# Validation



From the backend directory where appropriate:



```bash

gofmt -w internal/gameengine/result.go internal/gameengine/result\_test.go

go test ./...

```



From the repository root:



```bash

make backend-check

git diff --check

```



All checks must pass.



If a check fails:



1\. Fix only the issue inside this task's scope.

2\. Run the complete validation sequence again.

3\. Do not update `PROJECT\_STATUS.md` until every check passes.



\---



\# Pass Criteria



The task is complete only when all conditions are satisfied:



\- Roadmap task remains `M1-005 — Engine Result`.

\- `ResolveWalkResult` matches Calculator Rules Section 21.

\- The normative name `CalculationMetadata` is used.

\- `ValidationResult` uses the exact task-level contract.

\- `WalkMetrics` uses the exact task-level contract.

\- `VisitedHex` matches Section 12.1.

\- Existing `HexChange` is reused.

\- `ScoreChange` matches Section 19.

\- Canonical score metrics match Section 19.1.

\- `DomainEvent` includes all common fields from Section 20.1.

\- Canonical event types match Section 20.

\- Undefined reason sets are not invented.

\- Models contain no business logic.

\- No route, validation, territory, score or event behavior is implemented.

\- No unexpected production file is changed.

\- `gofmt` passes.

\- `go test ./...` passes.

\- `make backend-check` passes.

\- `git diff --check` passes.

\- `PROJECT\_STATUS.md` is updated only after all checks pass.



\---



\# Completion Report



Return:



```text

Task ID: M1-005



Changed files:

\- ...



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risks:

\- EventPayload concrete value contracts are intentionally deferred to M1-017.

\- ScoreChangeReason canonical values are not defined by Calculator Rules v1.0 and were not invented.

```

