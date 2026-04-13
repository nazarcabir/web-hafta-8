# 🎓 GoLearn Education Platform - Final Project

<div align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Gin-008EC4?style=for-the-badge&logo=gin&logoColor=white" alt="Gin" />
  <img src="https://img.shields.io/badge/GORM-3F51B5?style=for-the-badge&logo=gorm&logoColor=white" alt="GORM" />
  <img src="https://img.shields.io/badge/SQLite-003B57?style=for-the-badge&logo=sqlite&logoColor=white" alt="SQLite" />
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
  <img src="https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=white" alt="Swagger" />
</div>

---

## 👤 Öğrenci Bilgileri
- **Ad Soyad:** Nazar Cabir Cabır
- **Numara:** 24080410036

---

## 🚀 Proje Hakkında
GoLearn, Go dili ve Gin framework'ü kullanılarak geliştirilmiş, modern ve ölçeklenebilir bir Uzaktan Eğitim (LMS) Backend servisidir. Bu proje, kapsamlı bir eğitim platformunun tüm temel gereksinimlerini (Yetkilendirme, İçerik Yönetimi, Sınav Sistemleri, İlerleme Takibi ve Real-time İletişim) karşılayacak şekilde tasarlanmıştır.

## 🛠️ Uygulanan 13 Adım (Geliştirme Süreci)

1.  **Backend Altyapısı ve Gin Kurulumu:** Yüksek performanslı Gin framework'ü ile proje iskeleti oluşturuldu.
2.  **Veritabanı Katmanı (GORM & SQLite):** SQLite tabanlı, GORM ORM kütüphanesi ile veritabanı bağlantısı ve otomatik migrasyonlar (AutoMigrate) kuruldu.
3.  **JWT Tabanlı Kimlik Doğrulama:** Güvenli oturum yönetimi için JSON Web Token (JWT) entegrasyonu sağlandı.
4.  **Kullanıcı Kayıt ve Giriş Sistemi:** Şifreleri Bcrypt ile hash'leyen güvenli Register ve Login endpoint'leri geliştirildi.
5.  **Rol Tabanlı Yetkilendirme (RBAC):** `Teacher` ve `Student` rolleri için middleware seviyesinde yetki kontrolleri yapıldı.
6.  **Kurs Yönetim Servisi:** Kurs oluşturma, güncelleme, silme ve gelişmiş filtreleme (Kategori, Sıralama, Sayfalama) özellikleri eklendi.
7.  **Ders İçerik Yönetimi:** Kurslara bağlı derslerin (Lessons) eklenmesi ve sıralanması sağlandı.
8.  **Quiz Sistemi ve Soru Bankası:** Derslere bağlı sınavların ve soruların oluşturulabilmesi için gerekli modeller ve handler'lar yazıldı.
9.  **Sınav Çözme ve Puanlama Motoru:** Öğrencilerin cevaplarını kontrol eden ve başarı oranını anlık hesaplayan algoritma kuruldu.
10. **Quiz Sonuçlarının Kalıcılığı:** Çözülen sınavların sonuçlarının (Skor, Doğru Sayısı vb.) veritabanına kaydedilmesi sağlandı.
11. **Kurs Kayıt (Enrollment) Sistemi:** Öğrencilerin kurslara kaydolabilmesi ve kurs-öğrenci ilişkisinin takibi sağlandı.
12. **Gelişmiş İlerleme Takibi (Progress Tracking):** Tamamlanan derslerin ders sayısına oranlanarak yüzde bazında ilerleme gösterilmesi sağlandı.
13. **Real-time Sınıf Sohbeti (WebSocket) & Swagger:** WebSocket ile canlı sohbet altyapısı kuruldu ve tüm API Swagger ile dokümante edilerek Dockerize edildi.

---

## 📦 Kurulum ve Çalıştırma

### Docker ile (Önerilen)
Proje tamamen Dockerize edilmiştir. Tek bir komutla tüm bağımlılıkları çalıştırabilirsiniz:

```bash
docker-compose up --build
```

API'ye **http://localhost:8090** adresinden ulaşabilirsiniz.

### Manuel Çalıştırma (Windows PowerShell)
```powershell
./setup.ps1
go run main.go
```

---

## 📖 API Dokümantasyonu
Swagger UI üzerinden tüm endpoint'leri test edebilirsiniz:
👉 **[http://localhost:8090/swagger/index.html](http://localhost:8090/swagger/index.html)**

---

<div align="center">
  <sub>Web Geliştirme Dersi - Hafta 8 Final Projesi</sub>
</div>
