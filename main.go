package main

import (
	_ "gf_playground/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf_playground/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
