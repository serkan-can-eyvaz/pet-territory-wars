\# M0-009 — CI Bootstrap



Status: TODO



\## Goal



GitHub Actions tabanlı temel Continuous Integration (CI) pipeline'ını oluşturmak.



Bu görev yalnızca repository doğrulama (validation) pipeline'ını kapsar.



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



\- .github/workflows/ci.yml dosyasını oluştur.

\- Pipeline Ubuntu runner üzerinde çalışsın.

\- Go sürümünü repository'deki go.mod ile uyumlu kur.

\- Backend bağımlılıklarını indir.

\- Aşağıdaki kontrolleri çalıştır:



&#x20; - gofmt (format kontrolü)

&#x20; - go vet ./...

&#x20; - go test ./...

&#x20; - go build ./...



\- Pipeline yalnızca backend dizininde çalışsın.

\- Workflow push ve pull\_request event'lerinde tetiklensin.

\- Pipeline başarısız olursa job başarısız olsun.

\- Tüm doğrulamalar geçmeden PROJECT\_STATUS.md dosyasını güncelleme.



\---



\## Don't



\- Docker build ekleme.

\- Docker image publish etme.

\- Release workflow oluşturma.

\- CD (Continuous Deployment) ekleme.

\- GitHub Release oluşturma.

\- Code coverage raporu ekleme.

\- Linter (golangci-lint vb.) ekleme.

\- Security scan ekleme.

\- Dependabot ekleme.

\- Flutter workflow ekleme.

\- Business logic değiştirme.

\- Repository yapısını değiştirme.

\- Sonraki milestone'a ait kod yazma.



\---



\## Done



\- \[ ] GitHub Actions workflow oluşturuldu.

\- \[ ] push event'i çalışıyor.

\- \[ ] pull\_request event'i çalışıyor.

\- \[ ] Go kuruluyor.

\- \[ ] Dependency restore başarılı.

\- \[ ] gofmt kontrolü çalışıyor.

\- \[ ] go vet ./... çalışıyor.

\- \[ ] go test ./... çalışıyor.

\- \[ ] go build ./... çalışıyor.

\- \[ ] Workflow yalnızca backend üzerinde çalışıyor.

\- \[ ] Scope dışına çıkılmadı.

\- \[ ] PROJECT\_STATUS.md güncellendi.



\---



\## Commands



Local validation:



```bash

cd backend



gofmt -w .

go vet ./...

go test ./...

go build ./...

```



\---



\## Completion Update



Task tamamlandığında:



\- PROJECT\_STATUS.md içinde M0-009 Completed olarak işaretlenecek.

\- Milestone 0 Completed olarak işaretlenecek.

\- Sonraki aktif görev Milestone 1 roadmap'inden başlayacak.

\- Başka görev otomatik başlatılmayacak.



\---



\## Output



```text

Task ID: M0-009



Changed files:

\- .github/workflows/ci.yml

\- ...



Checks run:

\- gofmt

\- go vet

\- go test

\- go build



Remaining risk:

\- None

```

