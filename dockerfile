FROM golang:1.21 as builder

WORKDIR ./bin/url-services
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/main.go

FROM alpine:latest
WORKDIR .
COPY --from=builder /go/bin/url-services/main .
COPY --from=builder /go/bin/url-services/envs/ envs/
ENTRYPOINT ["./main", "pst"]