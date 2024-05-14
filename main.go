package main

import (
	"github.com/yurifrl/sturdy-succotash/cmd"

	_ "github.com/lib/pq"
)

func main() {
	cmd.Execute()
}
