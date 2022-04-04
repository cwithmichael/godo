# Build
FROM golang:1.18-alpine AS build
WORKDIR /go/src/app

COPY . .
RUN go mod download
RUN apk add git
RUN go build -o /godo ./cmd/web 

# Run
FROM alpine:latest
COPY --from=build /godo /godo
COPY --from=build /go/src/app/ui /ui
COPY --from=build /go/src/app/tls /tls
EXPOSE 4000
CMD ["./godo"]
