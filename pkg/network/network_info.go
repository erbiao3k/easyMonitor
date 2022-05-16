package network

import (
	"github.com/shirou/gopsutil/v3/net"
	"log"
)

// 网络

// 连接数	总数

// 连接数	LISTENING状态	正在侦听的端口数

// 连接数	ESTABLISHED状态	正在通信的连接数

// 连接数	ESTABLISHED状态	正在通信的连接数
//			1、周期内变化比较大，考虑应用或场景的问题

// 连接数	CLOSE_WAIT状态	对方主动关闭连接或者网络异常导致连接中断，这时我方的状态会变成CLOSE_WAIT 此时我方要调用close()来使得连接正确关闭
//			1、比较多时需要调优或判断服务器是否被攻击

// 连接数	TIME_WAIT状态	我方主动调用close()断开连接，收到对方确认后状态变为TIME_WAIT
//			1、比较多时需要调优或判断服务器是否被攻击

// 连接数	SYN_SENT状态	表示请求连接，当你要访问其它的计算机的服务时首先要发个同步信号给该端口，此时状态为SYN_SENT，如果连接成功了就变为 ESTABLISHED，此时SYN_SENT状态非常短暂
//			1、	SYN_SENT非常多且在向不同的机器发出，那你的机器可能中了冲击波或震荡波之类的病毒

//	每网卡流量上行速率

//	每网卡流量下行速率

// 	每网卡丢包

// 	每网卡错包

func Info() {
	i1, _ := net.Connections("inet")
	for _, stat := range i1 {
		log.Println(stat.Status)
	}

	i2, _ := net.IOCounters(true)
	for _, stat := range i2 {
		log.Println(stat)
	}
	log.Println(len(i1))
}
