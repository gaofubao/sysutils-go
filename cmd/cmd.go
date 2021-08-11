package main

import (
	"fmt"
	"sysutils-go/internal/process"
)

func main() {
	procs, err := process.Processes()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, proc := range procs {
		name, err := proc.Name()
		if err != nil {
			fmt.Println(err)
			return
		}

		if name == "java" {
			fmt.Println("pid: ", proc.Pid)

			cpuPercent, _ := proc.CPUPercent()
			fmt.Println("cpu percent: ", cpuPercent)

			memPercent, _ := proc.MemoryPercent()
			fmt.Println("memory percent: ", memPercent)

			netStat, _ := proc.NetIOCounters(true)
			for _, stat := range netStat {
				bytesSent := stat.BytesSent
				bytesRecv := stat.BytesRecv
				fmt.Println("net: ", bytesSent, bytesRecv)
			}
			//proc.Kill()
		}
	}

}
