\# M0-002 — Backend Bootstrap



Status: TODO



\## Goal



Go backend projesinin temel, derlenebilir ve test edilebilir başlangıç yapısını oluşturmak.



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Pet Territory Wars — Project Conventions v1.1.pdf

\- docs/roadmap/Pet Territory Wars — Milestone 0 Repository \& Development Environment.pdf



Do not read unrelated documents.



\## Do



\- `backend` klasöründe Go modülünü başlat.

\- Modül yolunu repository adresine uygun belirle.

\- Mimari dokümanlarda belirtilen temel backend klasör yapısını oluştur.

\- Backend'in hatasız derlenebildiğini doğrula.

\- Go test komutunun hatasız çalışmasını sağla.

\- Kullanılmayan örnek kod veya gereksiz bağımlılık ekleme.

\- Task tamamlandığında `PROJECT\_STATUS.md` dosyasını güncelle.



\## Don't



\- Configuration sistemi oluşturma.

\- Logging sistemi oluşturma.

\- HTTP endpoint ekleme.

\- PostgreSQL bağlantısı ekleme.

\- Docker yapılandırması ekleme.

\- Migration sistemi ekleme.

\- İş mantığı veya oyun motoru kodu yazma.

\- Gereksiz üçüncü taraf bağımlılık ekleme.

\- Sonraki tasklara geçme.



\## Done



\- \[ ] `backend/go.mod` mevcut.

\- \[ ] Go modül yolu doğru.

\- \[ ] Gerekli temel backend klasörleri mevcut.

\- \[ ] Go kaynak kodları formatlandı.

\- \[ ] `go test ./...` başarılı.

\- \[ ] `go build ./...` başarılı.

\- \[ ] Kapsam dışı özellik eklenmedi.

\- \[ ] `PROJECT\_STATUS.md` güncellendi.



\## Commands



Run inside the `backend` directory:



```bash

go mod tidy

gofmt -w .

go test ./...

go build ./...

