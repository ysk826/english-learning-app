# バックエンドアプリとPostgreSQLの両方を含む単一のDockerfile
FROM golang:1.23-alpine

WORKDIR /app

# PostgreSQLのインストール
RUN apk add --no-cache postgresql postgresql-contrib

# アプリケーションのビルド
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/api

# PostgreSQL初期化
USER postgres
RUN mkdir -p /run/postgresql && chown -R postgres:postgres /run/postgresql/
RUN mkdir -p /var/lib/postgresql/data && chmod 0700 /var/lib/postgresql/data
RUN initdb -D /var/lib/postgresql/data

# PostgreSQL設定
RUN echo "host all all 0.0.0.0/0 md5" >> /var/lib/postgresql/data/pg_hba.conf
RUN echo "listen_addresses='*'" >> /var/lib/postgresql/data/postgresql.conf

# 起動スクリプトの追加
USER root
COPY start.sh /start.sh
RUN chmod +x /start.sh

EXPOSE 5432
EXPOSE 10000

CMD ["/start.sh"]