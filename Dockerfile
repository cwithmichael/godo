FROM golang:1.16-alpine AS build
WORKDIR /go/src/app

COPY . .
RUN go mod download

RUN go build -o /godo ./cmd/web 

# Deploy
FROM alpine:latest
COPY --from=build /godo /godo
COPY --from=build /go/src/app/ui /ui
COPY --from=build /go/src/app/tls /tls
EXPOSE 4000
CMD ["./godo"]
