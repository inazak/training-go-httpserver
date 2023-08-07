package store

/*
他のRDBでは `ALTER TABLE` でテーブル構造を変更できる。
しかし、sqliteでは `ALTER TABLE` では下記の2点しか出来ない。

- テーブル名の変更
- カラムの追加

他の操作、つまりカラム名の変更や削除、データ型の変更などは
`ALTER TABLE` で行えない。

そのような場合は、次のステップで実行することになる。

1. 既存のテーブル名を変更する
2. 新しいテーブルを、元のテーブル名で作成する
3. 新しいテーブルに、元のテーブルのデータをコピーする
4. 元のテーブルを `DROP` する
*/

import (
  "database/sql"
  "embed"

  "github.com/pkg/errors"
  "github.com/golang-migrate/migrate/v4"
  "github.com/golang-migrate/migrate/v4/database/sqlite3"
  "github.com/golang-migrate/migrate/v4/source/iofs"
  _ "github.com/mattn/go-sqlite3"
)

//go:embed script/*.sql
var fs embed.FS

func RunMigrateUp(db *sql.DB) error {

  dbdriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
  if err != nil {
    return errors.Wrap(err, "creating sqlite3 db driver failed")
  }

  fsdriver, err := iofs.New(fs, "script")
  if err != nil {
    return errors.Wrap(err, "creating filesystem driver failed")
  }

  m, err := migrate.NewWithInstance("iofs", fsdriver, "sqlite3", dbdriver)
  if err != nil {
    return errors.Wrap(err, "initializing db migration failed")
  }

  err = m.Up()
  if err != nil && err != migrate.ErrNoChange {
    return errors.Wrap(err, "migrate database failed")
  }

  return nil
}


