FROM golang:1.18-alpine
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

#FROM golang:1.16-alpine AS build
#WORKDIR /go/src/app
#COPY --from=data /go/src/app .
#RUN go build -o /godo ./cmd/web 

# Run
#FROM alpine:latest
#COPY --from=build /godo /godo
#COPY --from=build /go/src/app/ui /ui
#COPY --from=build /go/src/app/tls /tls
RUN apk add --update gcc musl-dev git sqlite
COPY . .
RUN go build -v -o /godo ./cmd/web
EXPOSE 4000
CMD ["/godo"]
