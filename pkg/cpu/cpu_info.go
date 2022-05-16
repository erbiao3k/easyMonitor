package cpu

import (
	"easyMonitor/utils/number"
	"github.com/shirou/gopsutil/v3/cpu"
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
