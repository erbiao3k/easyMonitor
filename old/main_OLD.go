package old

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
)

var (
	fstype = []string{"xfs", "ext4", "ext3"}
)

// CPU
//
//	整体使用率
//	单个使用率
//  使用率异变情况

// 内存使用率
// swap

// 网络
// 	连接数
//	流量上行速率
//	流量下行速率
// 	丢包
// 	错包

// 系统
//	在线时间
//  /etc/passwd 文件是否被修改

// 进程
//	指定端口打开否
// 进程总数（进程总数变化）

// 基础服务监控

// diskItem 磁盘监控项目
type diskItem struct {
	mountPoint        string
	sysType           string
	spaceFree         string
	inodesUsedPercent string
	writable          string
	iops              string
	readRate          string
	writeRate         string
}

func memInfo() (percent string) {
	info, _ := mem.VirtualMemory()
	percent = fmt.Sprintf("%f", info.UsedPercent)
	return

}
func diskInfo() {
	info, _ := disk.Partitions(true)

	for _, stat := range info {
		for _, t := range fstype {
			if stat.Fstype != t {
				continue
			}
			mountPoint := stat.Mountpoint
			deviceName := stat.Device
			diskUsage, _ := disk.Usage(deviceName)
			diskIOCounters, _ := disk.IOCounters(deviceName)
			log.Println(mountPoint, "磁盘空间使用率：", diskUsage.UsedPercent)
			log.Println(mountPoint, "磁盘inode使用率：", diskUsage.InodesUsedPercent)
			log.Println(mountPoint, "----------", diskIOCounters)
		}
	}
}
