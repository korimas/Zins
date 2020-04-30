package main

import (
	"github.com/zpdev/zins/product/app"
	"github.com/zpdev/zins/product/url"
)

func main() {
	app.Init()
	url.Init()
	app.Run()
}
