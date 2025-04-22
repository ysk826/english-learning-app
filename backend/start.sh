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

# 実行ファイルの確認
echo "現在のディレクトリ:"
pwd
echo "ファイル一覧:"
ls -la

# 実行ファイルが存在するか確認
if [ -f "/app/main" ]; then
    echo "実行ファイルが存在します"
    # アプリケーションを起動
    exec /app/main
else
    echo "実行ファイルが見つかりません"
    exit 1
fi