# Download dependencies
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/sqlite
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get golang.org/x/time/rate
go get github.com/gorilla/websocket
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files

# Re-run swagger creation tools
swag init

# Run Docker
docker-compose up --build -d
