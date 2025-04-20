#!/bin/sh
# PostgreSQLを起動
su - postgres -c "pg_ctl -D /var/lib/postgresql/data -l /tmp/postgresql.log start"

# 少し待機してPostgreSQLが起動するのを待つ
sleep 5

# データベースの作成（初回のみ）
su - postgres -c "createdb -E UTF8 english_app || true"

# アプリケーション環境変数を設定
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=english_app
export DB_SSL_MODE=disable
export PORT=10000
export GIN_MODE=release

# アプリケーションを起動
/app/main