package main

import (
	"biuaxia.cn/bb/code/core"
	"biuaxia.cn/bb/code/support"
)

func main() {
	core.APPLICATION = &support.BbApplication{}
	core.APPLICATION.Start()
}
