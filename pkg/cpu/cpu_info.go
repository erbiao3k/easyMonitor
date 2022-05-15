package cpu

import (
	"easyMonitor/utils/number"
	"github.com/shirou/gopsutil/v3/cpu"
	"log"
)

// CpuInfo CPU信息
type CpuInfo struct {
	Vendor    string   // 厂商
	Cores     int32    // 核心数
	ModelName string   // 型号
	Mhz       int      // 主频
	Flags     []string // 标识
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
