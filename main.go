package main

import (
	_ "blog-api/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"blog-api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
