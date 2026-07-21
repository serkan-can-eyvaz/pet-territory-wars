\# M0-012 — Makefile



Status: TODO



\## Goal



Repository içindeki mevcut backend, Flutter, Docker Compose ve migration

komutlarını kök dizindeki tek bir Makefile üzerinden çalıştırılabilir hâle

getirmek.



Makefile yalnızca mevcut işlemleri sarmalayacak; yeni araç, bağımlılık veya

uygulama davranışı eklemeyecek.



\---



\## Read



\- AGENTS.md

\- PROJECT\_STATUS.md

\- docs/architecture/Pet Territory Wars — Project Conventions v1.1.pdf

\- docs/architecture/Pet Territory Wars — Backend Architecture v1.1.pdf

\- docs/architecture/Pet Territory Wars — Infrastructure \& Deployment v1.0.pdf

\- docs/roadmap/Pet Territory Wars — Milestone 0 Repository \& Development Environment.pdf



Do not read unrelated documents.



Belgelerde zorunlu hedef isimleri varsa onları kullan.

Belgeler çelişiyorsa dur ve raporla; tahmin etme.



\---



\## Do



\- Kök dizinde tek bir Makefile oluştur.

\- Varsayılan hedef yardım çıktısı versin.

\- Hedefleri `.PHONY` olarak tanımla.

\- Her hedef için kısa yardım açıklaması ekle.

\- Backend komutlarını `backend/` dizininde çalıştır.

\- Flutter komutlarını mevcut Flutter proje dizininde çalıştır.

\- Docker Compose komutlarını kök dizinde çalıştır.

\- Migration komutlarını `backend/migrations/` ile mevcut `DATABASE\_URL`

&#x20; üzerinden çalıştır.

\- Başarısız alt komutların Make hedefini başarısız etmesini sağla.

\- Environment değerlerini ve credential bilgilerini çıktıya yazdırma.

\- Tüm kontroller geçmeden PROJECT\_STATUS.md güncelleme.



\---



\## Required Targets



```text

help



backend-format

backend-vet

backend-test

backend-build

backend-check



flutter-format

flutter-analyze

flutter-test

flutter-check



compose-config

compose-up

compose-down

compose-logs



migrate-up

migrate-down

migrate-version

