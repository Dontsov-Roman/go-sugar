FROM golang:1.10.4

WORKDIR /home/fallengel/go-sugar
COPY . .
RUN go get github.com/gin-gonic/gin
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/go-sql-driver/mysql
EXPOSE 9000
CMD go run main.go
