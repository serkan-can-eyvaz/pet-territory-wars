\# M0-008 — Structured Logging



Status: Completed




\## Goal



Backend uygulamasına standart kütüphane tabanlı yapılandırılmış (structured) loglama altyapısını eklemek.



Bu görev yalnızca logging altyapısını kapsar.



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



\- backend/internal/infrastructure/logging paketini oluştur.

\- Sadece Go standart kütüphanesindeki log/slog paketini kullan.

\- Logger oluşturma ve konfigürasyonunu infrastructure katmanında tut.

\- JSON log formatı kullan.

\- Log seviyesini configuration paketinden al.

\- Composition root (cmd/api/main.go) logger'ı oluşturup bağımlılıkları buradan wire etsin.

\- HTTP handler, domain ve repository katmanlarını logger implementasyonuna bağlama.

\- Kodu formatla.

\- Tüm doğrulamalar geçmeden PROJECT\_STATUS.md dosyasını güncelleme.



\---



\## Don't



\- Zap ekleme.

\- Logrus ekleme.

\- Zerolog ekleme.

\- OpenTelemetry ekleme.

\- Metrics ekleme.

\- Request logging middleware ekleme.

\- HTTP access log ekleme.

\- Authentication ekleme.

\- Repository değiştirme.

\- Business logic ekleme.

\- Docker değiştirme.

\- Flutter koduna dokunma.

\- Sonraki göreve geçme.



\---



\## Done



\- \[ ] logging paketi oluşturuldu.

\- \[ ] slog tabanlı logger oluşturuldu.

\- \[ ] JSON handler kullanılıyor.

\- \[ ] Log seviyesi configuration üzerinden okunuyor.

\- \[ ] Composition root logger'ı oluşturuyor.

\- \[ ] gofmt başarılı.

\- \[ ] go vet ./... başarılı.

\- \[ ] go test ./... başarılı.

\- \[ ] go build ./... başarılı.

\- \[ ] Scope dışına çıkılmadı.

\- \[ ] PROJECT\_STATUS.md güncellendi.



\---



\## Commands



Run from repository root:



```bash

gofmt -w .

go vet ./...

go test ./...

go build ./...

```



\---



\## Completion Update



Task tamamlanınca:



\- PROJECT\_STATUS.md içinde M0-008 Completed olarak işaretlenecek.

\- Sonraki aktif görev M0-009 — CI Bootstrap olacak.

\- Başka task otomatik başlatılmayacak.



\---



\## Output



```text

Task ID: M0-008



Changed files:

\- ...



Checks run:

\- ...



Remaining risk:

\- None

```

