FROM golang:1.26.5-bookworm AS builder

WORKDIR /src

ARG TARGET=api

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN case "$TARGET" in \
        api|walk-worker|outbox-worker) CGO_ENABLED=0 go build -o /out/app "./cmd/$TARGET" ;; \
        *) echo "unsupported TARGET: $TARGET (expected api, walk-worker, or outbox-worker)" >&2; exit 1 ;; \
    esac

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder --chown=nonroot:nonroot /out/app /app/app

EXPOSE 8080

USER nonroot:nonroot

CMD ["/app/app"]
