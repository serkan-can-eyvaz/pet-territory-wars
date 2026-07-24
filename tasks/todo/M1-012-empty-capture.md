\# M1-012 — Empty Capture



\## Goal



Owner'ı olmayan (OwnerID == nil) bir hex'in ilk kez ele geçirilmesi

durumunu deterministik olarak hesaplamak.



Bu görev yalnızca empty capture uygular.



Defense, attack, ownership transfer, score, event ve ResolveWalk

kapsam dışındadır.



\---



\## Read



Önce:



\- AGENTS.md

\- PROJECT\_STATUS.md

\- bu görev dosyası



Ardından yalnızca:



\- Calculator Rules v1.0

&#x20; - Empty Capture

&#x20; - Processing Order

&#x20; - Error Model

&#x20; - Edge Cases

\- Project Conventions v1.1

\- backend/internal/gameengine/input.go

\- backend/internal/gameengine/result.go

\- backend/internal/gameengine/rules.go



Başka doküman okuma.



\---



\## Implement



Package-private fonksiyon ekle:



captureEmptyHex(...)



Mevcut HexState, WalkInput ve TerritoryRules tipleri yeniden kullanılacaktır.



Yeni public API veya yeni domain modeli oluşturulmayacaktır.



\---



\## Rules



\- Yalnızca OwnerID == nil olan hex işlenecektir.

\- OwnerID bulunan hex değiştirilmeyecektir.

\- Capture şartları yalnızca Calculator Rules'a göre değerlendirilecektir.

\- Girdi nesneleri değiştirilmeyecektir.

\- Sonuç yeni state olarak üretilecektir.

\- Sistem saati okunmayacaktır.

\- time.Now() kullanılmayacaktır.

\- Persistence yapılmayacaktır.



\---



\## Do Not



Eklenmeyecek:



\- owner defense

\- enemy attack

\- ownership transfer

\- score

\- domain event

\- persistence

\- database

\- validation

\- ResolveWalk

\- yeni public API



\---



\## Tests



Package içi testler:



\- owner olmayan hex

\- owner bulunan hex

\- capture başarılı

\- capture başarısız

\- deterministik sonuç

\- girdi değişmezliği



\---



\## Files



Beklenen değişiklikler:



\- backend/internal/gameengine/empty\_capture.go

\- backend/internal/gameengine/empty\_capture\_test.go



Tüm kontroller başarılı olursa:



\- PROJECT\_STATUS.md



\---



\## Validation



```bash

gofmt -w internal/gameengine/empty\_capture.go internal/gameengine/empty\_capture\_test.go

go test ./...

```



Repository root:



```bash

make backend-check

git diff --check

```



PROJECT\_STATUS.md yalnızca tüm kontroller başarılı olduktan sonra güncellenecektir.



\---



\## Pass Criteria



\- Calculator Rules'taki Empty Capture kuralları eksiksiz uygulanmıştır.

\- Sonuç deterministiktir.

\- Girdi değişmez.

\- Tüm testler geçmektedir.



\---



\## Completion Report



```text

Task ID: M1-012



Changed files:

\- backend/internal/gameengine/empty\_capture.go

\- backend/internal/gameengine/empty\_capture\_test.go

\- PROJECT\_STATUS.md



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- Owner Defense (M1-013) sonraki görevde uygulanacaktır.

```

