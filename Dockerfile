FROM golang:1.16 AS build-image
WORKDIR /root/
COPY . .
RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.13.2
WORKDIR /root/
COPY --from=build-image /root/main .
EXPOSE 3000
CMD ["./main"]
