\# M0-011 — Graceful Shutdown



Status: TODO



\## Goal



API servisinin graceful shutdown davranışını mimariye uygun şekilde

tamamlamak ve doğrulamak.



Bu görev yalnızca mevcut lifecycle'ın eksiklerini tamamlar.

Yeni worker mantığı veya yeni lifecycle altyapısı oluşturulmaz.



\---



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Pet Territory Wars — Backend Architecture v1.1.pdf

\- docs/architecture/Pet Territory Wars — Project Conventions v1.1.pdf

\- docs/roadmap/Pet Territory Wars — Milestone 0 Repository \& Development Environment.pdf



\---



\## Do



\- Mevcut graceful shutdown implementasyonunu incele.

\- Eksikse SIGINT ve SIGTERM desteğini tamamla.

\- HTTP server'ı kontrollü kapat.

\- HTTP\_SHUTDOWN\_TIMEOUT yapılandırmasını kullan.

\- HTTP kapandıktan sonra PostgreSQL pool'unu kapat.

\- Beklenmeyen hataları propagate et.

\- gofmt, go vet, go test ve go build çalıştır.

\- Tüm kontroller geçmeden PROJECT\_STATUS.md güncelleme.



\---



\## Don't



\- Yeni lifecycle framework oluşturma.

\- Worker business logic ekleme.

\- Worker loop yazma.

\- Dockerfile değiştirme.

\- compose.yaml değiştirme.

\- Migration değiştirme.

\- Yeni dependency ekleme.

\- Sonraki göreve geçme.



\---



\## Expected Scope



```text

backend/cmd/api/main.go

backend/cmd/api/\*\_test.go

```



Başarı sonrası:



```text

PROJECT\_STATUS.md

```



\---



\## Validation



```bash

cd backend



gofmt -w .

go vet ./...

go test ./...

go build ./...

```



API'nin SIGINT ve SIGTERM ile kontrollü kapandığını doğrula.



\---



\## Done



\- \[ ] Graceful shutdown mimariye uygun.

\- \[ ] SIGINT çalışıyor.

\- \[ ] SIGTERM çalışıyor.

\- \[ ] HTTP server düzgün kapanıyor.

\- \[ ] PostgreSQL pool kapanıyor.

\- \[ ] gofmt başarılı.

\- \[ ] go vet başarılı.

\- \[ ] go test başarılı.

\- \[ ] go build başarılı.

\- \[ ] PROJECT\_STATUS.md güncellendi.



\---



\## Output



```text

Task ID: M0-011



Changed files:

...



Checks run:

...



Remaining risk:

...

```

