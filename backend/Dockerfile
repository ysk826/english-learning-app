FROM golang:1.23-alpine AS builder

WORKDIR /app

# 依存関係をコピーしてダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main ./cmd/api

# PostgreSQLを含む最終イメージ
FROM postgres:15-alpine

# PostgreSQLのインストール
# Goアプリケーションの実行に必要なツール
RUN apk add --no-cache curl tzdata

# PostgreSQL用のディレクトリを準備（rootユーザーで実行）
RUN mkdir -p /run/postgresql && \
    mkdir -p /var/lib/postgresql/data && \
    chown -R postgres:postgres /run/postgresql/ && \
    chown -R postgres:postgres /var/lib/postgresql/data && \
    chmod 0700 /var/lib/postgresql/data

# PostgreSQL初期化
USER postgres
RUN initdb -D /var/lib/postgresql/data

# PostgreSQL設定
RUN echo "host all all 0.0.0.0/0 md5" >> /var/lib/postgresql/data/pg_hba.conf && \
    echo "listen_addresses='*'" >> /var/lib/postgresql/data/postgresql.conf

# Goアプリケーションをコピー
USER root
WORKDIR /app
COPY --from=builder /app/main .

# マイグレーションファイルをコピー
COPY --from=builder /app/../db/migrations /app/db/migrations

# 起動スクリプトをコピー
COPY start.sh /start.sh
RUN chmod +x /start.sh

# ポート設定
EXPOSE 5432
EXPOSE 10000

# 環境変数
ENV PORT=10000
ENV GIN_MODE=release

CMD ["/start.sh"]