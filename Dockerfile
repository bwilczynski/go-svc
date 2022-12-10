FROM golang:1.19-bullseye as builder

ENV CGO_ENABLED=0

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download -x
COPY . .
RUN go build -o /server ./cmd

FROM gcr.io/distroless/static

COPY --from=builder /server /server

EXPOSE 5000
EXPOSE 8000

ENTRYPOINT ["/server"]
CMD ["-port", "8000", "-admin-port", "5000"]
