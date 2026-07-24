\# M1-006 — Route Segment Calculation



\## Goal



Ardışık `LocationPoint` değerlerinden deterministic `RouteSegment` modelleri üretmek ve her segment için mesafe, süre, hız ile segment seviyesindeki geçerlilik sonucunu hesaplamak.



Bu görev yalnızca segment oluşturma ve segment seviyesindeki hesaplamaları içerir.



Walk-level metrics, walk acceptance kararı, interpolation, H3, territory, score veya event üretimi bu görevin kapsamında değildir.



\---



\## Roadmap Position



Current task:



\- M1-006 — Route Segment Calculation



Next task:



\- M1-007 — Walk Validation



Roadmap numaraları değiştirilmemelidir.



\---



\## Read



Yalnızca:



\- Calculator Rules v1.0

&#x20; - Section 7 — Engine Input Contract

&#x20; - Section 8 — Normative Execution Order

&#x20; - Section 9 — GPS Point and Segment Validation

&#x20; - Section 9.1 — Formulas

&#x20; - Section 9.2 — Invalid Reason Priority

&#x20; - Section 9.3 — Zero-Distance Segment

&#x20; - Section 25 — Relevant Edge Cases and Invariants

\- Project Conventions v1.1

\- Existing files:

&#x20; - backend/internal/gameengine/input.go

&#x20; - backend/internal/gameengine/rules.go

&#x20; - backend/internal/gameengine/rules\_validation.go



Do not read unrelated documents.



\---



\## Normative Scope



Calculator Rules v1.0 is the normative source.



This task implements:



1\. `RouteSegment`

2\. `SegmentInvalidReason`

3\. Ordered route-to-segment construction

4\. Deterministic geographic distance calculation

5\. Duration calculation

6\. Speed calculation

7\. Segment-level validation excluding city-boundary geometry evaluation

8\. Invalid-reason precedence



This task does not implement:



\- structural input validation

\- city-boundary geometry calculation

\- walk-level metrics

\- walk validity status

\- rejection reasons

\- interpolation

\- H3 conversion

\- territory behavior



\---



\# Required Types



\## 1. SegmentInvalidReason



Define:



```go

type SegmentInvalidReason string

```



Canonical values:



```go

const (

&#x20;   SegmentInvalidReasonNone                SegmentInvalidReason = ""

&#x20;   SegmentInvalidReasonNonPositiveDuration SegmentInvalidReason = "NON\_POSITIVE\_DURATION"

&#x20;   SegmentInvalidReasonMockLocation        SegmentInvalidReason = "MOCK\_LOCATION"

&#x20;   SegmentInvalidReasonLowAccuracy         SegmentInvalidReason = "LOW\_ACCURACY"

&#x20;   SegmentInvalidReasonLocationJump        SegmentInvalidReason = "LOCATION\_JUMP"

&#x20;   SegmentInvalidReasonImpossibleSpeed     SegmentInvalidReason = "IMPOSSIBLE\_SPEED"

&#x20;   SegmentInvalidReasonOutsideCity         SegmentInvalidReason = "OUTSIDE\_CITY"

)

```



Rules:



\- Empty string represents no invalid reason.

\- Do not add another reason.

\- Do not add parsing, `String()`, validation or marshal helpers.

\- `OUTSIDE\_CITY` must be declared for the complete normative enum, but this task must not produce it because city-boundary geometry evaluation is not implemented here.



\---



\## 2. RouteSegment



Define exactly:



```go

type RouteSegment struct {

&#x20;   From           LocationPoint

&#x20;   To             LocationPoint

&#x20;   DistanceMeters float64

&#x20;   Duration       time.Duration

&#x20;   SpeedMPS       float64

&#x20;   IsValid        bool

&#x20;   InvalidReason  SegmentInvalidReason

}

```



Do not change field names, order or types.



\---



\# Required Functions



Production functions must remain package-private.



\## 1. Route Segment Construction



Implement:



```go

func buildRouteSegments(

&#x20;   route \[]LocationPoint,

&#x20;   rules WalkRules,

) \[]RouteSegment

```



Behavior:



\- Input route order must be used as provided.

\- Structural input sorting must not be performed.

\- For `len(route) < 2`, return an empty non-nil or nil slice consistently.

\- For `n` points, produce exactly `n-1` segments.

\- Segment `i` must use:

