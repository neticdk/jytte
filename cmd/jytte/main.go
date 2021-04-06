package main

import (
	_ "net/http/pprof"

	"github.com/neticdk/jytte/pkg/cmd"
)

func main() {
	cmd.Execute()
}
