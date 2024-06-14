package main

import (
	"context"
	"github.com/liuzhaomax/ovo-sgw/internal/app"
	"github.com/liuzhaomax/ovo-sgw/internal/core"
)

func main() {
	app.Launch(
		context.Background(),
		app.SetConfigFile(core.LoadEnv()),
		app.SetWWWDir("www"),
	)
}