&#x20; - `From = route\[i]`

&#x20; - `To = route\[i+1]`

\- Output order must match input adjacency order.

\- Input slice and points must not be mutated.

\- Each segment must be calculated through the single-segment calculation function.



Do not validate duplicate sequence values or timestamp ordering at route level. Structural validation belongs to another layer.



\---



\## 2. Single Segment Calculation



Implement:



```go

func calculateRouteSegment(

&#x20;   from LocationPoint,

&#x20;   to LocationPoint,

&#x20;   rules WalkRules,

) RouteSegment

```



The function must:



1\. Preserve `from` and `to`.

2\. Calculate deterministic geographic distance.

3\. Calculate duration as:



```go

duration := to.RecordedAt.Sub(from.RecordedAt)

```



4\. Calculate speed according to the rules below.

5\. Assign exactly one invalid reason using canonical precedence.

6\. Set `IsValid` consistently with `InvalidReason`.



\---



\## 3. Geographic Distance



Implement one package-private deterministic helper:



```go

func geoDistanceMeters(

&#x20;   fromLatitude float64,

&#x20;   fromLongitude float64,

&#x20;   toLatitude float64,

&#x20;   toLongitude float64,

) float64

```



Use the Haversine formula.



Use this Earth radius constant:



```go

const earthRadiusMeters = 6371000.0

```



Required formula:



```text

lat1 = radians(fromLatitude)

lat2 = radians(toLatitude)

deltaLat = radians(toLatitude - fromLatitude)

deltaLon = radians(toLongitude - fromLongitude)



a =

&#x20;   sin(deltaLat/2)^2 +

&#x20;   cos(lat1) \* cos(lat2) \* sin(deltaLon/2)^2



c = 2 \* atan2(sqrt(a), sqrt(1-a))



distance = earthRadiusMeters \* c

```



Numerical requirement:



\- Clamp `a` to `\[0, 1]` before square roots to avoid floating-point drift.

\- Do not round the returned distance.

\- Do not introduce configurable Earth radius.

\- Do not use a third-party geospatial library.

\- Do not use H3 for distance.



\---



\# Duration and Speed Rules



\## Positive Duration



When:



```go

duration > 0

```



calculate:



```go

speedMPS := distanceMeters / duration.Seconds()

```



\## Non-Positive Duration



When:



```go

duration <= 0

```



the segment must be invalid with:



```text

NON\_POSITIVE\_DURATION

```



To prevent invalid floating-point division:



```go

SpeedMPS = 0

```



for non-positive duration.



Do not produce `NaN` or `Inf`.



\## Zero-Distance Segment



When:



\- duration is positive

\- distance is zero

\- no higher-priority invalid condition applies



the segment may be valid.



Required values:



```text

DistanceMeters = 0

SpeedMPS = 0

```



Zero distance alone is not an invalid reason.



\---



\# Segment Invalid-Reason Precedence



Exactly one reason may be assigned.



Apply conditions in this order:



1\. `NON\_POSITIVE\_DURATION`

2\. `MOCK\_LOCATION`

3\. `LOW\_ACCURACY`

4\. `LOCATION\_JUMP`

5\. `IMPOSSIBLE\_SPEED`

6\. `OUTSIDE\_CITY`



The first matching reason wins.



\## 1. NON\_POSITIVE\_DURATION



Condition:



```go

duration <= 0

```



\## 2. MOCK\_LOCATION



Condition:



```go

from.IsMockLocation || to.IsMockLocation

```



\## 3. LOW\_ACCURACY



Condition:



```go

from.AccuracyMeters > rules.MaxAccuracyMeters ||

&#x20;   to.AccuracyMeters > rules.MaxAccuracyMeters

```



Exact threshold is allowed:



```text

AccuracyMeters == MaxAccuracyMeters

```



must not be invalid for low accuracy.



Negative, `NaN` or infinite accuracy values are structural input errors and must not be handled in this task.



\## 4. LOCATION\_JUMP



Condition:



```go

distanceMeters > rules.MaxJumpMeters

```



Exact threshold is allowed:



```text

DistanceMeters == MaxJumpMeters

```



must not be invalid for location jump.



\## 5. IMPOSSIBLE\_SPEED



Condition:



```go

speedMPS > rules.MaxSpeedMPS

```



Exact threshold is allowed:



