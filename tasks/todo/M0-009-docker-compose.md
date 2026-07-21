\# M0-009 — Docker Compose



Status: Completed



\## Goal



Backend geliştirme ortamını Docker Compose ile ayağa kaldırmak.



Bu görev yalnızca development compose ortamını kapsar.



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



\- Repository root altında compose.yaml dosyasını oluştur.

\- Backend servisini ekle.

\- PostgreSQL servisini ekle.

\- Backend servisini M0-008'de oluşturulan backend/Dockerfile kullanarak build et.

\- Backend için gerekli environment değişkenlerini tanımla.

\- PostgreSQL için gerekli environment değişkenlerini tanımla.

\- Backend ile PostgreSQL aynı Docker network üzerinde çalışsın.

\- PostgreSQL için named volume kullan.

\- docker compose config komutunun başarılı olduğunu doğrula.

\- docker compose up ile servislerin başlayabildiğini doğrula.

\- Tüm doğrulamalar geçmeden PROJECT\_STATUS.md dosyasını güncelleme.



\---



\## Don't



\- Migration container ekleme.

\- Healthcheck ekleme.

\- CI değiştirme.

\- Kubernetes dosyaları ekleme.

\- Nginx ekleme.

\- Hot reload ekleme.

\- Monitoring ekleme.

\- Logging sistemi ekleme.

\- Production compose oluşturma.

\- Business logic değiştirme.

\- Flutter koduna dokunma.

\- Sonraki göreve geçme.



\---



\## Done



\- \[ ] compose.yaml oluşturuldu.

\- \[ ] Backend servisi eklendi.

\- \[ ] PostgreSQL servisi eklendi.

\- \[ ] Named volume kullanılıyor.

\- \[ ] Backend Dockerfile üzerinden build ediliyor.

\- \[ ] docker compose config başarılı.

\- \[ ] docker compose up başarılı.

\- \[ ] Scope dışına çıkılmadı.

\- \[ ] PROJECT\_STATUS.md güncellendi.



\---



\## Commands



Repository root:



```bash

docker compose config



docker compose up --build

```



\---



\## Completion Update



Task tamamlandığında:



\- PROJECT\_STATUS.md içinde M0-009 Completed olarak işaretlenecek.

\- Sonraki aktif görev M0-010 — Migration System olacak.

\- Başka görev otomatik başlatılmayacak.



\---



\## Output



```text

Task ID: M0-009



Changed files:

\- compose.yaml

\- ...



Checks run:

\- docker compose config

\- docker compose up --build



Remaining risk:

\- None

```

