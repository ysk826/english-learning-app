# バックエンドアプリとPostgreSQLの両方を含む単一のDockerfile
FROM golang:1.23-alpine

WORKDIR /app

# PostgreSQLのインストール
RUN apk add --no-cache postgresql postgresql-contrib

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

# Goアプリケーションのビルド
USER root
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/api

# 起動スクリプトの追加
COPY start.sh /start.sh
RUN chmod +x /start.sh

EXPOSE 5432
EXPOSE 10000

CMD ["/start.sh"]