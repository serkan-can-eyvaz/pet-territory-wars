\# M0-010 — Migration System



Status: Completed



\## Goal



PostgreSQL veritabanı için sürümlendirilebilir, tekrarlanabilir ve geri

alınabilir temel migration sistemini oluşturmak.



Bu görev yalnızca migration altyapısını ve başlangıç migration doğrulamasını

kapsar. Domain tablolarının tamamını tasarlamaz.



\---



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Pet Territory Wars — Project Conventions v1.1.pdf

\- docs/architecture/Pet Territory Wars — Backend Architecture v1.1.pdf

\- docs/architecture/Pet Territory Wars — Infrastructure \& Deployment v1.0.pdf

\- docs/roadmap/Pet Territory Wars — Milestone 0 Repository \& Development Environment.pdf



Do not read unrelated documents.



Architecture documents are normative.



If the documents do not clearly specify the migration tool, file convention,

execution model or required extensions, stop and report the ambiguity instead

of selecting a library or inventing a convention.



\---



\## Do



\- Mimari belgelerde tanımlanan migration aracını ve dosya düzenini kullan.

\- Migration dosyalarını backend altında belgelerde belirtilen konuma yerleştir.

\- Migration'ların sıralı ve değiştirilemez olmasını sağlayan adlandırma

&#x20; düzenini uygula.

\- Her migration için ileri ve geri alma yönünü destekle.

\- İlk migration ile yalnızca mimarinin zorunlu tuttuğu veritabanı

&#x20; extension veya başlangıç altyapısını oluştur.

\- PostGIS extension mimari belgelerde migration tarafından yönetiliyorsa

&#x20; idempotent şekilde etkinleştir.

\- Migration çalıştırıcısını uygulama başlangıcından bağımsız tut.

\- Migration bağlantısı için mevcut DATABASE\_URL sözleşmesini kullan.

\- Migration işleminin hata durumunda sıfır olmayan çıkış koduyla

&#x20; sonlanmasını sağla.

\- Aynı migration setinin boş bir veritabanına uygulanabildiğini doğrula.

\- Uygulanan migration'ların geri alınabildiğini doğrula.

\- Geri alma sonrasında migration'ların yeniden uygulanabildiğini doğrula.

\- Kodu ve migration dosyalarını mevcut proje kurallarına göre formatla.

\- Tüm kontroller geçmeden PROJECT\_STATUS.md dosyasını güncelleme.



\---



\## Don't



\- Migration'ları API başlangıcında otomatik çalıştırma.

\- API, Walk Worker veya Outbox Worker başlangıç kodunu değiştirme.

\- Docker Compose'a migration servisi ekleme.

\- Dockerfile değiştirme.

\- Graceful shutdown ekleme.

\- Domain tablolarının tamamını oluşturma.

\- Repository veya business logic ekleme.

\- Seed data ekleme.

\- Test kullanıcısı veya örnek üretim verisi ekleme.

\- Production deployment veya CI değiştirme.

\- Flutter koduna dokunma.

\- Belgelerde belirtilmeyen migration aracını tahmin ederek seçme.

\- Sonraki göreve geçme.



\---



\## Expected Scope



Belge sözleşmesine bağlı olarak değişmesi beklenen alanlar:



```text

backend/migrations/\*

```



Migration aracı uygulama içinde ayrı bir komut gerektiriyorsa yalnızca

belgelerde tanımlanan giriş noktası eklenebilir. Örneğin:



```text

backend/cmd/migrate/\*

```



Bu yol yalnızca mimari belgeler açıkça bu modeli gerektiriyorsa kullanılmalı;

aksi durumda tahmin edilmemelidir.



Bağımlılık dosyaları yalnızca onaylanan migration aracı yeni bir Go

bağımlılığı gerektiriyorsa değişebilir:



```text

backend/go.mod

backend/go.sum

```



Başarı sonrasında:



```text

PROJECT\_STATUS.md

```



\---



\## Required Migration Behaviour



Migration sistemi aşağıdaki işlemleri desteklemelidir:



```text

up

down

version/status

```



Komut isimleri ve çalıştırma biçimi kullanılan belgeli migration aracına göre

belirlenmelidir.



Migration sistemi:



\- boş veritabanında çalışmalı,

\- daha önce uygulanmış migration'ı tekrar uygulamamalı,

\- hata durumunda açık biçimde başarısız olmalı,

\- veritabanı parolasını veya DATABASE\_URL değerini loglamamalı,

\- migration geçmişini veritabanında takip etmelidir.



\---



\## Validation



Geçici veya local PostgreSQL/PostGIS veritabanı üzerinde şu akışı doğrula:



1\. Veritabanını boş durumda başlat.

2\. Tüm migration'ları uygula.

3\. Migration durumunu/version bilgisini doğrula.

4\. Son migration'ı geri al.

5\. Migration'ı yeniden uygula.

6\. Aynı migration komutunu tekrar çalıştır ve tekrar uygulanmadığını doğrula.



Go kodu veya bağımlılığı değiştiyse ayrıca:



```bash

cd backend



gofmt -w .

go vet ./...

go test ./...

go build ./...

```



Migration aracının belgelerde tanımlanan doğrulama komutlarını da çalıştır.



\---



\## Done



\- \[ ] Migration aracı mimari belgelere göre seçildi.

\- \[ ] Migration dosya düzeni oluşturuldu.

\- \[ ] Sıralı migration adlandırması kullanılıyor.

\- \[ ] İleri migration çalışıyor.

\- \[ ] Geri alma migration'ı çalışıyor.

\- \[ ] Migration geçmişi veritabanında tutuluyor.

\- \[ ] Boş PostgreSQL/PostGIS veritabanında migration başarılı.

\- \[ ] Geri alma sonrası yeniden uygulama başarılı.

\- \[ ] Tekrar çalıştırma uygulanmış migration'ları çoğaltmıyor.

\- \[ ] DATABASE\_URL veya credential bilgisi çıktıya sızmıyor.

\- \[ ] API ve worker başlangıç kodları değiştirilmedi.

\- \[ ] Dockerfile ve compose.yaml değiştirilmedi.

\- \[ ] Gerekliyse gofmt başarılı.

\- \[ ] Gerekliyse go vet ./... başarılı.

\- \[ ] Gerekliyse go test ./... başarılı.

\- \[ ] Gerekliyse go build ./... başarılı.

\- \[ ] Scope dışına çıkılmadı.

\- \[ ] PROJECT\_STATUS.md güncellendi.



\---



\## Completion Update



Task tamamlandığında:



\- PROJECT\_STATUS.md içinde M0-010 Completed olarak işaretlenecek.

\- Aktif görev M0-011 — Graceful Shutdown olacak.

\- Next Task M0-012 — Makefile olacak.

\- Başka görev otomatik başlatılmayacak.



Beklenen üst bölüm:



```text

Task: M0-011 — Graceful Shutdown

Task File: tasks/todo/M0-011-graceful-shutdown.md

Next Task: M0-012 — Makefile

Status: In Progress

```



\---



\## Output



```text

Task ID: M0-010



Architecture contract:

\- Migration tool:

\- Migration location:

\- Execution model:



Changed files:

\- ...



Checks run:

\- migration up:

\- migration status/version:

\- migration down:

\- migration up after down:

\- repeated migration up:

\- gofmt:

\- go vet:

\- go test:

\- go build:



Database validation:

\- PostgreSQL/PostGIS:

\- Applied version:

\- Rollback result:

\- Reapply result:



Remaining risk:

\- None

```

