FROM golang:1-alpine as build
WORKDIR /src
RUN apk update && apk add --no-cache ca-certificates git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/webwatch-server
WORKDIR /src/cmd/cli
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/webwatch

FROM scratch
LABEL maintainer="Bogdan Kosarevskyi <bogdan.kosarevskyi@gmail.com>"
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app /app
EXPOSE 8080
CMD ["/app/webwatch-server"]