package main

import (
	"easyMonitor/pkg/cpu"
	"easyMonitor/pkg/mem"
	"easyMonitor/pkg/network"
	"log"
)

func main() {
	cpuItem := cpu.Item()
	memInfo := mem.Item()

	cpuInfo := cpu.Info()

	log.Printf(`【CPU信息】 厂商: "%s", 核心数: "%d", 型号: "%s", 主频: "%d", 标识: "%s", 使用率: "%.2f%%"`, cpuInfo.Vendor, cpuInfo.Cores, cpuInfo.ModelName, cpuInfo.Mhz, cpuInfo.Flags, cpuItem.UsedPercentTotal)
	log.Printf(`【内存信息】 内存总量: "%d", 内存使用率: "%.2f%%", swap总量: "%d", swap空闲大小: "%d"`, memInfo.Total, memInfo.UsedPercent, memInfo.SwapTotal, memInfo.SwapFree)

	network.Info()
}
