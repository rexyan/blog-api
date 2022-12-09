package main

import (
	_ "blog-api/internal/logic"
	_ "blog-api/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"blog-api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
