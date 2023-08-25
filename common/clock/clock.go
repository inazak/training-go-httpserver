package clock

import (
	"time"
)

// time.Time を使いたいが、sqlite3 に date/time の型がない
// string で記録する形に変更した

func NowString() string {
	location := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(location)
	return now.Format(time.RFC3339)
}
