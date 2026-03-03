package monitor

import (
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/mem"
)

type SystemStats struct {
    CPU    float64 `json:"cpu"`
    Memory float64 `json:"memory"`
}

func GetStats() SystemStats {
    c, _ := cpu.Percent(0, false)
    m, _ := mem.VirtualMemory()
    
    cpuUsage := 0.0
    if len(c) > 0 {
        cpuUsage = c[0]
    }

    return SystemStats{
        CPU:    cpuUsage,
        Memory: m.UsedPercent,
    }
}
