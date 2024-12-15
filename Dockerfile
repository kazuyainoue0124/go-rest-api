# ビルドステージ
FROM golang:1.23-alpine AS builder
WORKDIR /app

# go.mod, go.sumのみを先にコピーし、依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコード全体をコピー
COPY . .

# CGOを無効化（特にC依存がなければ推奨）、バイナリサイズ縮小のためのldflagsも指定
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o server ./cmd/server

# 本番実行用ステージ
FROM alpine:3.21 AS final
WORKDIR /app

# ca証明書とユーザー作成（非root実行）
RUN apk add --no-cache ca-certificates \
    && addgroup -S appgroup \
    && adduser -S appuser -G appgroup

COPY --from=builder /app/server /app/server

USER appuser:appgroup

EXPOSE 8080
CMD ["./server"]
