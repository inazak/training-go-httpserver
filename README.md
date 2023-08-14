# training-go-httpserver

## 履歴

### v0.0.5

- sqlite3をgolang-migrateでUPする
- sqlite3にdate/time型がないので調整
- sqlxをラップするinterfaceの定義
- clockパッケージの追加
- DBへのIOを行うメソッドをRepository構造に追加

```
repository.go にある New() は、*sqlx.DB を返す、
それとは別に Repository構造の定義があり、これは Clocker を持っているだけ

```

直すべき点

```

```

### v0.0.4

- entity.Task でタスク構造の定義
- handler でタスクを追加する http.HandlerFunc を満たす ServeHTTP を作成
- タスクの登録時にリクエストボディのJSONを検証するため go-playground/validator を利用
- handler にJSONを返すヘルパー関数を追加
- goldenファイルによるテストを追加
- テストにおけるJSONの検証に go-cmp を利用
- タスクを一覧するエンドポイントの追加
- go-chi をルーティングに利用


```
entity.Task
がタスクの構造で

store.TaskStore = { LastID, map[TaskID]*Task }
               .Add( Task )
               .All() Tasks
がタスク操作の構造

storeに var Tasks で store.TaskStore がある状態

hander には store.TaskStore をさらにラップした
AddTask, ListTask があり、これは
.ServeHTTP をサポートしていて、muxでhandlerとして使われる
```
 
### v0.0.3

- シグナルを受け取る
- Server構造体に機能を分離してserver.goを作成
- mux.go を作成
 
### v0.0.2

- 起動ポートを環境変数から読み込み

### v0.0.1

- contextを利用したキャンセル可能なhttpserver起動
- testの作成


## 参考

- 詳解GO言語Webアプリケーション開発 9784863543720


