FROM golang:1.17-buster AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-lang-app

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /go-lang-app .

USER nonroot:nonroot

ENTRYPOINT ["/go-lang-app"]