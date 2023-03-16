# blog_api_go

## 概要
記事とコメントを投稿する機能を持ったAPIです。
ターミナルからコマンドを実行して動作確認を行います。

## 技術構成
* go1.19
* docker
* mysql5.7


## API
|    |メソッド|URI|
| --------- | ----------- | ------- |
|記事投稿|POST|/article|
|全記事データを取得|GET|/article/list|
|指定したIDの記事取得|GET|/article/:id|
|指定した記事のniceを1増やす|POST|/article/nice|
|指定した記事にコメント投稿|POST|/comment|


## DB事前準備
※GoとDockerはインストール済みの前提で記載しています。
未インストールの際は下記ページ参照の上インストールして下さい。

- Goインストール
　https://go.dev/dl/

- Dockerインストール
　https://matsuand.github.io/docs.docker.jp.onthefly/get-docker/

Dockerコンテナを起動する。
```
docker-compose up
```

データベースに接続し、プログラムを起動する。
```
go run main.go DB_USER=docker DB_PASSWORD=docker DB_NAME=sampledb
```

データベースにテストデータを投入する。

※データベースに接続する際、ポート3306を使用しているため、既に使用している場合killする必要があると思われます。
```
mysql -h 127.0.0.1 -u docker sampledb -p < repositories/testdata/setupDB.sql
```
パスワードに「docker」と入力する。


## API動作確認

### 正常系

・POST /article をテスト
```
curl http://localhost:8080/article -X POST -d '{"title":"a","contents":"b","user_name":"c"}'
```
任意の記事を投稿する。
<br>
<br>
・GET /article/list をテスト
```
curl http://localhost:8080/article/list -X GET
```
投稿した記事が取得できていることを確認する。
<br>
<br>
・GET /article/id をテスト
```
curl http://localhost:8080/article/1 -X GET
```
指定したIDの記事が取得できていることを確認する。
<br>
<br>
・POST /article/nice をテスト
```
curl http://localhost:8080/article/nice -X POST -d 
'{"article_id": 1,"title":"firstPost","contents": "This is my first blog","user_name": "user"}'
```
該当記事のniceを1増やす。

GET /article/id で確認
```
curl http://localhost:8080/article/1 -X GET
```
該当記事のniceが1増えていることを確認する。
<br>
<br>
・POST /comment をテスト
```
curl http://localhost:8080/comment -X POST -d '{"article_id": 1,"message": "テストコメント"}'
```
任意のコメントを投稿する。

GET /article/id で確認
```
curl http://localhost:8080/article/1 -X GET
```
指定したIDの記事のコメントが取得できていることを確認する。


### 異常系

下記コマンドでデータベースサーバを止めておく。
```
docker-compose down
```

・GET /article/list をテスト
```
curl http://localhost:8080/article/list -X GET
```
{"ErrCode":"S002","Message":"fail to get data"}が表示される。
<br>
<br>
・GET /article/id をテスト
```
curl http://localhost:8080/article/1 -X GET
```
{"ErrCode":"S002","Message":"fail to get data"}が表示される。




登録系操作確認のため下記コマンドでデータベースサーバを動作させる。
```
docker-compose up
```
・POST /article をテスト
```
curl http://localhost:8080/article -X POST -d '{"title":"a","contents":"b","user_name":"c",test}'
```
{"ErrCode":"R001","Message":"bad request body"}が表示される。
<br>
<br>
・POST /article/nice をテスト
```
curl http://localhost:8080/article/nice -X POST -d '{"article_id": 1,"title":"firstPost","contents": "This is my first blog","user_name": "user",test}'
```
{"ErrCode":"S004","Message":"does not exist target article"}が表示される。
<br>
<br>
・POST /comment をテスト
```
curl http://localhost:8080/comment -X POST -d '{"article_id": 1,"message": "テストコメント",test}'
```
{"ErrCode":"R001","Message":"bad request body"}
{"ErrCode":"S001","Message":"fail to record data"}が表示される。


## データベース確認手順

データベースにログインする。

```
mysql -h 127.0.0.1 -u docker sampledb -p
```
パスワードに「docker」と入力する。

```
use sampledb;
```

記事テーブルを確認する。

```
select * from articles;
```

コメントテーブルを確認する。
```
select * from comments;
```
