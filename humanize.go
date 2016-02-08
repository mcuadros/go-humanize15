package humanize15

import (
	"strings"
	"time"

	"github.com/dustin/go-humanize"

	"gopkg.in/inconshreveable/log15.v2"
)

// Kind defines the supported humanizables kinds
type Kind int

// Keyword defines the keywords where look for humanizable kinds
type Keyword string

const (
	// Rate any number is translete to a string like [x]kb/s
	Rate Kind = iota
	// Size any number is translated to a string like [x]kb
	Size
	// Duration any number is translated to a time.Duration
	Duration
)

// Supported defines how they keywords are translated
var Supported = map[Keyword]Kind{
	"rate": Rate, "speed": Rate,
	"size": Size, "bytes": Size,
	"duration": Duration, "time": Duration, "elapsed": Duration,
}

// HumanizeHandler returns a log15.Handler that humanize the supported context
// keys to human readable strings
func HumanizeHandler(h log15.Handler) log15.Handler {
	return log15.FuncHandler(func(r *log15.Record) error {
		l := len(r.Ctx)
		for i := 0; i < l; i = i + 2 {
			k, ok := r.Ctx[i].(string)
			if !ok {
				continue
			}

			if res, done := processKey(k, r.Ctx[i+1]); done {
				r.Ctx[i+1] = res
			}
		}

		return h.Log(r)
	})
}

func processKey(k string, v interface{}) (res interface{}, done bool) {
	keyword := getKeyword(k)
	if kind, ok := Supported[keyword]; ok {
		return applyHumanize(kind, v), true
	}

	return v, false
}

func getKeyword(k string) Keyword {
	parts := strings.Split(k, "_")
	l := len(parts)
	if l > 1 {
		return Keyword(parts[l-1])
	}

	return Keyword(parts[0])
}

func applyHumanize(kind Kind, v interface{}) interface{} {
	switch kind {
	case Size:
		return applyHumanizeSize(v)
	case Rate:
		return applyHumanizeRate(v)
	case Duration:
		return applyHumanizeDuration(v)
	}

	return nil
}

func applyHumanizeRate(v interface{}) interface{} {
	if v, ok := applyHumanizeSize(v).(string); ok {
		return v + "/s"
	}

	return v
}

func applyHumanizeSize(v interface{}) interface{} {
	if bytes, ok := castToUint64(v); ok {
		return humanize.Bytes(bytes)
	}

	return v
}

func applyHumanizeDuration(v interface{}) interface{} {
	if bytes, ok := castToUint64(v); ok {
		return time.Duration(bytes).String()
	}

	return v
}

func castToUint64(v interface{}) (uint64, bool) {
	var u64 uint64
	switch v.(type) {
	case int:
		u64 = uint64(v.(int))
	case int8:
		u64 = uint64(v.(int8))
	case int16:
		u64 = uint64(v.(int16))
	case int32:
		u64 = uint64(v.(int32))
	case int64:
		u64 = uint64(v.(int64))
	case uint8:
		u64 = uint64(v.(uint8))
	case uint16:
		u64 = uint64(v.(uint16))
	case uint32:
		u64 = uint64(v.(uint32))
	case uint64:
		u64 = v.(uint64)
	case float32:
		u64 = uint64(v.(float32))
	case float64:
		u64 = uint64(v.(float64))
	default:
		return 0, false
	}

	return u64, true
}
