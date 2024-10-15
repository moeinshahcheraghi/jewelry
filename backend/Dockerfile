FROM golang:1.21-alpine AS build
WORKDIR /src
COPY go.mod  .
RUN go mod download
COPY . .
RUN go build -o /app/main ./main.go

FROM alpine:latest
COPY --from=build /app/main /app/main
WORKDIR /app
EXPOSE 8080
CMD ["/app/main"]
