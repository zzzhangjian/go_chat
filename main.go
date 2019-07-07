package main

import (
	_ "github.com/zzzhangjian/go_chat/boot"
	_ "github.com/zzzhangjian/go_chat/router"
	"github.com/gogf/gf/g"
)

func main() {
	g.Server().Run()
}