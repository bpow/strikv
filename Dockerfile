FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum *.go ./
RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o /strikv

FROM gcr.io/distroless/base-debian12 as release
WORKDIR /
COPY --from=builder /strikv /strikv
EXPOSE 8080
ENTRYPOINT ["/strikv"]