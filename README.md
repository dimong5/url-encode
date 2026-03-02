# url-encode

Сервис для создания коротких ссылок

### Docker

docker-compose up --build

### Создать короткую ссылку

curl -X POST "http://localhost:8080/shorten?url=https://google.com"

### Тесты

go test ./...