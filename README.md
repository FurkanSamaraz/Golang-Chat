# Golang-Chat
Redis - WebSocket - Go

# Golang Chat Uygulaması

Bu, Golang tabanlı bir sohbet uygulamasıdır. Uygulama, Redis ve PostgreSQL veritabanlarını kullanmaktadır.

## Başlangıç

Bu talimatlar, projeyi yerel ortamınızda çalıştırmak için gereken adımları içermektedir. Aşağıdaki önkoşulların sağlandığından emin olun.

### Önkoşullar

Bu projeyi çalıştırmak için aşağıdaki yazılımların yüklü olduğundan emin olun:

- Golang: [Golang Resmi Websitesi](https://golang.org/)
- Redis: [Redis Resmi Websitesi](https://redis.io/)
- PostgreSQL: [PostgreSQL Resmi Websitesi](https://www.postgresql.org/)

### Kurulum

1. Projeyi yerel makinenize klonlayın:

   ```shell
   git clone https://github.com/FurkanSamaraz/Golang-Chat.git
   
2. Proje dizinine gidin:
    ```
    cd Golang-Chat
3. Gerekli bağımlılıkları yükleyin:
   go get
   
4. .env adında bir dosya oluşturun ve aşağıdaki içeriği ekleyin:
   REDIS_URL=redis://localhost:6379
   POSTGRES_URL=postgresql://username:password@localhost:5432/database_name?sslmode=disable
   
   Not: username, password ve database_name değerlerini kendi PostgreSQL ayarlarınıza göre güncelleyin.
   
5. Uygulamayı başlatmak için aşağıdaki komutu çalıştırın:
   go run main.go

6. Tarayıcınızda http://localhost:8080 adresini açın.
7. Uygulamayı kullanmaya başlayın!

## API Rotaları

- POST /register: Yeni bir kullanıcı kaydı oluşturur.
- POST /login: Bir kullanıcıyı giriş yapar.
- GET /verify-contact: Bir kullanıcının iletişim bilgilerini doğrular.
- GET /chat-history: İki kullanıcı arasındaki sohbet geçmişini alır.
- GET /contact-list: Bir kullanıcının iletişim listesini alır.
- GET /ws: WebSocket bağlantısını sağlar.
      
      Not: API rotalarıyla ilgili daha fazla ayrıntıyı swagger.yaml dosyasında bulabilirsiniz.


### Katkıda Bulunma
Katkıda bulunmak isteyen geliştiriciler, fork yaparak kendi projelerinde çalışabilirler. Yapılan değişiklikleri bir talep (pull request) ile gönderebilirler.
   







