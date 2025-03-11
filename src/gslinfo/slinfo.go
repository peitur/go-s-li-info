package main

import (
	"core"
	"flag"
	"fmt"
)

func main() {

	var (
		withAllInfo       bool
		withHostnameInfo  bool
		withLoadInfo      bool
		withUptimeInfo    bool
		withModulesInfo   bool
		withCpuInfo       bool
		withNicInfo       bool
		withMemoryInfo    bool
		withDiskInfo      bool
		withUsbInfo       bool
		withPartInfo      bool
		withConectionInfo bool
		withUserInfo      bool
		withProcessesInfo bool
		withProcessInfo   uint64
	)

	flag.BoolVar(&withAllInfo, "all", false, "All Info")

	flag.BoolVar(&withLoadInfo, "load", false, "System Load Average Info")
	flag.BoolVar(&withHostnameInfo, "hostname", false, "System Uptime Info")
	flag.BoolVar(&withUptimeInfo, "uptime", false, "System Uptime Info")
	flag.BoolVar(&withModulesInfo, "kmodules", false, "Loaded Kernel Module Info")
	flag.BoolVar(&withCpuInfo, "cpu", false, "CPU Info")
	flag.BoolVar(&withNicInfo, "nic", false, "NIC Info")
	flag.BoolVar(&withDiskInfo, "disk", false, "Disk Info")
	flag.BoolVar(&withMemoryInfo, "mem", false, "Memory Info")
	flag.BoolVar(&withUsbInfo, "usb", false, "Connected USB Info")
	flag.BoolVar(&withPartInfo, "partition", false, "Partition Info")
	flag.BoolVar(&withConectionInfo, "connection", false, "Connection Info")
	flag.BoolVar(&withUserInfo, "user", false, "User Info")
	flag.BoolVar(&withProcessesInfo, "processes", false, "Processes Info")
	flag.Uint64Var(&withProcessInfo, "process", 0, "Processe Info")

	flag.Parse()

	fmt.Println("CPUInfo ", withCpuInfo)
	fmt.Println(core.GOSLI_CONST_PROC)
}
