\# M0-008 — Dockerfile



Status: Complete



\## Goal



Backend uygulaması için production odaklı, çok aşamalı (multi-stage) bir Dockerfile oluşturmak.



Bu görev yalnızca Docker image oluşturmayı kapsar.



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



\- backend/Dockerfile dosyasını oluştur.

\- Multi-stage build kullan.

\- Builder aşamasında resmi Go imajı kullan.

\- Final aşamasında minimal Linux tabanlı bir runtime image kullan.

\- Backend binary'sini builder aşamasında derle.

\- Final image içinde yalnızca çalıştırmak için gerekli dosyaları bulundur.

\- Uygulamayı varsayılan olarak cmd/api üzerinden başlat.

\- Çalışma dizinini (WORKDIR) tanımla.

\- Uygun EXPOSE portunu ekle.

\- Docker build'in başarılı olduğunu doğrula.

\- Tüm doğrulamalar geçmeden PROJECT\_STATUS.md dosyasını güncelleme.



\---



\## Don't



\- Docker Compose ekleme.

\- PostgreSQL servisi ekleme.

\- Migration container ekleme.

\- Healthcheck ekleme.

\- CI workflow değiştirme.

\- Kubernetes dosyaları ekleme.

\- Nginx ekleme.

\- Hot reload ekleme.

\- Docker volume ekleme.

\- Business logic değiştirme.

\- Flutter koduna dokunma.

\- Sonraki göreve geçme.



\---



\## Done



\- \[ ] backend/Dockerfile oluşturuldu.

\- \[ ] Multi-stage build kullanılıyor.

\- \[ ] Binary başarıyla derleniyor.

\- \[ ] Final image minimal.

\- \[ ] Docker build başarılı.

\- \[ ] gofmt başarılı.

\- \[ ] go vet ./...

\- \[ ] go test ./...

\- \[ ] go build ./...

\- \[ ] Scope dışına çıkılmadı.

\- \[ ] PROJECT\_STATUS.md güncellendi.



\---



\## Commands



Repository root:



```bash

cd backend



gofmt -w .

go vet ./...

go test ./...

go build ./...



docker build -t pet-territory-wars-backend .

```



\---



\## Completion Update



Task tamamlandığında:



\- PROJECT\_STATUS.md içinde M0-008 Completed olarak işaretlenecek.

\- Sonraki aktif görev M0-009 — Docker Compose olacak.

\- Başka görev otomatik başlatılmayacak.



\---



\## Output



```text

Task ID: M0-008



Changed files:

\- backend/Dockerfile



Checks run:

\- gofmt

\- go vet

\- go test

\- go build

\- docker build



Remaining risk:

\- None

```

