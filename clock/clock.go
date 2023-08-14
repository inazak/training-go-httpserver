package clock

import (
  "time"
)

// time.Time を使いたいが、sqlite3 に date/time の型がない
// string で記録する形に変更した

type Clocker interface {
  Now() string
}

type RealClocker struct {}

func (r RealClocker) Now() string {
  location := time.FixedZone("Asia/Tokyo", 9*60*60)
  now := time.Now().In(location)
  return now.Format(time.RFC3339)
}

type FixedClocker struct {}

func (f FixedClocker) Now() string {
  location := time.FixedZone("Asia/Tokyo", 9*60*60)
  now := time.Date(2023, 4, 5, 12, 34, 56, 0, time.UTC).In(location)
  return now.Format(time.RFC3339)
}

