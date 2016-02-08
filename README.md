go-humanize15 [![Build Status](https://travis-ci.org/mcuadros/go-humanize15.png?branch=master)](https://travis-ci.org/mcuadros/go-humanize15) [![GoDoc](https://godoc.org/github.com/mcuadros/go-humanize15?status.png)](http://godoc.org/github.com/mcuadros/go-humanize15)
=============

log15.Handler that humanize the supported context keys to human readable strings

Installation
------------

The recommended way to install go-humanize15

```
go get -u github.com/mcuadros/go-humanize15
```

Examples
--------

```go
import (
    "os"

    "github.com/mcuadros/go-humanize15"
    "gopkg.in/inconshreveable/log15.v2"
)

func main() {
    log := log15.New()
    log.SetHandler(humanize15.HumanizeHandler(
        log15.StreamHandler(os.Stdout, log15.LogfmtFormat())
    ))

    log.Warn("foo", "rate", 1024, "elapsed", 1024)
}

//t=2016-02-08T00:13:32+0100 lvl=warn msg=foo rate="1.0 kB/s" elapsed="1.024µs"
```

License
-------

MIT, see [LICENSE](LICENSE)
