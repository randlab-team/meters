FROM golang:1.17
ADD ./ /app/
WORKDIR /app
CMD go mod download -x
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/meters_service cmd/meters_service/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/build/meters_service ./

EXPOSE 8080

CMD ["./meters_service"]