```text

SpeedMPS == MaxSpeedMPS

```



must not be invalid for impossible speed.



\## 6. OUTSIDE\_CITY



`OUTSIDE\_CITY` belongs to the normative reason set, but this task must not evaluate city-boundary geometry.



Therefore:



\- declare the constant

\- do not produce it in `calculateRouteSegment`

\- do not add `CityBoundary` as a function parameter

\- do not implement point-in-polygon

\- do not invent partial-segment clipping behavior



Boundary filtering will be integrated when the engine flow reaches the boundary-evaluation stage.



\---



\# Validity Assignment



After selecting the invalid reason:



```go

IsValid = InvalidReason == SegmentInvalidReasonNone

```



No independent validity condition may contradict the reason.



Examples:



```text

InvalidReason == ""                     => IsValid true

InvalidReason == "MOCK\_LOCATION"        => IsValid false

InvalidReason == "IMPOSSIBLE\_SPEED"     => IsValid false

```



\---



\# Architecture Decisions



\## Rules Input



The segment calculation receives `WalkRules`, not the complete `RuleSet`.



Reason:



\- only walk-level segment thresholds are required

\- the function must not depend on Movement or Territory rules



\## No Error Return



The calculation functions do not return Go `error`.



Reason:



\- inputs are assumed structurally validated before this stage

\- segment rejection is represented by `InvalidReason`

\- structural input errors belong to input-validation behavior



\## Package Visibility



Functions must remain package-private.



Do not export:



\- `BuildRouteSegments`

\- `CalculateRouteSegment`

\- `GeoDistanceMeters`



The public engine entry point remains `ResolveWalk` in a later task.



\---



\# Files



Create:



```text

backend/internal/gameengine/route\_segment.go

backend/internal/gameengine/route\_segment\_test.go

```



Update only after all checks pass:



```text

PROJECT\_STATUS.md

```



Do not modify:



```text

backend/internal/gameengine/input.go

backend/internal/gameengine/result.go

backend/internal/gameengine/hex\_change.go

backend/internal/gameengine/rules.go

```



unless compilation exposes a direct defect that blocks this task. Report such a conflict before changing those files.



\---



\# Implementation Requirements



\- Use `package gameengine`.

\- Use only the Go standard library.

\- Use `math` for Haversine calculations.

\- Use `time.Duration` in `RouteSegment`.

\- Do not mutate route input.

\- Do not sort route points.

\- Do not create constructors.

\- Do not create builders.

\- Do not add interfaces.

\- Do not add configurable distance strategies.

\- Do not add JSON, database or ORM tags.

\- Do not add logging.

\- Do not add metrics instrumentation.

\- Do not add third-party dependencies.

\- Do not round calculated values.

\- Do not use current time.

\- Do not use randomness.



\---



\# Required Tests



Create focused tests in:



```text

backend/internal/gameengine/route\_segment\_test.go

```



\## 1. Canonical Invalid Reasons



Verify exact values:



\- `""`

\- `NON\_POSITIVE\_DURATION`

\- `MOCK\_LOCATION`

\- `LOW\_ACCURACY`

\- `LOCATION\_JUMP`

\- `IMPOSSIBLE\_SPEED`

\- `OUTSIDE\_CITY`



\## 2. Route Construction



Verify:



\- zero points produce zero segments

\- one point produces zero segments

\- two points produce one segment

\- three points produce two ordered segments

\- `From` and `To` values are preserved

\- input route remains unchanged



Do not test sorting.



\## 3. Deterministic Distance



Verify:



\- identical coordinates produce zero distance

\- a known short coordinate pair produces an expected approximate distance

\- reversing endpoints produces the same distance

\- repeated calls produce the same result



Use tolerance for floating-point comparison.



Do not require exact binary equality against an externally calculated decimal.



\## 4. Duration and Speed



Verify:



\- positive duration is preserved

\- speed equals distance divided by duration seconds

\- zero distance with positive duration produces speed zero

\- zero duration produces speed zero and `NON\_POSITIVE\_DURATION`

\- negative duration produces speed zero and `NON\_POSITIVE\_DURATION`



\## 5. Valid Segment



Create a segment that satisfies all rules and verify:



```text

IsValid == true

InvalidReason == ""

```



\## 6. Mock Location



Verify:



\- mock `From`

\- mock `To`



both produce:



