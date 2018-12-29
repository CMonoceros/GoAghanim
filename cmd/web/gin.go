package main

import (
	"cmonoceros.com/GoAghanim/pkg/lib"
	"cmonoceros.com/GoAghanim/pkg/web/router"
)

func main() {
	defer lib.CatchErrorFunc()()

	router.Init()
}
