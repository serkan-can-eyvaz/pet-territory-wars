\# M1-016 — Score Changes



\## Goal



M1-012, M1-013, M1-014 ve M1-015 tarafından üretilen HexChange sonucundan

oyuncuya yazılacak ScoreChange kayıtlarını deterministik olarak üretmek.



Bu görev yalnızca ScoreChange üretir.



Score hesaplaması tekrar yapılmayacak.

Ownership, dominance, event veya ResolveWalk davranışı uygulanmayacaktır.



\---



\## Read



Önce:



\- AGENTS.md

\- PROJECT\_STATUS.md

\- bu görev dosyası



Ardından yalnızca:



\- Calculator Rules v1.0

&#x20; - Score Changes

&#x20; - Processing Order

&#x20; - Error Model

&#x20; - Edge Cases

\- Project Conventions v1.1

\- backend/internal/gameengine/result.go

\- backend/internal/gameengine/rules.go



Başka doküman okuma.



\---



\## Implement



Package-private fonksiyon:



resolveScoreChanges(

&#x20;   input ResolveWalkInput,

&#x20;   change HexChange,

&#x20;   rules ScoreRules,

) \[]ScoreChange



Mevcut:



\- ResolveWalkInput

\- HexChange

\- ScoreRules

\- ScoreChange



tipleri yeniden kullanılacaktır.



Yeni public API veya yeni domain modeli oluşturulmayacaktır.



\---



\## Rules



Bu görev yalnızca mevcut HexChange değerini yorumlayacaktır.



Ownership Transfer, Enemy Attack, Owner Defense veya Empty Capture

tekrar hesaplanmayacaktır.



Score yalnızca Calculator Rules'taki ChangeType → Score eşlemesine göre

üretilecektir.



ChangeType için karşılık gelen puan değeri mevcut ScoreRules alanlarından

okunacaktır.



Yeni sabit, enum veya puan tablosu oluşturulmayacaktır.



\---



\## NO\_CHANGE



change.ChangeType == NO\_CHANGE ise



boş \[]ScoreChange döndürülür.



Nil yerine boş slice kullanılacaktır.



\---



\## Successful Result



Skor oluştuğunda mevcut ScoreChange modeli kullanılacaktır.



PlayerID doğrudan ResolveWalkInput.PlayerID olacaktır.



HexID doğrudan HexChange.HexID olacaktır.



Reason alanı mevcut ChangeType ile eşleşen normatif reason olacaktır.



Score değeri Calculator Rules ve mevcut ScoreRules alanlarından alınacaktır.



Yeni reason veya helper enum oluşturulmayacaktır.



\---



\## Do Not



Eklenmeyecek:



\- Ownership hesaplama

\- Enemy Attack

\- Owner Defense

\- Empty Capture

\- ResolveWalk

\- Domain Event

\- Validation

\- Persistence

\- Database

\- I/O

\- Sistem saati

\- Yeni public API

\- Yeni domain modeli



\---



\## Tests



Package-içi tablo testleri:



\- NO\_CHANGE → boş slice

\- EMPTY\_CAPTURE

\- DEFENSE

\- ENEMY\_ATTACK

\- OWNERSHIP\_TRANSFER

\- doğru PlayerID

\- doğru HexID

\- doğru Reason

\- doğru Score

\- determinizm

\- input değişmezliği



\---



\## Files



Beklenen değişiklikler:



\- backend/internal/gameengine/score\_changes.go

\- backend/internal/gameengine/score\_changes\_test.go



Tüm kontroller başarılı olursa:



\- PROJECT\_STATUS.md



\---



\## Validation



```bash

gofmt -w internal/gameengine/score\_changes.go internal/gameengine/score\_changes\_test.go

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



\- Calculator Rules'taki Score Mapping eksiksiz uygulanmıştır.

\- NO\_CHANGE boş slice döndürmektedir.

\- Yeni public API eklenmemiştir.

\- Deterministik sonuç üretilmektedir.

\- Tüm testler geçmektedir.



\---



\## Completion Report



```text

Task ID: M1-016



Changed files:

\- backend/internal/gameengine/score\_changes.go

\- backend/internal/gameengine/score\_changes\_test.go

\- PROJECT\_STATUS.md



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- Domain Events (M1-017) sonraki görevde uygulanacaktır.

```

