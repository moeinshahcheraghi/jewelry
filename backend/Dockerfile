FROM docker.arvancloud.ir/golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod  .
RUN apk update && apk add --no-cache git
ARG http_proxy=http://45.94.254.18:43129
ARG https_proxy=http://45.94.254.18:43129
#RUN go mod download
RUN go mod download
COPY . .
RUN go build -o /app/main ./main.go
FROM docker.arvancloud.ir/alpine:latest
COPY --from=build /app/main /app/main
WORKDIR /app
EXPOSE 8080
CMD ["/app/main"]
