# training-go-httpserver

## 履歴

### v0.0.15

- go fmt する

### v0.0.14

- httpserver と service の結合
- main.go の最低限の動作確認


### v0.0.13

- service の追加
- httpserver/handler/api の追加
- common/mock の生成


### v0.0.12

- repository/database/sqlite の追加
- model の追加


### v0.0.11

- v0.0.6 から再構成
- go-kit/log の導入
- パッケージの再構成


### v0.0.6

- ハンドラにDBとのIOをserviceパッケージ経由で紐付け
- gomockでのモック生成、モックを使ったテスト
- clockをtime.Timeではなくstringに変更
- handler は store ではなく、Serviceインターフェイスを利用する


gomock使うために
```
go install github.com/golang/mock/mockgen@v1.6.0
```

この辺はもう分かりにくいので書き出すと
```
handler パッケージ
  service.go に AddTaskService などが interface として定義されており
  これらの実装は service パッケージの add_task.go など

service パッケージ
  interface.go TaskAdder などがinterfaceとして定義されており
  これらの実装は storeパッケージのデータベースI/O


handler.AddTaskService = (service.AddTask)AddTask
service.TaskAddr = (store.Repository)AddTask
となっているけど、

service/add_task.go で
  struct AddTask { store.Execer, TaskAdder } を定義していて
  そのAddTaskのメソッドとして、AddTaskを定義していて、
  メソッド内で TaskAdderつまり store.RepositoryのAddTaskを実行している
    func (a *AddTask) AddTask(){ ... Taskadder.AddTask(ctx,db,..) ... }
  名前の使い方が分かりにくいレベルを超えている
  
mux.go でこうなっているのも重すぎる
  r := store.Repository{}
  at := &handler.AddTask{
    Service: &service.AddTask{ *sqlx.DB, &r }
  }
```


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

を参考にして始めたが、途中 v0.0.11から全然違う道を進んでしまった


## TODO

- 今回使ったパッケージの使用例部分の抜粋と説明の抜き書き


