\# M1-015 — Ownership Transfer



\## Goal



Enemy Attack sonucu effective dominance değeri sıfıra düşen bir hex için

Ownership Transfer kararını hesaplamak.



Bu görev yalnızca sahiplik transferini uygular.



Score, Domain Event ve ResolveWalk kapsam dışındadır.



\---



\## Read



Önce:



\- AGENTS.md

\- PROJECT\_STATUS.md

\- bu görev dosyası



Ardından yalnızca:



\- Calculator Rules v1.0

&#x20; - Ownership Transfer

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



Package-private fonksiyon:



resolveOwnershipTransfer(

&#x20;   input ResolveWalkInput,

&#x20;   state HexState,

&#x20;   attackChange HexChange,

&#x20;   territoryRules TerritoryRules,

) HexChange



Yeni public API veya yeni domain tipi oluşturulmayacaktır.



Mevcut:



\- ResolveWalkInput

\- HexState

\- TerritoryRules

\- HexChange



tipleri yeniden kullanılacaktır.



\---



\## Rules



Ownership Transfer yalnızca aşağıdaki koşullar birlikte sağlandığında uygulanacaktır:



\- attackChange.ChangeType == ENEMY\_ATTACK

\- attackChange.NewDominance == 0

\- state.OwnerID != nil



Bu görev Enemy Attack hesabını tekrar yapmayacaktır.



Başarılı transfer durumunda:



\- NewOwnerID = input.PlayerID

\- PreviousOwnerID = state.OwnerID

\- StoredDominance = territoryRules.InitialDominance

\- EffectiveDominance = territoryRules.InitialDominance

\- NewDominance = territoryRules.InitialDominance

\- ExpectedVersion = state.ExpectedVersion

\- ChangeType = OWNERSHIP\_TRANSFER



HexID her durumda state.HexID olacaktır.



\---



\## NO\_CHANGE



Aşağıdaki durumlarda NO\_CHANGE döndürülür:



\- attackChange.ChangeType != ENEMY\_ATTACK

\- attackChange.NewDominance > 0

\- state.OwnerID == nil



NO\_CHANGE sonucunda:



\- HexID korunur

\- Owner alanları korunur

\- StoredDominance korunur

\- EffectiveDominance = attackChange.EffectiveDominance

\- NewDominance = attackChange.NewDominance

\- ExpectedVersion korunur

\- ChangeType = NO\_CHANGE



\---



\## Do Not



Eklenmeyecek:



\- Enemy Attack hesabı

\- Empty Capture

\- Owner Defense

\- Score

\- Domain Event

\- ResolveWalk

\- Validation

\- Persistence

\- Database

\- time.Now()

\- I/O

\- State mutation

\- Yeni public API



\---



\## Tests



Package-içi tablo testleri:



\- attack olmayan change

\- dominance sıfır değil

\- owner olmayan state

\- başarılı ownership transfer

\- InitialDominance atanması

\- owner değişimi

\- ExpectedVersion korunması

\- HexID korunması

\- determinizm

\- input değişmezliği



\---



\## Files



Beklenen değişiklikler:



\- backend/internal/gameengine/ownership\_transfer.go

\- backend/internal/gameengine/ownership\_transfer\_test.go



Tüm kontroller başarılı olursa:



\- PROJECT\_STATUS.md



\---



\## Validation



```bash

gofmt -w internal/gameengine/ownership\_transfer.go internal/gameengine/ownership\_transfer\_test.go

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



\- Ownership yalnızca ENEMY\_ATTACK sonucu dominance sıfır olduğunda değişir.

\- Yeni owner input.PlayerID olur.

\- InitialDominance doğru uygulanır.

\- Enemy Attack bu görevde tekrar hesaplanmaz.

\- Deterministik sonuç üretilir.

\- Tüm testler geçer.



\---



\## Completion Report



```text

Task ID: M1-015



Changed files:

\- backend/internal/gameengine/ownership\_transfer.go

\- backend/internal/gameengine/ownership\_transfer\_test.go

\- PROJECT\_STATUS.md



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- Score Changes (M1-016) sonraki görevde uygulanacaktır.

```

