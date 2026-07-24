\# M1-014 — Enemy Attack



\## Goal



Başka bir oyuncuya ait sahipli bir hex için ENEMY\_ATTACK kararını

hesaplamak.



Bu görev yalnızca saldırı sonucu oluşacak HexChange'i hesaplar.



Ownership transfer, score, event ve ResolveWalk kapsam dışındadır.



\---



\## Read



Önce:



\- AGENTS.md

\- PROJECT\_STATUS.md

\- bu görev dosyası



Ardından yalnızca:



\- Calculator Rules v1.0

&#x20; - Enemy Attack

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



resolveEnemyAttack(

&#x20;   input ResolveWalkInput,

&#x20;   state HexState,

&#x20;   visit VisitedHex,

&#x20;   movementRules MovementRules,

&#x20;   territoryRules TerritoryRules,

&#x20;   effectiveDominance int,

) HexChange



Yeni public API veya domain tipi oluşturulmayacaktır.



Mevcut ResolveWalkInput, HexState, VisitedHex, MovementRules,

TerritoryRules ve HexChange yeniden kullanılacaktır.



\---



\## Rules



Ön koşullar:



\- state.OwnerID != nil

\- state.OwnerID != input.PlayerID

\- ziyaret qualified olmalıdır



Qualification M1-012 ile aynıdır:



\- visit.PresenceSeconds >= movementRules.MinHexPresenceSeconds

&#x20; VEYA

\- visit.DistanceMeters >= movementRules.MinHexDistanceMeters



Karşılaştırmalar inclusive (>=) olacaktır.



Decay bu görevde tekrar hesaplanmayacaktır.



effectiveDominance çağıran katmandan gelecektir.



Başarılı saldırı formülü yalnızca Calculator Rules'taki normatif

Enemy Attack formülü kullanılarak hesaplanacaktır.



StoredDominance değiştirilmeyecektir.



Owner değiştirilmeyecektir.



\---



\## NO\_CHANGE



Aşağıdaki durumlarda NO\_CHANGE döndürülür:



\- owner yok

\- owner input.PlayerID ile aynı

\- ziyaret qualified değil

\- normatif saldırı sonucu değişiklik oluşturmuyor



NO\_CHANGE durumunda:



\- HexID = state.HexID

\- owner alanları korunur

\- StoredDominance korunur

\- EffectiveDominance = verilen effectiveDominance

\- NewDominance = effectiveDominance

\- ExpectedVersion korunur

\- ChangeType = NO\_CHANGE



\---



\## Do Not



Eklenmeyecek:



\- Empty Capture

\- Owner Defense

\- Ownership Transfer

\- Score

\- Domain Event

\- ResolveWalk

\- Decay hesabı

\- Validation

\- Persistence

\- Database

\- time.Now()

\- yeni public API



\---



\## Tests



Package-içi tablo testleri:



\- owner yok

\- aynı owner

\- qualified olmayan ziyaret

\- presence eşiği

\- distance eşiği

\- normatif attack formülü

\- saldırı sonucu değişiklik oluşmayan durum

\- determinizm

\- girdi değişmezliği



\---



\## Files



Beklenen değişiklikler:



\- backend/internal/gameengine/enemy\_attack.go

\- backend/internal/gameengine/enemy\_attack\_test.go



Tüm kontroller başarılı olursa:



\- PROJECT\_STATUS.md



\---



\## Validation



```bash

gofmt -w internal/gameengine/enemy\_attack.go internal/gameengine/enemy\_attack\_test.go

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



\- Enemy Attack yalnızca farklı owner için uygulanır.

\- Qualification M1-012 ile birebir aynıdır.

\- Decay tekrar hesaplanmaz.

\- Calculator Rules'taki normatif attack formülü uygulanır.

\- Deterministik sonuç üretilir.

\- Tüm testler geçer.



\---



\## Completion Report



```text

Task ID: M1-014



Changed files:

\- backend/internal/gameengine/enemy\_attack.go

\- backend/internal/gameengine/enemy\_attack\_test.go

\- PROJECT\_STATUS.md



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- Ownership Transfer (M1-015) sonraki görevde uygulanacaktır.

```

