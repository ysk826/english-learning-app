# English Learning App - Backend

英語学習を支援するアプリケーションのバックエンド API サーバーです。ユーザー認証、単語の保存、文章の作成などの機能を提供します。

## 技術スタック

-   Go 1.23
-   Gin (Web フレームワーク)
-   GORM (ORM ライブラリ)
-   PostgreSQL (データベース)
-   JWT (認証)

## 機能

-   ユーザー登録・ログイン
-   単語の保存と取得
-   文章の作成と取得

## 開発環境のセットアップ

### 前提条件

-   Go 1.23 以上
-   Docker および Docker Compose
-   PostgreSQL

### インストールと実行

1. リポジトリのクローン
   git clone https://github.com/ysk826/english-learning-app
   cd english-learning-app

2. 環境変数の設定
   cp .env.example .env
   必要に応じて`.env`ファイルを編集してください。

3. Docker Compose を使用して開発環境を起動
   docker-compose up

4. バックエンドのみを起動（Docker Compose なし）
   cd backend
   go run cmd/api/main.go

### テスト

テストを実行するには以下のコマンドを使用します：
cd backend
go test ./...

特定のパッケージのみテストを実行：
go test ./internal/services

## API エンドポイント

### 認証

-   `POST /api/v1/auth/register` - ユーザー登録
-   `POST /api/v1/auth/login` - ログイン

### 単語

-   `GET /api/v1/words` - 単語リスト取得
-   `POST /api/v1/words` - 新しい単語を追加

### 文章

-   `GET /api/v1/sentences` - 文章リスト取得
-   `POST /api/v1/sentences` - 新しい文章を追加

## デプロイ

アプリケーションは Render を使用してデプロイされています。main ブランチへのプッシュが行われると、GitHub Actions によって自動的にテストとデプロイが実行されます。

## プロジェクト構造

backend/
├── cmd/ # エントリーポイント
│ ├── api/ # API サーバー
│ └── migrate/ # マイグレーションツール
├── db/ # データベース関連
│ ├── migrations/ # マイグレーションファイル
│ └── seeds/ # シードデータ
├── internal/ # 内部パッケージ
│ ├── api/ # API ハンドラー
│ ├── config/ # 設定
│ ├── database/ # データベース接続
│ ├── models/ # データモデル
│ ├── repository/ # データアクセス層
│ ├── services/ # ビジネスロジック層
│ └── utils/ # ユーティリティ
└── .air.toml # Air 設定（ホットリロード）
