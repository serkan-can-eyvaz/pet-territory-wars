\# M1-017 — Domain Events



\## Goal



M1-012–M1-016 görevlerinde üretilen HexChange ve ScoreChange

sonuçlarından normatif DomainEvent kayıtlarını üretmek.



Bu görev yalnızca DomainEvent üretir.



ResolveWalk, persistence, publishing ve event dispatch kapsam dışındadır.



\---



\## Read



Önce:



\- AGENTS.md

\- PROJECT\_STATUS.md

\- bu görev dosyası



Ardından yalnızca:



\- Calculator Rules v1.0

&#x20; - Domain Events

&#x20; - Processing Order

&#x20; - Error Model

&#x20; - Edge Cases

\- Project Conventions v1.1

\- backend/internal/gameengine/result.go



Başka doküman okuma.



\---



\## Implement



Package-private fonksiyon:



resolveDomainEvents(

&#x20;   input ResolveWalkInput,

&#x20;   change HexChange,

&#x20;   scoreChanges \[]ScoreChange,

) \[]DomainEvent



Mevcut:



\- ResolveWalkInput

\- HexChange

\- ScoreChange

\- DomainEvent



tipleri yeniden kullanılacaktır.



Yeni public API veya yeni domain modeli oluşturulmayacaktır.



\---



\## Rules



Bu görev yalnızca mevcut sonuçları DomainEvent'e dönüştürecektir.



Ownership, dominance, score veya attack hesaplaması tekrar

yapılmayacaktır.



DomainEvent üretimi yalnızca Calculator Rules'taki normatif

event mapping'e göre yapılacaktır.



\---



\## NO\_CHANGE



change.ChangeType == NO\_CHANGE ise



nil olmayan boş \[]DomainEvent döndürülür.



\---



\## Event Mapping



Calculator Rules'taki event eşlemesi birebir uygulanacaktır.



Yeni event tipi oluşturulmayacaktır.



Repository'de mevcut event tipleri kullanılacaktır.



Eksik normatif event sabitleri varsa yalnızca mevcut EventType tipi

için package-level sabitler tanımlanacaktır.



Yeni enum veya yeni tip oluşturulmayacaktır.



\---



\## Event Payload



DomainEvent modelindeki mevcut alanlar kullanılacaktır.



Yeni payload modeli oluşturulmayacaktır.



Mevcut Payload tipi dışında alan eklenmeyecektir.



WalkID doğrudan ResolveWalkInput.WalkID olacaktır.



\---



\## Output



Çıktı sırası Calculator Rules'taki event üretim sırasını

aynen koruyacaktır.



Map iteration sonucu etkilemeyecektir.



\---



\## Do Not



Eklenmeyecek:



\- ResolveWalk

\- Event publishing

\- Kafka

\- Outbox

\- Persistence

\- Database

\- Validation

\- I/O

\- Sistem saati

\- Yeni public API

\- Yeni domain modeli



\---



\## Tests



Package-içi tablo testleri:



\- NO\_CHANGE

\- EMPTY\_CAPTURE event

\- DEFENSE event

\- ENEMY\_ATTACK event

\- OWNERSHIP\_TRANSFER event

\- score event mapping

\- doğru WalkID

\- determinizm

\- input değişmezliği



\---



\## Files



Beklenen değişiklikler:



\- backend/internal/gameengine/domain\_events.go

\- backend/internal/gameengine/domain\_events\_test.go



Gerekirse yalnızca mevcut EventType sabitlerini eklemek için:



\- backend/internal/gameengine/result.go



Tüm kontroller başarılı olursa:



\- PROJECT\_STATUS.md



\---



\## Validation



```bash

gofmt -w internal/gameengine/domain\_events.go internal/gameengine/domain\_events\_test.go

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



\- Calculator Rules'taki event mapping eksiksiz uygulanmıştır.

\- NO\_CHANGE boş slice döndürmektedir.

\- Yeni public API eklenmemiştir.

\- Çıktı sırası deterministiktir.

\- Tüm testler geçmektedir.



\---



\## Completion Report



```text

Task ID: M1-017



Changed files:

\- backend/internal/gameengine/domain\_events.go

\- backend/internal/gameengine/domain\_events\_test.go

\- backend/internal/gameengine/result.go (yalnızca normatif EventType sabitleri gerekiyorsa)

\- PROJECT\_STATUS.md



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- ResolveWalk (M1-018) sonraki görevde uygulanacaktır.

```

