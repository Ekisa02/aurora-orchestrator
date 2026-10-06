FROM golang:1.20-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://proxy.golang.org
COPY . .
RUN go build -o /usr/local/bin/aurora ./cmd/server

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /usr/local/bin/aurora /usr/local/bin/aurora
EXPOSE 8080
CMD ["/usr/local/bin/aurora"]
