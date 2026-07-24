\# M1-013 — Owner Defense



\## Goal



Owner tarafından ziyaret edilen sahipli bir hex için DEFENSE kararını

hesaplamak.



Bu görev yalnızca Owner Defense davranışını uygular.



\---



\## Read



Önce:



\- AGENTS.md

\- PROJECT\_STATUS.md

\- bu görev dosyası



Ardından yalnızca:



\- Calculator Rules v1.0

&#x20; - Owner Defense

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



resolveOwnerDefense(...) HexChange



Yeni public API veya yeni domain modeli oluşturulmayacaktır.



Mevcut:



\- ResolveWalkInput

\- HexState

\- VisitedHex

\- MovementRules

\- TerritoryRules

\- HexChange



tipleri yeniden kullanılacaktır.



\---



\## Rules



\- Yalnızca OwnerID != nil olan hex değerlendirilecektir.

\- ResolveWalkInput.PlayerID == HexState.OwnerID olmalıdır.

\- Qualification M1-012 ile aynıdır.



Qualified ziyaret:



\- PresenceSeconds >= MinHexPresenceSeconds

&#x20; VEYA

\- DistanceMeters >= MinHexDistanceMeters



(>= inclusive)



Qualified değilse NO\_CHANGE döndürülür.



Owner farklıysa NO\_CHANGE döndürülür.



Başarılı defense durumunda:



\- Owner değişmez.

\- StoredDominance mevcut dominance + DefenseIncrease kadar artırılır.

\- Sonuç MaximumDominance değerini geçemez.

\- EffectiveDominance aynı yeni dominance değeri olur.

\- ExpectedVersion korunur.

\- ChangeType DEFENSE olur.



\---



\## Do Not



Eklenmeyecek:



\- Empty Capture

\- Enemy Attack

\- Ownership Transfer

\- Score

\- Domain Event

\- ResolveWalk

\- Validation

\- Persistence

\- Database

\- time.Now()

\- yeni public API



\---



\## NO\_CHANGE



Capture görevindeki kuralla aynı davranış uygulanacaktır.



HexChange mevcut state'i aynen temsil edecektir.



\---



\## Tests



Package içi testler:



\- owner olmayan hex

\- farklı owner

\- qualified olmayan ziyaret

\- presence eşiği

\- distance eşiği

\- başarılı defense

\- MaximumDominance clamp

\- ExpectedVersion korunması

\- determinizm

\- girdi değişmezliği



\---



\## Files



Beklenen değişiklikler:



\- backend/internal/gameengine/owner\_defense.go

\- backend/internal/gameengine/owner\_defense\_test.go



Tüm kontroller başarılı olursa:



\- PROJECT\_STATUS.md



\---



\## Validation



```bash

gofmt -w internal/gameengine/owner\_defense.go internal/gameengine/owner\_defense\_test.go

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



\- DEFENSE yalnızca owner tarafından uygulanır.

\- Qualification M1-012 ile birebir aynıdır.

\- Dominance MaximumDominance değerini geçmez.

\- Owner değişmez.

\- Deterministik sonuç üretilir.

\- Tüm testler geçer.



\---



\## Completion Report



```text

Task ID: M1-013



Changed files:

\- backend/internal/gameengine/owner\_defense.go

\- backend/internal/gameengine/owner\_defense\_test.go

\- PROJECT\_STATUS.md



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- Enemy Attack (M1-014) sonraki görevde uygulanacaktır.

```

