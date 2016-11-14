package model

var (
	templates []string = []string{
		"machine", "alarm", "collect", "doc",
		// "init", "deploy", "acl", "owner", "route",
	}

	TemplatePrefix string = "_template_"
)

var (
	docTemplate     Resources = Resources{Resource{"describe": "loda"}}
	collectTemplate Resources = Resources{Resource{
		"name":             "app.service.coredump",
		"interval":         "60",
		"measurement_type": "COREDUMP",
		"comment":          "检测 /home/coresave 下生成的 core",
	}, Resource{
		"comment":          "",
		"name":             "cpu.idle",
		"interval":         "10",
		"measurement_type": "CPU",
	}, Resource{
		"comment":          "机器单核采集",
		"name":             "cpu.idle.core",
		"interval":         "10",
		"aggregate":        "avg",
		"measurement_type": "CPU",
	}, Resource{
		"comment":          "最近一分钟服务器负载",
		"name":             "cpu.loadavg.1",
		"interval":         "10",
		"measurement_type": "CPU",
	}, Resource{
		"comment":          "最近十五分钟服务器负载",
		"name":             "cpu.loadavg.15",
		"interval":         "10",
		"measurement_type": "CPU",
	}, Resource{
		"comment":          "最近五分钟服务器负载",
		"name":             "cpu.loadavg.5",
		"interval":         "10",
		"measurement_type": "CPU",
	}, Resource{
		"comment":          "机器磁盘 inode 使用率",
		"name":             "disk.inodes.used.percent",
		"interval":         "120",
		"measurement_type": "DISK",
	}, Resource{
		"comment":          "机器磁盘使用率",
		"name":             "disk.used.percent",
		"interval":         "120",
		"measurement_type": "DISK",
	}, Resource{
		"comment":          "检测文件系统故障. 0 表示文件系统为只读, 1表示文件系统正常",
		"name":             "fs.disk.rw",
		"interval":         "300",
		"measurement_type": "FS",
	}, Resource{
		"comment":          "整个系统被分配的file handles",
		"name":             "fs.files.allocated",
		"interval":         "300",
		"measurement_type": "FS",
	}, Resource{
		"comment":          "整个系统剩余可以分配的 file handles",
		"name":             "fs.files.left",
		"interval":         "300",
		"measurement_type": "FS",
	}, Resource{
		"comment":          "整个系统所有进程能够打开的最多文件数",
		"name":             "fs.files.max",
		"interval":         "300",
		"measurement_type": "FS",
	}, Resource{
		"comment":          "整个系统的file handles 的使用率",
		"name":             "fs.files.used.percent",
		"interval":         "300",
		"measurement_type": "FS",
	}, Resource{
		"comment":          "CPU等待 IO 操作时间",
		"name":             "io.await",
		"interval":         "10",
		"measurement_type": "IO",
	}, Resource{
		"comment":          "io使用率",
		"name":             "io.util",
		"interval":         "10",
		"measurement_type": "IO",
	}, Resource{
		"collect_type":     "FLOW",
		"degree":           "0",
		"file_path":        "/var/log/kernel",
		"func":             "cnt",
		"interval":         "10",
		"measurement_type": "LOG",
		"name":             "kernel.log.OOM",
		"pattern":          "Out of memory",
		"tags":             "service",
		"tag_service":      "kill process \\d+ \\((\\S+)\\)",
		"time_format":      "yyyy-mm-ddTHH:MM:SS",
	}, Resource{
		"comment":          "内核错误日志(I/O error|EXT3-fs error|ERROR on|Medium Error|error recovery|disk error|Illegal block|Out of Memory|dead device|readonly)条数. ",
		"name":             "kernel_log_monitor",
		"interval":         "300",
		"measurement_type": "KERNEL",
	}, Resource{
		"comment":          "服务器心跳",
		"name":             "machine.heartbeat",
		"interval":         "10",
		"measurement_type": "HEALTH",
	}, Resource{
		"comment":          "内存缓存量",
		"name":             "mem.buffers",
		"interval":         "10",
		"measurement_type": "MEM",
	}, Resource{
		"comment":          "内存空闲量",
		"name":             "mem.free",
		"interval":         "10",
		"measurement_type": "MEM",
	}, Resource{
		"comment":          "机器物理内存总量",
		"name":             "mem.total",
		"interval":         "10",
		"measurement_type": "MEM",
	}, Resource{
		"comment":          "机器内存使用率",
		"name":             "mem.used",
		"interval":         "10",
		"measurement_type": "MEM",
	}, Resource{
		"comment":          "内存使用率",
		"name":             "mem.used.percent",
		"interval":         "10",
		"measurement_type": "MEM",
	}, Resource{
		"comment":          "网卡入口流量",
		"name":             "net.in",
		"interval":         "10",
		"measurement_type": "NET",
	}, Resource{
		"comment":          "网络入口丢包数",
		"name":             "net.in.dropped",
		"interval":         "10",
		"measurement_type": "NET",
	}, Resource{
		"comment":          "网卡出口流量",
		"name":             "net.out",
		"interval":         "10",
		"aggregate":        "sum",
		"measurement_type": "NET",
	}, Resource{
		"comment":          "网络出口丢包数",
		"name":             "net.out.dropped",
		"interval":         "10",
		"measurement_type": "NET",
	}, Resource{
		"comment":          "正在使用（正在侦听）的TCP socket 数量",
		"name":             "net.sockets.tcp.inuse",
		"interval":         "10",
		"measurement_type": "NET",
	}, Resource{
		"comment":          "已使用的所有协议socket总量",
		"name":             "net.sockets.used",
		"interval":         "10",
		"measurement_type": "NET",
	}, Resource{
		"comment":          "机器timewait连接数",
		"name":             "net.tcp.timewait",
		"interval":         "15",
		"measurement_type": "NET",
	}, Resource{
		"comment":          "机器和 ntp server 时间差(ms)",
		"name":             "time.offset",
		"interval":         "300",
		"measurement_type": "TIME",
	}}
	RootTemplate map[string]Resources = map[string]Resources{
		TemplatePrefix + "doc":     docTemplate,
		TemplatePrefix + "collect": collectTemplate,
		TemplatePrefix + "alarm":   nil,
		TemplatePrefix + "machine": nil,
	}
)

// TODO: post put