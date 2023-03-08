# blog_api_go

## 技術構成
* go
* gorilla
* mysql


```
function hello(){
   return "hello world!";
}
```

```
docker-compose up
```

データベースに接続し、プログラムを起動する。
```
go run main.go DB_USER=docker DB_PASSWORD=docker DB_NAME=sampledb
```

データベースにテストデータを投入する。
```
mysql -h 127.0.0.1 -u docker sampledb -p < repositories/testdata/setupDB.sql
```
パスワードに「docker」と入力する。


## 動作確認

POST /article をテスト
```
curl http://localhost:8080/article -X POST -d '{"title":"a","contents":"b","user_name":"c"}'
```
任意の記事を投稿する。

GET /article/list をテスト
```
curl http://localhost:8080/article/list -X GET
```
投稿した記事が取得できていることを確認する。

GET /article/id をテスト
```
curl http://localhost:8080/article/1 -X GET
```
指定したIDの記事が取得できていることを確認する。


POST /article/nice をテスト
```
curl http://localhost:8080/article/nice -X POST -d 
'{"article_id": 1,"title":"firstPost","contents": "This is my first blog","user_name": "user"}'
```
該当記事のniceを1増やす。

GET /article/id をテスト
```
curl http://localhost:8080/article/1 -X GET
```
該当記事のniceが1増えていることを確認する。

POST /comment をテスト
```
curl http://localhost:8080/comment -X POST -d '{"article_id": 1,"message": "テストコメント"}'
```
任意のコメントを投稿する。

GET /article/id をテスト
```
curl http://localhost:8080/article/1 -X GET
```
指定したIDの記事のコメントが取得できていることを確認する。



## データベース確認手順

ログイン

```
mysql -h 127.0.0.1 -u docker sampledb -p
```
パスワードに「docker」と入力する。

```
use sampledb;
```

記事テーブル確認

```
select * from articles;
```

コメントテーブル確認
```
select * from comments;
```




