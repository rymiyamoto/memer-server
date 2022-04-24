# memer-server

## Getting Started

事前に以下のものを入れてください

|項目|目的|
|-|-|
|[goenv](https://github.com/syndbg/goenv)|goのバージョン管理|
|[sql-migrate](https://github.com/rubenv/sql-migrate)|マイグレーションで利用|
|[direnv](https://github.com/direnv/direnv)|.envrc利用のため|
|[golangci-lint](https://github.com/golangci/golangci-lint)|linter|
|[pre-commit](https://pre-commit.com)|commit時の事前チェック|
|[air](https://github.com/cosmtrek/air)|hot reload|
|[mockgen](https://github.com/golang/mock)|mock作成|
|[swag](https://github.com/swaggo/swag)|swaggerの生成|
|[echo-swagger](https://github.com/swaggo/echo-swagger)|swaggerのecho周りの補完|

```sh
# .go-version に従いインストール
$ goenv install

# 環境変数に追加
$ cp .envrc.dev .envrc
$ vim .envrc
$ direnv allow .

# pre-commit の追加
$ pre-commit install

# dockerでMySQL起動
$ docker-compose up -d rdb

# マイグレーションの実行
$ sql-migrate up -config asset/db/dbconf.yml

# アプリ起動
$ air
```

## Testing

### exec

```sh
# 対象のディレクトリに移動
$ cd target_dir

# ディレクトリ内すべて
$ go test -run ""

# 特定のテストのみ
$ go test -run "TestHoge"
```

### mock

```sh
# mock作成
$ mockgen -source=hoge.go -destination=hoge_mock.go -package=pkg
```

## Swagger

### generate

[format](https://github.com/swaggo/swag#general-api-info)

```sh
# ドキュメント生成
$ swag init
```

### URL

|env|url|
|-|-|
|local|http://localhost:1324/swagger/index.html|
