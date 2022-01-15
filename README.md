# memer

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

```sh
# .go-version に従いインストール
$ goenv install

# 環境変数に追加
$ cp .envrc.dev .envrc
$ vim .envrc
$ direnv allow .

# dockerでMySQL起動
$ docker-compose up -d rdb

# マイグレーションの実行
$ sql-migrate up -config asset/db/dbconf.yml

# アプリ起動
$ air
```
