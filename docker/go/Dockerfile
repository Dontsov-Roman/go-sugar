FROM golang

ADD . /go/src/go-sugar
WORKDIR /go/src/go-sugar
RUN go get ./
RUN go install
CMD go run main.go
