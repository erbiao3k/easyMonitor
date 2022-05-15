package mem

import (
	"github.com/shirou/gopsutil/v3/mem"
	"log"
)

// MemItem 内存监控项
type MemItem struct {
	Total       uint64  // 总内存，bytes
	Available   uint64  // 可用内存
	Used        uint64  // 已用内存
	UsedPercent float64 // 已用内存百分比
	Free        uint64  // 空闲内存
	Active      uint64  // 活跃内存
	Inactive    uint64  // 非活跃内存
	SwapTotal   uint64  // swap总大小
	SwapFree    uint64  // 空闲swap大小
}

func Info() (mi MemItem) {
	m, _ := mem.VirtualMemory()

	mi.Total = m.Total
	mi.Available = m.Available
	mi.Used = m.Used
	mi.UsedPercent = m.UsedPercent
	mi.Free = m.Free
	mi.Active = m.Active
	mi.Inactive = m.Inactive
	mi.SwapTotal = m.SwapTotal
	mi.SwapFree = m.SwapFree

	return
}

func init() {
	log.Printf(`【内存信息】 内存总量: "%d", 内存使用率: "%.2f%%", swap总量: "%d", swap空闲大小: "%d"`, Info().Total, Info().UsedPercent, Info().SwapTotal, Info().SwapFree)
}
