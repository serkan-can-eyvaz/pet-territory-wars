\# M1-004 — CityBoundary Model



\## Goal



Calculator Engine tarafından şehir sınırı kontrollerinde kullanılacak minimal CityBoundary veri modelini oluşturmak.



MVP kapsamında şehir sınırı, sıralı koordinatlardan oluşan tek bir dış polygon halkası olarak temsil edilir.



Bu görev yalnızca veri modelini tanımlar.

Koordinat doğrulaması, polygon hesaplaması veya point-in-polygon mantığı içermez.



\---



\## Read



\- Calculator Rules v1.0

&#x20; - Section 7 — Engine Input Contract

&#x20; - Section 9 — GPS Point and Segment Validation

&#x20; - Section 25 — Edge Cases and Invariants



\---



\## Architecture Decision



MVP CityBoundary modeli:



\- Tek bir şehir sınırını temsil eder.

\- Şehrin domain kimliğini taşır.

\- Polygon dış halkasını sıralı koordinatlar olarak taşır.

\- İlk ve son koordinatın aynı olması zorunluluğu bu görevde doğrulanmaz.

\- Polygon hole desteklemez.

\- MultiPolygon desteklemez.

\- Geometry library veya third-party geospatial type kullanmaz.



Hole ve MultiPolygon desteği MVP kapsamı dışındadır.



\---



\## Implement



Aşağıdaki veri modellerini oluştur:



\- GeoCoordinate

\- CityBoundary



\### GeoCoordinate



Alanlar:



\- Latitude float64

\- Longitude float64



\### CityBoundary



Alanlar:



\- CityID id.CityID

\- OuterRing \[]GeoCoordinate



\---



\## File



Create:



\- backend/internal/gameengine/city\_boundary.go

\- backend/internal/gameengine/city\_boundary\_test.go



\---



\## Model Requirements



\- Struct alanları exported olmalı.

\- `CityID`, önceki görevlerde tanımlanan `id.CityID` tipini kullanmalı.

\- Third-party geometry tipi kullanılmamalı.

\- JSON, database veya ORM etiketi eklenmemeli.

\- Model business logic içermemeli.

\- Constructor veya builder eklenmemeli.

\- Varsayılan boundary oluşturulmamalı.



\---



\## Tests



Odaklı testlerle aşağıdakileri doğrula:



\- `CityBoundary.CityID`, `id.CityID` tipini kullanıyor.

\- `OuterRing`, `\[]GeoCoordinate` tipinde.

\- `GeoCoordinate`, Latitude ve Longitude alanlarını taşıyor.

\- Alanlar değer atamasıyla doğru şekilde korunuyor.

\- Modelde JSON veya ORM davranışı bulunmuyor.



Testlerde polygon doğrulaması veya geospatial hesaplama yapma.



\---



\## Do Not



\- Point-in-polygon hesaplaması ekleme.

\- Boundary validation ekleme.

\- Latitude veya longitude validation ekleme.

\- Polygon kapatma davranışı ekleme.

\- Hole desteği ekleme.

\- MultiPolygon desteği ekleme.

\- H3 entegrasyonu ekleme.

\- Geometry dependency ekleme.

\- Constructor veya factory ekleme.

\- JSON tag ekleme.

\- Database veya API tipi ekleme.

\- ResolveWalkInput modelini bu görevde ekleme.



\---



\## Expected Changed Files



\- backend/internal/gameengine/city\_boundary.go

\- backend/internal/gameengine/city\_boundary\_test.go

\- PROJECT\_STATUS.md



Başka dosya değiştirilmemeli.



\---



\## Validation



Run:



\- gofmt

\- go test ./...

\- make backend-check

\- git diff --check



\---



\## Pass Criteria



\- CityBoundary modeli açık ve minimal şekilde tanımlanmış.

\- CityID doğru güçlü domain tipini kullanıyor.

\- Boundary dış halkası sıralı koordinatlarla temsil ediliyor.

\- Third-party geospatial type kullanılmıyor.

\- Validation veya geospatial business logic eklenmiyor.

\- Tüm kontroller başarıyla geçiyor.

\- PROJECT\_STATUS.md yalnızca kontroller geçtikten sonra güncelleniyor.

