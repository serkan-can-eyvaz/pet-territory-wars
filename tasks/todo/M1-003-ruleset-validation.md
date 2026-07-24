\# M1-003 — RuleSet Validation



\## Goal



Calculator Rules v1.0 normatif sözleşmesine uygun olarak RuleSet doğrulamasını uygulamak.



Bu görev yalnızca RuleSet değerlerinin geçerliliğini kontrol eder.

Business hesaplaması, varsayılan değer üretimi veya RuleSet oluşturma sorumluluğu içermez.



\---



\## Read



\- Calculator Rules v1.0

&#x20; - Section 3 — RuleSet Model

&#x20; - Section 4 — RuleSet Field Catalog

&#x20; - Section 6 — RuleSet Validation



\---



\## Implement



\- RuleSet doğrulama fonksiyonunu oluştur.

\- WalkRules alanlarını doğrula.

\- MovementRules alanlarını doğrula.

\- TerritoryRules alanlarını doğrula.

\- Cross-field kurallarını doğrula.

\- Float alanlarda NaN ve ±Inf kontrolü yap.

\- Başarılı doğrulamada `nil` döndür.

\- Başarısız doğrulamada açıklayıcı error döndür.



\---



\## Validation Rules



\### WalkRules



\- Version boş olamaz.

\- MinDurationSeconds >= 0

\- MinDistanceMeters >= 0

\- MinValidRouteRatio ∈ \[0,1]

\- MaxSpeedMPS > 0

\- MaxAccuracyMeters >= 0

\- MaxJumpMeters > 0



\### MovementRules



\- H3Resolution > 0

\- MinHexPresenceSeconds >= 0

\- MinHexDistanceMeters >= 0

\- InterpolationMeters > 0



\### TerritoryRules



\- MaxDominance > 0

\- InitialDominance ∈ \[1, MaxDominance]

\- OwnerVisitGain >= 0

\- EnemyAttackDamage > 0

\- DailyDecay >= 0

\- MinimumOwnedDominance ∈ \[1, MaxDominance]

\- ThreatThreshold ∈ \[0, MaxDominance]



\### Float Validation



Tüm float64 alanları:



\- NaN olamaz.

\- +Inf olamaz.

\- -Inf olamaz.



\---



\## Do Not



\- Default değer üretme.

\- RuleSet oluşturma.

\- Builder ekleme.

\- Constructor ekleme.

\- Engine kodu yazma.

\- Territory hesabı yapma.

\- JSON ekleme.

\- Persistence ekleme.



\---



\## Files



Expected:



\- backend/internal/gameengine/rules\_validation.go

\- backend/internal/gameengine/rules\_validation\_test.go



\---



\## Validation



Run:



\- gofmt

\- go test ./...

\- make backend-check

\- git diff --check



\---



\## Pass Criteria



\- Calculator Rules v1.0 Section 6 tamamen karşılanıyor.

\- Tüm validation kuralları test ediliyor.

\- Geçerli RuleSet başarıyla doğrulanıyor.

\- Geçersiz RuleSet'ler deterministik şekilde hata döndürüyor.

\- Başka hiçbir business logic eklenmiyor.

