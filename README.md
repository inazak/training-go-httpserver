# training-go-httpserver

## 履歴

### v0.0.4

- entity.Task でタスク構造の定義
- handler でタスクを追加する http.HandlerFunc を満たす ServeHTTP を作成
- タスクの登録時にリクエストボディのJSONを検証するため go-playground/validator を利用
- handler にJSONを返すヘルパー関数を追加
- goldenファイルによるテストを追加
- テストにおけるJSONの検証に go-cmp を利用
- タスクを一覧するエンドポイントの追加
- go-chi をルーティングに利用

 
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


