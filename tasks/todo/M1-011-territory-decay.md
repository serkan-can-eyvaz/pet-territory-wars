\# M1-011 — Territory Decay



\## Goal



Belirli bir değerlendirme anında (EvaluatedAt), mevcut HexState için

effective dominance değerini deterministik olarak hesaplamak.



Bu görev yalnızca decay hesabını yapar.



Persistence, capture, ownership değişikliği, score, event ve

ResolveWalk kapsamında değildir.



\---



\## Read



Önce:



\- AGENTS.md

\- PROJECT\_STATUS.md

\- bu görev dosyası



Ardından yalnızca:



\- Calculator Rules v1.0

&#x20; - Territory Decay

&#x20; - Processing Order

&#x20; - Error Model

&#x20; - Edge Cases

\- Project Conventions v1.1

\- backend/internal/gameengine/input.go

\- backend/internal/gameengine/rules.go



Başka doküman okuma.



\---



\## Implement



Package-private fonksiyon ekle:



calculateEffectiveDominance(...)



Mevcut HexState ve TerritoryRules tipleri yeniden kullanılacaktır.



Yeni public API oluşturulmayacaktır.



\---



\## Rules



\- Hesaplama deterministik olacaktır.

\- Girdi yapıları değiştirilmeyecektir.

\- HexState güncellenmeyecektir.

\- Sonuç yalnızca hesaplanacaktır.

\- Negatif dominance üretilemez.

\- Minimum ve maksimum sınırlar Calculator Rules'a göre uygulanacaktır.

\- Değerlendirme zamanı çağıran katman tarafından verilecektir.

\- Sistem saati okunmayacaktır.

\- Time.Now() kullanılmayacaktır.



\---



\## Do Not



Eklenmeyecek:



\- ownership

\- territory resolution

\- capture

\- defense

\- attack

\- score

\- event

\- persistence

\- database

\- H3

\- validation

\- ResolveWalk

\- yeni public API



\---



\## Tests



Package içi testler:



\- decay uygulanmayan durum

\- normal decay

\- maksimum süre

\- minimum dominance sınırı

\- deterministik çıktı

\- farklı evaluation zamanı

\- girdi değişmezliği



\---



\## Files



Beklenen değişiklikler:



\- backend/internal/gameengine/territory\_decay.go

\- backend/internal/gameengine/territory\_decay\_test.go



Tüm kontroller başarılı olursa:



\- PROJECT\_STATUS.md



\---



\## Validation



```bash

gofmt -w internal/gameengine/territory\_decay.go internal/gameengine/territory\_decay\_test.go

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



\- Calculator Rules'taki decay formülü eksiksiz uygulanmıştır.

\- Sonuç deterministiktir.

\- Time.Now() kullanılmamıştır.

\- Girdi değişmemektedir.

\- Tüm testler geçmektedir.



\---



\## Completion Report



```text

Task ID: M1-011



Changed files:

\- backend/internal/gameengine/territory\_decay.go

\- backend/internal/gameengine/territory\_decay\_test.go

\- PROJECT\_STATUS.md



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- Territory ownership ve resolution sonraki görevde uygulanacaktır.

```

