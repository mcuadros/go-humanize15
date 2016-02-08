package humanize15

import (
	"testing"

	. "gopkg.in/check.v1"
	"gopkg.in/inconshreveable/log15.v2"
)

func Test(t *testing.T) { TestingT(t) }

type Humanize15Suite struct{}

var _ = Suite(&Humanize15Suite{})

func (s *Humanize15Suite) TestHumanizeHandler(c *C) {
	cases := [][]interface{}{
		[]interface{}{"none", 1024, 1024},
		[]interface{}{"rate", 1024, "1.0 kB/s"},
		[]interface{}{"speed", 1024, "1.0 kB/s"},
		[]interface{}{"size", 1024, "1.0 kB"},
		[]interface{}{"foo_size", 1024, "1.0 kB"},
		[]interface{}{"bytes", 1024, "1.0 kB"},
		[]interface{}{"duration", 1024, "1.024µs"},
		[]interface{}{"time", 1024, "1.024µs"},
		[]interface{}{"elapsed", 1024, "1.024µs"},
	}

	for _, t := range cases {
		s.testHumanizeHandler(c, t[0].(string), t[1], t[2])
	}
}

func (s *Humanize15Suite) testHumanizeHandler(c *C, key string, value, expected interface{}) {
	var record *log15.Record
	h := HumanizeHandler(log15.FuncHandler(func(r *log15.Record) error {
		record = r
		return nil
	}))

	h.Log(&log15.Record{Ctx: []interface{}{key, value}})
	c.Assert(record.Ctx[1], Equals, expected)
}
