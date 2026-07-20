\# M0-007 — Health Endpoints



Status: Completed



\## Goal



Backend servisine temel health endpoint'lerini eklemek.



Bu görev yalnızca uygulamanın canlılık (liveness) ve hazır olma (readiness)

kontrollerini kapsar.



\---



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Pet Territory Wars — Project Conventions v1.1.pdf

\- docs/architecture/Pet Territory Wars — Backend Architecture v1.1.pdf

\- docs/roadmap/Pet Territory Wars — Milestone 0 Repository \& Development Environment.pdf



Do not read unrelated documents.



\---



\## Do



\- HTTP health endpointlerini oluştur.

\- En az liveness ve readiness endpointlerini uygula.

\- Response formatını mimari belgeye uygun tut.

\- Readiness kontrolünde yalnızca temel bağımlılık doğrulamalarını yap.

\- Handler'ları infrastructure katmanında konumlandır.

\- Kodu formatla.

\- Tüm doğrulamalar geçmeden PROJECT\_STATUS.md dosyasını güncelleme.



\---



\## Don't



\- Authentication ekleme.

\- Middleware ekleme.

\- Structured logging ekleme.

\- Metrics ekleme.

\- OpenTelemetry ekleme.

\- Swagger ekleme.

\- Domain servislerini çağırma.

\- Repository kullanma.

\- İş mantığı ekleme.

\- Docker değiştirme.

\- Flutter koduna dokunma.

\- Sonraki göreve geçme.



\---



\## Done



\- \[ ] Liveness endpointi eklendi.

\- \[ ] Readiness endpointi eklendi.

\- \[ ] Response formatı tutarlı.

\- \[ ] gofmt başarılı.

\- \[ ] go vet ./... başarılı.

\- \[ ] go test ./... başarılı.

\- \[ ] go build ./... başarılı.

\- \[ ] Scope dışına çıkılmadı.

\- \[ ] PROJECT\_STATUS.md güncellendi.



\---



\## Commands



Run from the repository root:



```bash

gofmt -w .

go vet ./...

go test ./...

go build ./...

```



\---



\## Completion Update



Task tamamlanınca:



\- PROJECT\_STATUS.md içinde M0-007 Completed olarak işaretlenecek.

\- Sonraki aktif görev M0-008 — Structured Logging olacak.

\- Başka task otomatik başlatılmayacak.



\---



\## Output



```text

Task ID: M0-007



Changed files:

\- ...



Checks run:

\- ...



Remaining risk:

\- None

```

