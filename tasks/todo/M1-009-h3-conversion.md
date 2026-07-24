\# M1-009 — H3 Conversion



\## Goal



M1-008 tarafından üretilen interpolasyon noktalarını H3 hex kimliklerine dönüştürmek.



Bu görev yalnızca koordinat → H3 dönüşümünü gerçekleştirir. Hex birleştirme, territory, skor, event veya walk sonucu üretmeyecektir.



\---



\## Read



Sadece aşağıdaki doküman ve dosyaları oku:



\- Calculator Rules v1.0

&#x20; - H3 Conversion

\- Project Conventions v1.1

\- backend/internal/gameengine/route\_interpolation.go

\- backend/internal/gameengine/result.go

\- backend/internal/gameengine/rules.go



Başka doküman okuma.



\---



\## Implement



Yeni package-private fonksiyonlar ekle:



\- convertInterpolationToH3(...)

\- pointToH3(...)



Kurallar Calculator Rules v1.0 ile birebir uyumlu olacaktır.



\---



\## Rules



\- Yalnızca interpolasyon çıktısındaki noktalar işlenecek.

\- Her koordinat tam bir H3 hücresine dönüştürülecek.

\- Aynı koordinat her zaman aynı H3 indeksini üretmelidir.

\- Girdi sırası korunacaktır.

\- Girdi değiştirilmeyecektir.

\- H3 çözünürlüğü RuleSet'ten okunacaktır.

\- Yeni public API eklenmeyecektir.



\---



\## Do Not



Eklenmeyecek:



\- Hex birleştirme

\- Hex ziyaret süreleri

\- Territory davranışı

\- Score hesaplama

\- Event üretimi

\- Walk validation

\- ResolveWalk

\- JSON/ORM tag

\- Third-party wrapper

\- Caching



\---



\## Tests



Odaklı testler:



\- aynı koordinat → aynı H3

\- farklı koordinatlar

\- çıktı sırası korunması

\- boş girdi

\- tek nokta

\- determinizm

\- girdi değişmezliği



\---



\## Files



Beklenen değişiklikler:



\- backend/internal/gameengine/h3\_conversion.go

\- backend/internal/gameengine/h3\_conversion\_test.go



Tüm kontroller geçerse:



\- PROJECT\_STATUS.md



\---



\## Validation



```bash

gofmt -w internal/gameengine/h3\_conversion.go internal/gameengine/h3\_conversion\_test.go

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

Task ID: M1-009



Changed files:

\- backend/internal/gameengine/h3\_conversion.go

\- backend/internal/gameengine/h3\_conversion\_test.go

\- PROJECT\_STATUS.md



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- Hex aggregation, territory, score ve event davranışları sonraki görevlere bırakılmıştır.

```

