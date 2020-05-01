FROM golang

ADD . /go/src/go-sugar
WORKDIR /go/src/go-sugar
RUN go get ./
RUN go install
# RUN go get -u github.com/gin-gonic/gin
# RUN go get github.com/dgrijalva/jwt-go
# RUN go get github.com/go-sql-driver/mysql
# RUN go get github.com/joho/godotenv
EXPOSE 9000
CMD go run main.go
