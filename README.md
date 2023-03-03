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

```
use sampledb;
```

```
select * from articles;
```

```
select * from comments;
```
