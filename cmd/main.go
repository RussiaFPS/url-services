package main

import (
	"os"
	"url-services/envs"
	"url-services/internal/app"
)

func init() {
	envs.LoadEnvs()
}

func main() {
	app.Run(os.Args[1])
}
