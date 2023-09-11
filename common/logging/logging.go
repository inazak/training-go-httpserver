package logging

import (
	"github.com/go-kit/log"
	"os"
	"time"
)

var (
	timestampFormat = log.TimestampFormat(
		func() time.Time {
			return time.Now().UTC().In(time.FixedZone("Asia/Tokyo", 9*60*60))
		},
		"2006-01-02T15:04:05.000Z07:00",
	)
)

func NewLogger() log.Logger {
	l := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	l = log.With(l, "ts", timestampFormat, "caller", log.DefaultCaller)
	return l
}
