FROM golang:1.18-alpine
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN apk add --update gcc musl-dev git sqlite
COPY . .

RUN go build -v -o /godo ./cmd/web
EXPOSE 4000
CMD ["/godo"]
