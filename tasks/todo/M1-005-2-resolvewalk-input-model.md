\# M1-005 — ResolveWalk Input Model



\## Goal



Calculator Engine'in yürüyüş çözümleme işlemi için kullandığı immutable giriş veri modelini oluşturmak.



Bu görev yalnızca `ResolveWalkInput`, `LocationPoint` ve `HexState` modellerini tanımlar.



Validation, hesaplama, engine execution veya persistence davranışı içermez.



\---



\## Read



\- Calculator Rules v1.0

&#x20; - Section 7 — Engine Input Contract

\- MVP Domain Model v1.0

\- M1-004 — CityBoundary Model



\---



\## Implement



Aşağıdaki veri modellerini oluştur:



\- ResolveWalkInput

\- LocationPoint

\- HexState



\---



\## Required Models



\### ResolveWalkInput



Alanlar:



\- WalkID id.WalkID

\- PlayerID id.PlayerID

\- CityID id.CityID

\- StartedAt time.Time

\- EndedAt time.Time

\- EvaluatedAt time.Time

\- Route \[]LocationPoint

\- ExistingHexes map\[id.HexID]HexState

\- Boundary CityBoundary

\- EngineVersion id.EngineVersion



\### LocationPoint



Alanlar:



\- Sequence int

\- Latitude float64

\- Longitude float64

\- RecordedAt time.Time

\- AccuracyMeters float64

\- IsMockLocation bool



\### HexState



Alanlar:



\- HexID id.HexID

\- OwnerID \*id.PlayerID

\- Dominance int

\- LastUpdatedAt time.Time

\- Version int64



\## Architecture Rules



\- `Boundary`, M1-004 görevinde tanımlanan `CityBoundary` tipini kullanmalı.

\- Tüm domain kimlikleri `internal/domain/id` paketindeki güçlü tipleri kullanmalı.

\- Zaman alanları `time.Time` olmalı.

\- `Route`, sıralı konum noktalarını taşımalı.

\- `ExistingHexes`, `id.HexID` anahtarlı mevcut hex durumlarını taşımalı.

\- Modeller yalnızca veri taşımalı.

\- Slice ve map alanları için defensive copy veya accessor eklenmemeli.



\---



\## Files



Create:



\- backend/internal/gameengine/input.go

\- backend/internal/gameengine/input\_test.go



Expected additional change:



\- PROJECT\_STATUS.md



Başka dosya değiştirilmemeli.



\---



\## Tests



Odaklı testlerle aşağıdakileri doğrula:



\- `ResolveWalkInput` doğru güçlü domain kimlik tiplerini kullanıyor.

\- `Boundary`, `CityBoundary` tipinde.

\- `Route`, `\[]LocationPoint` tipinde.

\- `ExistingHexes`, `map\[id.HexID]HexState` tipinde.

\- Zaman alanları `time.Time` değerlerini koruyor.

\- `LocationPoint` alan değerleri korunuyor.

\- `HexState` alan değerleri korunuyor.

\- Route, boundary ve existing hex verileri model içinde doğru taşınıyor.



Testler validation veya business behavior içermemeli.



\---



\## Do Not



\- Validation ekleme.

\- Constructor veya factory ekleme.

\- Builder ekleme.

\- Default değer ekleme.

\- JSON, database veya ORM etiketi ekleme.

\- Defensive copy ekleme.

\- Getter veya setter ekleme.

\- Route sıralama davranışı ekleme.

\- Koordinat doğrulaması ekleme.

\- Timestamp karşılaştırması ekleme.

\- Boundary doğrulaması ekleme.

\- Hex ownership veya dominance mantığı ekleme.

\- ResolveWalk result modeli ekleme.

\- Engine execution ekleme.

\- H3 entegrasyonu ekleme.

\- Persistence veya API modeli ekleme.



\---



\## Validation



Run:



\- gofmt

\- go test ./...

\- make backend-check

\- git diff --check



\---



\## Pass Criteria



\- Üç veri modeli Calculator Rules sözleşmesine uygun şekilde tanımlanmış.

\- Tüm kimlik alanları güçlü domain tiplerini kullanıyor.

\- `CityBoundary` mevcut modelden yeniden kullanılıyor.

\- Modeller yalnızca immutable değer sözleşmesi olarak kalıyor.

\- Validation veya business logic eklenmiyor.

\- Tüm kontroller başarıyla geçiyor.

\- `PROJECT\_STATUS.md` yalnızca kontroller geçtikten sonra güncelleniyor.

