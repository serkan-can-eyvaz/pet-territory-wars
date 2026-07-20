\# M0-003 — Flutter Bootstrap



Status: TODO



\## Goal



Flutter mobile projesinin temel, derlenebilir ve test edilebilir başlangıç yapısını oluşturmak.



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Pet Territory Wars — Project Conventions v1.1.pdf

\- docs/roadmap/Pet Territory Wars — Milestone 0 Repository \& Development Environment.pdf



Do not read unrelated documents.



\## Do



\- `mobile` klasöründe Flutter projesini oluştur.

\- Mimari dokümanlarda belirtilen temel klasör yapısını oluştur.

\- Varsayılan örnek uygulama kodlarını kaldır.

\- Derlenebilir boş uygulama başlangıcı bırak.

\- Gerekiyorsa `.gitignore` ve Flutter yapılandırma dosyalarını oluştur.

\- Flutter analyze ve test komutlarının başarılı çalışmasını sağla.

\- Tüm Done koşulları geçmeden `PROJECT\_STATUS.md` dosyasını değiştirme.



\## Don't



\- UI tasarlama.

\- Sayfa ekleme.

\- Navigation oluşturma.

\- State management ekleme.

\- Riverpod, Bloc, Provider veya benzeri paket ekleme.

\- Firebase ekleme.

\- REST API istemcisi ekleme.

\- Authentication yazma.

\- Oyun ekranları oluşturma.

\- Asset ekleme.

\- Business logic yazma.

\- Sonraki tasklara geçme.



\## Done



\- \[ ] Flutter projesi oluşturuldu.

\- \[ ] Varsayılan demo uygulaması kaldırıldı.

\- \[ ] Temel klasör yapısı oluşturuldu.

\- \[ ] Kod formatlandı.

\- \[ ] `flutter analyze` başarılı.

\- \[ ] `flutter test` başarılı.

\- \[ ] Kapsam dışına çıkılmadı.

\- \[ ] PROJECT\_STATUS.md güncellendi.



\## Commands



Run inside the `mobile` directory:



```bash

flutter pub get

dart format .

flutter analyze

flutter test

```



\## Completion Update



Task tamamlanınca:



\- `PROJECT\_STATUS.md` içinde M0-003 Completed olarak işaretlenecek.

\- Sonraki aktif görev M0-004 olacak.

\- Otomatik olarak başka task başlatılmayacak.



\## Output



```text

Task ID: M0-003



Changed files:

\- ...



Checks run:

\- ...



Remaining risk:

\- None

```

