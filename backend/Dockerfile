FROM golang:1.23-alpine AS builder

WORKDIR /app

# 依存関係をコピーしてダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main ./cmd/api

# 最終イメージ
FROM alpine:latest

WORKDIR /app

# ビルダーからバイナリをコピー
COPY --from=builder /app/main .

# ポート設定
EXPOSE 10000

# アプリケーションを実行
CMD ["./main"]