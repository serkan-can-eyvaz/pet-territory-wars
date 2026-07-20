\# M0-005 — Docker Development Environment



Status: Completed



\## Goal



Yerel geliştirme için tekrar üretilebilir Docker geliştirme ortamını hazırlamak.



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Pet Territory Wars — Project Conventions v1.1.pdf

\- docs/roadmap/Pet Territory Wars — Milestone 0 Repository \& Development Environment.pdf



Do not read unrelated documents.



\---



\## Do



\- Docker Compose geliştirme ortamını oluştur.

\- Mimari dokümanlarda tanımlanan servisleri ekle.

\- Gerekli volume ve network yapılandırmalarını oluştur.

\- Backend ve PostgreSQL servislerini geliştirme amacıyla yapılandır.

\- Gerekliyse `.env.example` ile uyumlu environment tanımlarını kullan.

\- Yapılandırma dosyalarını doğrula.

\- Tüm Done koşulları geçmeden `PROJECT\_STATUS.md` dosyasını güncelleme.



\---



\## Don't



\- Uygulama iş mantığı yazma.

\- HTTP endpoint ekleme.

\- PostgreSQL migration çalıştırma.

\- Seed data ekleme.

\- RabbitMQ ekleme.

\- Redis ekleme.

\- Kubernetes dosyaları oluşturma.

\- Production deployment hazırlama.

\- CI/CD pipeline oluşturma.

\- Flutter koduna dokunma.

\- Sonraki tasklara geçme.



\---



\## Done



\- \[ ] Docker Compose dosyası oluşturuldu.

\- \[ ] Geliştirme servisleri tanımlandı.

\- \[ ] Volume yapılandırmaları mevcut.

\- \[ ] Network yapılandırması mevcut.

\- \[ ] Environment değişkenleri yapılandırıldı.

\- \[ ] `docker compose config` başarılı.

\- \[ ] Scope dışına çıkılmadı.

\- \[ ] `PROJECT\_STATUS.md` güncellendi.



\---



\## Commands



Run from the repository root:



```bash

docker compose config

```



If available:



```bash

docker compose up --build

docker compose down

```



Use the second command only to verify the development environment if it is required by the architecture document.



\---



\## Completion Update



Task tamamlanınca:



\- `PROJECT\_STATUS.md` içinde M0-005 Completed olarak işaretlenecek.

\- Sonraki aktif görev M0-006 olacak.

\- Başka task otomatik başlatılmayacak.



\---



\## Output



```text

Task ID: M0-005



Changed files:

\- ...



Checks run:

\- ...



Remaining risk:

\- None

```

