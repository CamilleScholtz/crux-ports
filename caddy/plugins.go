package main

import (
	"github.com/mholt/caddy/caddy/caddymain"
	_ "github.com/onodera-punpun/punfed"
)

func main() {
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
