package main

import (
	_ "gf_playground/internal/logic"
	_ "gf_playground/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"gf_playground/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
