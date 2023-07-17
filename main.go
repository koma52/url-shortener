package main

import (
	"github.com/koma52/url-shortener/backend"
)

func main() {
	a := backend.App{}
	a.Initialize()

	a.Run()

}
