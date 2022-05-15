package main

import (
	"easyMonitor/pkg/cpu"
	"easyMonitor/pkg/mem"
	"log"
	"runtime"
)

func main() {
	cpu.Info()
	mem.Info()
	log.Println(runtime.GOOS)
}
