
# SMS ve E-posta Doğrulamalı Servis

Bu servis, SMS ve e-posta doğrulama işlevsellikleri sunarak diğer hizmetlere güvenli ve doğrulanmış erişim sağlar. Doğrulama süreci tamamlanmadan diğer hizmetlere erişimi engelleyen bir yetkilendirme mekanizması içerir. Servis, doğrulama ve yetkilendirme işlemleri için özel bir kimlik doğrulama servisi ile entegre edilmiştir. Veriler PostgreSQL kullanılarak yönetilmektedir.


## Özellikler

- SMS Doğrulama: Kullanıcılara doğrulama kodu içeren SMS gönderir.
- E-posta Doğrulama: Kullanıcılara e-posta doğrulama bağlantısı içeren  e-postalar gönderir.
- Yetkilendirme Mekanizması: Doğrulama tamamlanmadan diğer hizmetlere erişim engellenir.
- Entegrasyon: Kimlik doğrulama ve yetkilendirme işlemleri için harici bir kimlik doğrulama servisi ile sorunsuz entegrasyon.
- Veritabanı: PostgreSQL kullanılarak doğrulama verileri ve kullanıcı bilgileri saklanır.


## Installation



```bash
git clone https://github.com/kullanici-adi/verification-service.git
cd verification-service
```
```bash
Her go.mod dosyası olan dizin için
go mod tidy
```
```bash
Önemli verilerin okunacağı config.json yapılandır.
{
    "Database":{
            "Host":"...",
            "Port":...,
            "User":"...",
            "Password":"...",
            "Dbname":"..."
    },
    "Verification":{
            "accountSid" : "...",
            "authToken" : "...",
            "verificationString" : "..."
    },
    "Smtp":{
        "Address":"...",
        "Password":"..."
    }
}
```
```bash
go run Main/Main.go
```
    
