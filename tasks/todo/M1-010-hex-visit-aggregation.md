\# M1-010 — Hex Visit Aggregation



\## Goal



M1-009 tarafından üretilen sıralı \[]id.HexID listesinden normatif VisitedHex kayıtlarını oluşturmak.



Bu görev yalnızca ardışık aynı hex ziyaretlerini birleştirir ve ziyaret sırasını korur.



\---



\## Read



Sadece aşağıdaki doküman ve dosyaları oku:



\- Calculator Rules v1.0

&#x20; - Hex Visit Aggregation

&#x20; - VisitedHex

\- Project Conventions v1.1

\- backend/internal/gameengine/result.go

\- backend/internal/gameengine/h3\_conversion.go



Başka doküman okuma.



\---



\## Implement



Yeni package-private fonksiyon ekle:



\- aggregateVisitedHexes(...)



Mevcut result.go içindeki VisitedHex modeli yeniden kullanılacaktır.



Yeni domain modeli oluşturulmayacaktır.



\---



\## Rules



\- Girdi sırası korunacaktır.

\- Girdi değiştirilmeyecektir.

\- Ardışık aynı HexID tek bir VisitedHex olarak birleştirilecektir.

\- Aynı HexID daha sonra tekrar görülürse yeni bir VisitedHex oluşturulacaktır.

\- Çıktı sırası ilk görülme sırasını koruyacaktır.

\- HexID karşılaştırması mevcut id.HexID tipi üzerinden yapılacaktır.

\- Deterministik sonuç üretilecektir.



\---



\## Do Not



Eklenmeyecek:



\- Territory hesaplama

\- Hex ownership

\- Score hesaplama

\- Event üretimi

\- ResolveWalk

\- H3 dönüşümü

\- Validation

\- Public API

\- Third-party dependency



\---



\## Tests



Package içi testler:



\- boş girdi

\- tek hex

\- ardışık tekrar

\- ardışık olmayan tekrar

\- çıktı sırası

\- determinizm

\- girdi değişmezliği



\---



\## Files



Beklenen değişiklikler:



\- backend/internal/gameengine/hex\_visit\_aggregation.go

\- backend/internal/gameengine/hex\_visit\_aggregation\_test.go



Tüm kontroller geçerse:



\- PROJECT\_STATUS.md



\---



\## Validation



```bash

gofmt -w internal/gameengine/hex\_visit\_aggregation.go internal/gameengine/hex\_visit\_aggregation\_test.go

go test ./...

```



Repository root:



```bash

make backend-check

git diff --check

```



PROJECT\_STATUS.md yalnızca tüm kontroller başarılı olduktan sonra güncellenecektir.



\---



\## Completion Report



```text

Task ID: M1-010



Changed files:

\- backend/internal/gameengine/hex\_visit\_aggregation.go

\- backend/internal/gameengine/hex\_visit\_aggregation\_test.go

\- PROJECT\_STATUS.md



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- Territory ownership, score ve event üretimi sonraki görevlere bırakılmıştır.

```

