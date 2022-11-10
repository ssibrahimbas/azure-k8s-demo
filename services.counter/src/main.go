package main

import "github.com/ssibrahimbas/azure-k8s-demo.counter/src/app"

func main() {
	app.New().Init().Serve()
}