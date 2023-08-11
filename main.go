package main

import (
	_ "llmfarm/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"llmfarm/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.New())
}
