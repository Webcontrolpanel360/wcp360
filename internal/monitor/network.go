package monitor

import (
    "os"
    "strings"
    "strconv"
)

// GetNetworkStats récupère les octets reçus et envoyés sur eth0
func GetNetworkStats() (int64, int64) {
    data, _ := os.ReadFile("/proc/net/dev")
    lines := strings.Split(string(data), "\n")
    for _, line := range lines {
        if strings.Contains(line, "eth0") {
            fields := strings.Fields(line)
            rx, _ := strconv.ParseInt(fields[1], 10, 64)
            tx, _ := strconv.ParseInt(fields[9], 10, 64)
            return rx, tx
        }
    }
    return 0, 0
}
