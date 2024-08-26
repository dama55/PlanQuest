# backend/Dockerfile

# Goの公式イメージを使用
FROM golang:1.23

# 作業ディレクトリを設定
WORKDIR /usr/src/app/my-go-backend

RUN apt update && apt install -y bash

# Goの依存関係をキャッシュするためのステップ
COPY ./my-go-backend/go.mod ./
COPY ./my-go-backend/go.sum ./
RUN go mod download
RUN go mod tidy

# Airをインストール
RUN go install github.com/air-verse/air@latest

# Goモジュールファイルをコピー（まだプロジェクトがない場合はコメントアウト可）
# COPY go.mod go.sum ./

# Goモジュールをダウンロード（まだ依存関係がない場合はコメントアウト可）
# RUN go mod download

# プロジェクトディレクトリ全体をコピー
# COPY . .

# アプリケーションをビルド・実行（開発環境が整ったら使用）
# RUN go build -o main .
# CMD ["./main"]

# 開発が整うまで、基本のシェルを起動しておく
CMD ["tail", "-f", "/dev/null"]