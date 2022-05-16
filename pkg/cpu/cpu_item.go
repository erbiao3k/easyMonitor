package cpu

import (
	"easyMonitor/utils/number"
	"github.com/shirou/gopsutil/v3/cpu"
	"time"
)

// CPU
//	整体使用率
//	单个使用率
//  使用率异变情况

type CpuItem struct {
	UsedPercentTotal float64 // 所有核心总的使用率
}

func Item() (ci CpuItem) {
	cpuPercentList, _ := cpu.Percent(time.Second, false)
	ci.UsedPercentTotal = number.Decimal(cpuPercentList[0])
	return
}
