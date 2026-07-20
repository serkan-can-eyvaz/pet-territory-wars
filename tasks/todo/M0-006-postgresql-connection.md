\# M0-006 — PostgreSQL Connection



Status: Completed



\## Goal



PostgreSQL bağlantı altyapısını oluşturmak.



Bu görev yalnızca veritabanı bağlantısının kurulmasını kapsar.

Repository, migration, transaction yönetimi veya iş mantığı bu görevin kapsamı değildir.



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



\- `internal/infrastructure/database` paketini oluştur.

\- PostgreSQL bağlantısını pgx kullanarak kur.

\- Mevcut configuration paketini kullan.

\- Bağlantı oluşturma fonksiyonlarını ekle.

\- Bağlantıyı güvenli şekilde kapatacak yardımcı fonksiyon ekle.

\- Connection pool yapılandırmasını mimari belgede belirtilen değerlerle sınırla.

\- Hata durumlarını açık şekilde döndür.

\- Birim test gerekmiyorsa ekleme.

\- Tüm doğrulamalar geçmeden `PROJECT\_STATUS.md` güncelleme.



\---



\## Don't



\- Repository yazma.

\- SQL sorgusu yazma.

\- Migration çalıştırma.

\- Transaction helper yazma.

\- Outbox implementasyonu yapma.

\- Domain koduna dokunma.

\- HTTP server ekleme.

\- Logging ekleme.

\- Docker değiştirme.

\- Flutter koduna dokunma.

\- Sonraki tasklara geçme.



\---



\## Done



\- \[ ] PostgreSQL bağlantı paketi oluşturuldu.

\- \[ ] Configuration sistemi kullanılıyor.

\- \[ ] Connection pool oluşturuluyor.

\- \[ ] Close fonksiyonu mevcut.

\- \[ ] Kod formatlandı.

\- \[ ] go mod tidy başarılı.

\- \[ ] go test ./... başarılı.

\- \[ ] go build ./... başarılı.

\- \[ ] Scope dışına çıkılmadı.

\- \[ ] PROJECT\_STATUS.md güncellendi.



\---



\## Commands



Run from the repository root:



```bash

go mod tidy

gofmt -w .

go test ./...

go build ./...

```



\---



\## Completion Update



Task tamamlanınca:



\- `PROJECT\_STATUS.md` içinde M0-006 Completed olarak işaretlenecek.

\- Sonraki aktif görev M0-007 olacak.

\- Başka task otomatik başlatılmayacak.



\---



\## Output



```text

Task ID: M0-006



Changed files:

\- ...



Checks run:

\- ...



Remaining risk:

\- None

```

