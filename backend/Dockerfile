# alpineベースのGo開発環境をマルチステージビルドで構築
FROM golang:1.23-alpine AS builder

# ルートの直下に作業ディレクトリを作成
WORKDIR /app

# # ホストマシン上にあるgo.mod go.sumをコンテナの/appにコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
# 第一引数(コピー元)、第二引数(コピー先)を指定
# COPY . .は、ホストマシン上の全てのファイルをコンテナの/appにコピーする
COPY . .

# アプリケーションをビルド
# -oオプションで、出力ファイル名を指定(main)
# ./cmd/apiは、ビルドするGoファイルのパスを指定
RUN go build -o main ./cmd/api

# alpineベースのPostgreSQL環境をマルチステージビルドで構築
# 最終的なイメージは、PostgreSQLとGoアプリケーションを含む
FROM postgres:15-alpine

# apkを使用してcurlとtzdataをインストール
# tzdataは、タイムゾーンデータを提供するパッケージ
# --no-cacheオプションは、キャッシュを使用せずにインストールすることを指定
# これにより、イメージのサイズを小さく保つことができる
# tzdataは、PostgreSQLのタイムゾーン設定に必要
RUN apk add --no-cache curl tzdata

# PostgreSQL用のディレクトリを準備
# /run/postgresqlは、PostgreSQLのソケットファイルを格納するためのディレクトリ
# /var/lib/postgresql/dataは、PostgreSQLのデータファイルを格納するためのディレクトリ
# chmod 0700は、ディレクトリのパーミッションを設定するためのコマンド
RUN mkdir -p /run/postgresql && \
    mkdir -p /var/lib/postgresql/data && \
    chown -R postgres:postgres /run/postgresql/ && \
    chown -R postgres:postgres /var/lib/postgresql/data && \
    chmod 0700 /var/lib/postgresql/data

# 実行ユーザーをpostgresに変更
# -Dオプションは、PostgreSQLのデータディレクトリを指定するためのオプション
USER postgres
RUN initdb -D /var/lib/postgresql/data

# PostgreSQL設定
RUN echo "host all all 0.0.0.0/0 md5" >> /var/lib/postgresql/data/pg_hba.conf && \
    echo "listen_addresses='*'" >> /var/lib/postgresql/data/postgresql.conf

# 実行ユーザーをrootに変更
USER root

# ルートの直下に作業ディレクトリを作成(2番目のステージ)
WORKDIR /app

# ビルダーステージからビルドしたアプリケーションをコピー
COPY --from=builder /app/main ./

# マイグレーション用のディレクトリを作成
RUN mkdir -p /app/db/migrations

# ビルダーステージからマイグレーションファイルをコピー
COPY --from=builder /app/db/migrations /app/db/migrations

# 起動スクリプトをコピー
COPY start.sh /start.sh

# スクリプトの実行権限を指定したパスに付与
# シェルスクリプトは、実行権限が必要
RUN chmod +x /start.sh

# コンテナが5432ポートでリッスンするように設定
# PostgreSQLのデフォルトポートは5432
# 10000ポートは、Goアプリケーションのポート
EXPOSE 5432
EXPOSE 10000

# 環境変数
ENV PORT=10000
ENV GIN_MODE=release

# start.shを実行
CMD ["/start.sh"]