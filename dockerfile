FROM golang:1.9.1 as builder
WORKDIR /go/src/github.com/martin-magakian/bidder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bidder .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/martin-magakian/bidder/bidder .
EXPOSE 3001
CMD ["./bidder"]