package cpu

import (
	"easyMonitor/utils/number"
	"github.com/shirou/gopsutil/v3/cpu"
	"log"
)

type CpuInfo struct {
	// 厂商
	Vendor string
	// 核心数
	Cores int32
	// 型号
	ModelName string
	// 主频
	Mhz int
	// 标识
	Flags []string
}

func Info() (ci CpuInfo) {
	c, _ := cpu.Info()
	ci.Vendor = c[0].VendorID
	ci.Cores = c[0].Cores
	ci.ModelName = c[0].ModelName
	ci.Mhz = number.Round(c[0].Mhz)
	ci.Flags = c[0].Flags
	return
}

func init() {
	log.Printf(`【CPU信息】 厂商: "%s", 核心数: "%d", 型号: "%s", 主频: "%d", 标识: "%s"`, Info().Vendor, Info().Cores, Info().ModelName, Info().Mhz, Info().Flags)
}
