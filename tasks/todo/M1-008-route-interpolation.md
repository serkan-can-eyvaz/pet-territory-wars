\# M1-008 — Route Interpolation



\## Goal



M1-006 tarafından üretilen geçerli RouteSegment'lerden belirlenen aralıklarla ara konumlar (interpolated points) üretmek.



Bu görev yalnızca rota interpolasyonunu gerçekleştirir. Hex üretimi, H3 dönüşümü, territory, skor ve event davranışları bu görevin kapsamında değildir.



\---



\## Read



Sadece aşağıdaki doküman ve dosyaları oku:



\- Calculator Rules v1.0

&#x20; - Route Interpolation

\- Project Conventions v1.1

\- backend/internal/gameengine/route\_segment.go

\- backend/internal/gameengine/input.go



Başka doküman okuma.



\---



\## Implement



Yeni package-private fonksiyonlar ekle:



\- interpolateRoute(...)

\- interpolateSegment(...)



Kurallar Calculator Rules v1.0 ile birebir uyumlu olacaktır.



\---



\## Interpolation Output Contract



Interpolation output uses this package-private model:



```go

type interpolatedPoint struct {
    Latitude  float64
    Longitude float64
}

```



\- Interpolated points must not invent a `LocationPoint.Sequence` value.
\- Source `LocationPoint` values and their `Sequence` fields must remain unchanged.
\- The combined output includes the first valid segment's starting source point.
\- When a valid source segment directly continues the preceding valid source segment, its shared source endpoint appears only once in the combined output.
\- A non-valid segment separates valid route parts. Do not deduplicate across that boundary and do not create a synthetic connection.



\---



\## Rules



\- Yalnızca geçerli (`IsValid == true`) segmentler işlenecek.

\- Geçersiz segmentlerden interpolasyon noktası üretilmeyecek.

\- Segment başlangıç ve bitiş sırası korunacak.

\- Noktalar deterministik üretilecek.

\- Aynı giriş her zaman aynı interpolasyon çıktısını üretmelidir.

\- Segment uzunluğu interpolasyon aralığından kısa ise yalnızca mevcut uç noktalar kullanılacaktır.

\- Girdi segmentleri değiştirilmeyecek.



\---



\## Do Not



Eklenmeyecek:



\- H3 dönüşümü

\- Hex üretimi

\- Territory davranışı

\- Score hesaplama

\- Event üretimi

\- Walk validation

\- ResolveWalk

\- Public API

\- Third-party dependency



\---



\## Tests



Odaklı testler:



\- kısa segment

\- uzun segment

\- geçersiz segment

\- interpolasyon sırası

\- deterministik sonuç

\- giriş verisinin değişmemesi

\- tekrar eden hesaplamalarda aynı çıktı



\---



\## Files



Beklenen değişiklikler:



\- backend/internal/gameengine/route\_interpolation.go

\- backend/internal/gameengine/route\_interpolation\_test.go



Tüm kontroller geçerse:



\- PROJECT\_STATUS.md



\---



\## Validation



```bash

gofmt -w internal/gameengine/route\_interpolation.go internal/gameengine/route\_interpolation\_test.go

go test ./...

```



Repository root:



```bash

make backend-check

git diff --check

```



PROJECT\_STATUS.md yalnızca tüm kontroller başarılı olduktan sonra güncellenecektir.



\---



\## Completion Report



```text

Task ID: M1-008



Changed files:

\- backend/internal/gameengine/route\_interpolation.go

\- backend/internal/gameengine/route\_interpolation\_test.go

\- PROJECT\_STATUS.md



Checks run:

\- Pass: gofmt

\- Pass: go test ./...

\- Pass: make backend-check

\- Pass: git diff --check



Remaining risk:

\- H3 dönüşümü, hex üretimi, territory, score ve event davranışları sonraki görevlere bırakılmıştır.

```

