FROM golang:latest as builder

COPY ./src/main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /main

FROM scratch

COPY --from=builder /main /main

EXPOSE 8000

ENTRYPOINT ["/main"]