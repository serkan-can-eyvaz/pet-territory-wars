\# M0-004 — Configuration System



Status: TODO



\## Goal



Backend için merkezi, genişletilebilir ve test edilebilir configuration altyapısını oluşturmak.



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Pet Territory Wars — Project Conventions v1.1.pdf

\- docs/roadmap/Pet Territory Wars — Milestone 0 Repository \& Development Environment.pdf



Do not read unrelated documents.



\---



\## Do



\- Configuration paketini oluştur.

\- Ortak configuration yapısını tanımla.

\- Environment variable okuma altyapısını oluştur.

\- Varsayılan configuration yükleme mekanizmasını oluştur.

\- Configuration doğrulama mekanizmasını ekle.

\- Kodun derlenebilir ve test edilebilir olmasını sağla.

\- Tüm Done koşulları geçmeden PROJECT\_STATUS.md dosyasını güncelleme.



\---



\## Don't



\- HTTP server oluşturma.

\- Logger oluşturma.

\- PostgreSQL bağlantısı kurma.

\- Redis ekleme.

\- RabbitMQ ekleme.

\- Dependency Injection sistemi yazma.

\- Repository katmanı oluşturma.

\- Service katmanı oluşturma.

\- Migration sistemi yazma.

\- Docker yapılandırmasına dokunma.

\- Business logic yazma.

\- Flutter koduna dokunma.

\- Sonraki tasklara geçme.



\---



\## Done



\- \[ ] Configuration paketi oluşturuldu.

\- \[ ] Environment variable desteği eklendi.

\- \[ ] Varsayılan configuration yükleniyor.

\- \[ ] Configuration doğrulaması mevcut.

\- \[ ] Kod formatlandı.

\- \[ ] go test ./... başarılı.

\- \[ ] go build ./... başarılı.

\- \[ ] Scope dışına çıkılmadı.

\- \[ ] PROJECT\_STATUS.md güncellendi.



\---



\## Commands



Run inside the backend directory:



```bash

go mod tidy

gofmt -w .

go test ./...

go build ./...

```



\---



\## Completion Update



Task tamamlanınca:



\- PROJECT\_STATUS.md içinde M0-004 Completed olarak işaretlenecek.

\- Sonraki aktif görev M0-005 olacak.

\- Başka task otomatik başlatılmayacak.



\---



\## Output



```text

Task ID: M0-004



Changed files:

\- ...



Checks run:

\- ...



Remaining risk:

\- None

```

