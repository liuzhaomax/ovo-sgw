package main

import (
	"context"
	"github.com/liuzhaomax/ovo/internal/app"
	"github.com/liuzhaomax/ovo/internal/core"
)

func main() {
	app.Launch(
		context.Background(),
		app.SetConfigFile(core.LoadEnv()),
		app.SetWWWDir("www"),
	)
}
