FROM golang:1.22 AS builder

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o /app/server cmd/**/*.go


FROM gcr.io/distroless/static-debian12 AS runtime
WORKDIR /app

COPY --from=builder --chown=nonroot:nonroot /app/server /server

ENTRYPOINT ["/server"]