```text

MOCK\_LOCATION

```



when duration is positive.



\## 7. Low Accuracy



Verify:



\- `From.AccuracyMeters > MaxAccuracyMeters`

\- `To.AccuracyMeters > MaxAccuracyMeters`



produce:



```text

LOW\_ACCURACY

```



Verify exact threshold remains allowed.



\## 8. Location Jump



Verify:



\- distance above `MaxJumpMeters` produces `LOCATION\_JUMP`

\- exact threshold does not produce `LOCATION\_JUMP`



Use test values or a rule threshold derived from calculated distance to avoid brittle coordinate assumptions.



\## 9. Impossible Speed



Verify:



\- speed above `MaxSpeedMPS` produces `IMPOSSIBLE\_SPEED`

\- exact threshold is allowed



Use duration or threshold values derived from the calculated distance.



\## 10. Precedence



Table-driven tests must verify at least:



\- non-positive duration beats mock location

\- mock location beats low accuracy

\- low accuracy beats location jump

\- location jump beats impossible speed



Only the highest-priority reason must be returned.



\## 11. OUTSIDE\_CITY Deferral



Verify only that the canonical constant exists.



Do not write a test expecting segment calculation to produce `OUTSIDE\_CITY`.



\---



\# Test Restrictions



Tests must not:



\- test walk status

\- calculate valid route ratio

\- aggregate walk duration or distance

\- use city-boundary geometry

\- use H3

\- test interpolation

\- test territory outcomes

\- test score changes

\- test domain events

\- test `ResolveWalk`

\- inspect struct tags with reflection

\- require third-party packages



\---



\# Do Not



\- Implement M1-007 Walk Validation.

\- Add `ValidationResult` calculation.

\- Produce `WalkResolutionStatus`.

\- Calculate aggregate walk metrics.

\- Produce rejection reasons.

\- Sort route points.

\- Validate IDs.

\- Validate latitude or longitude ranges.

\- Validate duplicate sequence values.

\- Validate route chronology as a technical error.

\- Implement city-boundary geometry.

\- Clip segments to a city boundary.

\- Implement interpolation.

\- Implement H3 conversion.

\- Implement hex aggregation.

\- Implement territory logic.

\- Implement score logic.

\- Implement event generation.

\- Implement `ResolveWalk`.

\- Add public API functions.

\- Add third-party libraries.



\---



\# Expected Changed Files



Expected:



```text

backend/internal/gameengine/route\_segment.go

backend/internal/gameengine/route\_segment\_test.go

PROJECT\_STATUS.md

```



Task file:



```text

tasks/todo/M1-006-route-segment-calculation.md

```



No other production file should change.



\---



\# Validation



From `backend`:



```bash

gofmt -w internal/gameengine/route\_segment.go internal/gameengine/route\_segment\_test.go

go test ./...

```



From repository root:



```bash

make backend-check

git diff --check

```



Every command must pass.



If any check fails:



1\. Fix only the issue within M1-006 scope.

2\. Run the complete validation sequence again.

3\. Do not update `PROJECT\_STATUS.md` until all checks pass.



\---



\# Pass Criteria



The task is complete only when:



\- Roadmap remains unchanged.

\- `RouteSegment` matches Calculator Rules Section 9.

\- Canonical invalid reasons are exact.

\- Adjacent route points produce ordered segments.

\- Distance uses deterministic Haversine calculation.

\- Duration uses `To.RecordedAt - From.RecordedAt`.

\- Positive-duration speed is calculated correctly.

\- Non-positive duration does not create `NaN` or `Inf`.

\- Zero-distance positive-duration segments can remain valid.

\- Invalid-reason precedence is exact.

\- Only one invalid reason is assigned.

\- `OUTSIDE\_CITY` is declared but not calculated.

\- Input route is not mutated.

\- No walk-level behavior is implemented.

\- No boundary, interpolation, H3 or territory behavior is implemented.

\- `gofmt` passes.

\- `go test ./...` passes.

\- `make backend-check` passes.

\- `git diff --check` passes.

\- `PROJECT\_STATUS.md` is updated only after every check passes.



\---



\# Completion Report



Return:



```text

Task ID: M1-006



Changed files:

\- ...



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- OUTSIDE\_CITY evaluation is intentionally deferred because city-boundary segment evaluation is not part of M1-006.

```

