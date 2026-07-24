\# M1-006 — HexChange Model



\## Goal



Calculator Engine tarafından üretilen hex değişikliklerinin normatif veri sözleşmesini oluşturmak.



Bu görev yalnızca aşağıdaki tipleri tanımlar:



\- `HexChange`

\- `HexChangeType`

\- `HexChangeReason`



Territory hesaplaması, validation, persistence, event üretimi veya `ResolveWalkResult` bu görevin kapsamında değildir.



\---



\## Read



Yalnızca:



\- Calculator Rules v1.0

&#x20; - Section 18 — HexChange Contract

&#x20; - Appendix A — Canonical Enumerations

\- Project Conventions v1.1



\---



\## Normative Source of Truth



Bu görevde Calculator Rules v1.0 normatif kaynaktır.



Section 18, `HexChange` alanlarını ve alanların anlamlarını tanımlar.



Appendix A yalnızca aşağıdaki canonical `HexChangeType` değerlerini tanımlar:



\- `EMPTY\_CAPTURE`

\- `OWNER\_DEFENSE`

\- `ENEMY\_ATTACK`

\- `OWNERSHIP\_TRANSFER`

\- `NO\_CHANGE`



Calculator Rules v1.0, `HexChangeReason` için canonical veya kapalı bir değer kümesi tanımlamaz.



Bu nedenle bu görevde normatif belgede bulunmayan reason sabitleri üretilmemelidir.



\---



\## Architecture Decision — HexChangeReason



`HexChangeReason`, kapalı bir enum değil, strongly typed bir açıklama kodu olarak tanımlanacaktır:



```go

type HexChangeReason string

```



Bu görevde:



\- `HexChangeReason` sabitleri tanımlanmayacak.

\- Reason kodları varsayılmayacak.

\- Reason validation eklenmeyecek.

\- Reason parsing veya formatting helper'ı eklenmeyecek.

\- Boş string davranışı belirlenmeyecek.



İleride canonical reason kodları normatif belgede tanımlanırsa ayrı bir görevle sabitler ve validation eklenebilir.



Bu karar, eksik normatif değerleri uydurmadan Section 18 sözleşmesindeki `Reason` alanının güçlü tipli olarak temsil edilmesini sağlar.



\---



\## Implement



\### HexChangeType



Aşağıdaki strongly typed string tipini oluştur:



```go

type HexChangeType string

```



Appendix A ile birebir uyumlu aşağıdaki canonical sabitleri tanımla:



```go

const (

&#x20;   HexChangeTypeEmptyCapture      HexChangeType = "EMPTY\_CAPTURE"

&#x20;   HexChangeTypeOwnerDefense      HexChangeType = "OWNER\_DEFENSE"

&#x20;   HexChangeTypeEnemyAttack       HexChangeType = "ENEMY\_ATTACK"

&#x20;   HexChangeTypeOwnershipTransfer HexChangeType = "OWNERSHIP\_TRANSFER"

&#x20;   HexChangeTypeNoChange          HexChangeType = "NO\_CHANGE"

)

```



Başka `HexChangeType` sabiti ekleme.



\### HexChangeReason



Aşağıdaki strongly typed string tipini oluştur:



```go

type HexChangeReason string

```



Bu görevde herhangi bir `HexChangeReason` sabiti ekleme.



\### HexChange



Aşağıdaki modeli oluştur:



```go

type HexChange struct {

&#x20;   HexID              id.HexID

&#x20;   ExpectedVersion    int64

&#x20;   PreviousOwnerID    \*id.PlayerID

&#x20;   NewOwnerID         \*id.PlayerID

&#x20;   StoredDominance    int

&#x20;   EffectiveDominance int

&#x20;   NewDominance       int

&#x20;   ChangeType         HexChangeType

&#x20;   Reason             HexChangeReason

}

```



Alan adlarını, sırasını veya tiplerini değiştirme.



\---



\## Field Semantics



\- `HexID`, değişikliğin uygulandığı hex kimliğidir.

\- `ExpectedVersion`, input içindeki authoritative `HexState.Version` değeridir.

\- `PreviousOwnerID`, resolution öncesindeki owner değeridir.

\- `NewOwnerID`, resolution sonrasındaki owner değeridir.

\- `nil` owner, sahipsiz hex'i temsil eder.

\- `StoredDominance`, input state içinde saklanan dominance değeridir.

\- `EffectiveDominance`, decay uygulandıktan sonra fakat yürüyüş etkisi uygulanmadan önceki dominance değeridir.

\- `NewDominance`, yürüyüş etkisi uygulandıktan sonraki dominance değeridir.

\- `ChangeType`, Appendix A'daki canonical değişiklik türlerinden biridir.

\- `Reason`, kullanıcıya gösterilmeyen deterministik açıklama kodudur.



Bu açıklamalar yalnızca veri sözleşmesinin anlamını belirtir. Bu görevde bu davranışlar hesaplanmaz veya doğrulanmaz.



\---



\## Files



Create:



\- `backend/internal/gameengine/hex\_change.go`

