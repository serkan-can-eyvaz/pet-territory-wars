\# M1-007 — Walk Validation



\## Goal



M1-006 tarafından üretilen RouteSegment listesinden walk seviyesindeki doğrulamayı gerçekleştirmek.



Bu görev yalnızca walk seviyesindeki geçerlilik kararını, reddedilme nedenlerini ve temel walk metriklerini hesaplar.



\## Read



Sadece aşağıdaki doküman ve dosyaları oku:



\- Calculator Rules v1.0

&#x20; - Walk Validation

&#x20; - Walk Metrics

&#x20; - Valid Route Ratio

&#x20; - Rejection Reasons

\- Project Conventions v1.1

\- backend/internal/gameengine/route\_segment.go

\- backend/internal/gameengine/result.go

\- backend/internal/gameengine/rules.go



Başka doküman okuma.



\---



\## Implement



Üret:



\- WalkValidationStatus

\- WalkRejectionReason

\- ValidationResult hesaplaması

\- WalkMetrics hesaplaması

\- valid route ratio hesaplaması

\- walk acceptance kararı



Kurallar Calculator Rules v1.0 ile birebir uyumlu olmalıdır.



\---



\## Rules



\- Walk metrikleri yalnızca RouteSegment listesinden hesaplanacaktır.

\- Geçersiz segmentler valid distance hesabına dahil edilmeyecektir.

\- Total distance tüm segmentlerden hesaplanacaktır.

\- Valid duration yalnızca geçerli segmentlerden hesaplanacaktır.

\- Valid route ratio Calculator Rules formülüyle hesaplanacaktır.

\- Walk rejection reason yalnızca normatif sebeplerden üretilecektir.

\- Aynı walk için deterministik sonuç üretilecektir.



\---



\## Do Not



Eklenmeyecek:



\- Route interpolation

\- H3

\- Hex üretimi

\- Territory

\- Score hesaplama

\- Event üretimi

\- ResolveWalk

\- JSON/ORM tag

\- Public API

\- Third-party dependency



\---



\## Tests



Odaklı testler:



\- valid walk

\- invalid walk

\- rejection reasons

\- valid distance

\- total distance

\- valid duration

\- valid route ratio

\- eşik değerleri

\- deterministik sonuçlar



\---



\## Files



Beklenen değişiklikler:



\- backend/internal/gameengine/walk\_validation.go

\- backend/internal/gameengine/walk\_validation\_test.go



Tüm kontroller geçerse:



\- PROJECT\_STATUS.md



\---



\## Validation



Çalıştır:



```bash

gofmt -w internal/gameengine/walk\_validation.go internal/gameengine/walk\_validation\_test.go

go test ./...

```



Repository root:



```bash

make backend-check

git diff --check

```



PROJECT\_STATUS.md yalnızca tüm kontroller başarılı olduktan sonra güncellenecektir.



\---



\## Completion Report



```text

Task ID: M1-007



Changed files:

\- backend/internal/gameengine/walk\_validation.go

\- backend/internal/gameengine/walk\_validation\_test.go

\- PROJECT\_STATUS.md



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- Route interpolation, H3, territory, score ve event üretimi sonraki görevlere bırakılmıştır.

```

