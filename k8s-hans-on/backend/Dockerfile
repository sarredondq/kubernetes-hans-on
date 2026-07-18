# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY src/main.go ./main.go

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/server \
    ./main.go

FROM gcr.io/distroless/static-debian12:nonroot AS runtime

COPY --from=builder --chown=nonroot:nonroot /out/server /server

EXPOSE 9090

USER nonroot:nonroot

ENTRYPOINT ["/server"]