\- `backend/internal/gameengine/hex\_change\_test.go`



Update only after all checks pass:



\- `PROJECT\_STATUS.md`



Eski ve hatalı görev dosyası varsa kaldır:



\- `tasks/todo/M1-006-resolvewalk-result-model.md`



Task dosyası değişiklikleri dışında başka üretim dosyası değiştirilmemelidir.



\---



\## Implementation Requirements



\- Tipler `gameengine` paketinde tanımlanmalı.

\- Kimlikler `internal/domain/id` paketindeki güçlü tipleri kullanmalı.

\- Owner alanları pointer olmalı.

\- Dominance alanları `int` olmalı.

\- `ExpectedVersion` alanı `int64` olmalı.

\- Modeller yalnızca veri taşımalı.

\- JSON, database veya ORM etiketi kullanılmamalı.

\- Third-party dependency eklenmemeli.



\---



\## Tests



`hex\_change\_test.go` içinde yalnızca veri sözleşmesini doğrulayan odaklı testler yaz.



\### Required Test Coverage



1\. `HexChange` tüm alan değerlerini değişmeden taşıyor.

2\. `PreviousOwnerID == nil` ile önceki sahipsiz durum temsil edilebiliyor.

3\. `NewOwnerID == nil` ile sonraki sahipsiz durum temsil edilebiliyor.

4\. Dolu `\*id.PlayerID` owner değerleri korunuyor.

5\. `ExpectedVersion` değeri korunuyor.

6\. `StoredDominance`, `EffectiveDominance` ve `NewDominance` tamsayı değerlerini koruyor.

7\. `Reason`, strongly typed serbest bir deterministik kod değerini taşıyabiliyor.

8\. Aşağıdaki canonical `HexChangeType` değerleri tam olarak doğrulanıyor:

&#x20;  - `EMPTY\_CAPTURE`

&#x20;  - `OWNER\_DEFENSE`

&#x20;  - `ENEMY\_ATTACK`

&#x20;  - `OWNERSHIP\_TRANSFER`

&#x20;  - `NO\_CHANGE`



Testler reflection kullanarak struct alanlarını veya tag yokluğunu denetlememelidir.



\---



\## Do Not



\- `ResolveWalkResult` ekleme.

\- `ScoreChange` ekleme.

\- `VisitedHex` ekleme.

\- `ValidationResult` ekleme.

\- `WalkMetrics` ekleme.

\- `DomainEvent` ekleme.

\- `CalculationMetadata` ekleme.

\- Normatif belgede bulunmayan `HexChangeReason` sabitleri üretme.

\- `HexChangeType` için ek enum değeri üretme.

\- Validation ekleme.

\- Constructor, factory veya builder ekleme.

\- Getter veya setter ekleme.

\- String conversion helper ekleme.

\- Parse helper ekleme.

\- Default değer ekleme.

\- Territory resolution mantığı ekleme.

\- Decay hesaplaması ekleme.

\- Capture, defense, attack veya transfer hesaplaması ekleme.

\- Optimistic-lock kontrolü ekleme.

\- Database veya persistence davranışı ekleme.

\- JSON, database veya ORM etiketi ekleme.

\- API modeli ekleme.

\- Third-party dependency ekleme.



\---



\## Validation



Repository kökünden veya mevcut proje komut düzenine uygun dizinden çalıştır:



```bash

gofmt -w backend/internal/gameengine/hex\_change.go backend/internal/gameengine/hex\_change\_test.go

go test ./...

make backend-check

git diff --check

```



Komutların çalışma dizinini repository yapısına göre ayarla; doğrulama kapsamını değiştirme.



\---



\## Pass Criteria



Görev yalnızca aşağıdaki koşulların tamamı sağlandığında tamamlanır:



\- `HexChange`, Calculator Rules v1.0 Section 18 sözleşmesiyle uyumludur.

\- `HexChangeType` yalnızca Appendix A'daki beş canonical değeri içerir.

\- `HexChangeReason`, strongly typed string olarak tanımlanmıştır.

\- Normatif belgede bulunmayan reason sabitleri eklenmemiştir.

\- Owner alanları `\*id.PlayerID` tipindedir.

\- Sahipsiz owner `nil` ile temsil edilebilir.

\- Dominance alanları `int` tipindedir.

\- `ExpectedVersion`, `int64` tipindedir.

\- Modelde validation veya business logic bulunmaz.

\- Beklenmeyen üretim dosyaları değiştirilmemiştir.

\- `gofmt` başarılıdır.

\- `go test ./...` başarılıdır.

\- `make backend-check` başarılıdır.

\- `git diff --check` başarılıdır.

\- `PROJECT\_STATUS.md` yalnızca bütün kontroller geçtikten sonra güncellenmiştir.



\---



\## Completion Report



Görev sonunda aşağıdaki formatta rapor ver:



```text

Task ID: M1-006



Changed files:

\- ...



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- HexChangeReason canonical değerleri Calculator Rules v1.0 tarafından henüz tanımlanmıyor.

```

