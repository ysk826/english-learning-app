#!/bin/sh
# PostgreSQLを起動
su - postgres -c "pg_ctl -D /var/lib/postgresql/data -l /var/log/postgresql/postgresql.log start"

# データベースの作成（初回のみ）
su - postgres -c "createdb -E UTF8 english_app || true"

# アプリケーションを起動
/app/main