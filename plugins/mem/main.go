package mem

import (
  "github.com/shirou/gopsutil/mem"
)

func Get() *mem.VirtualMemoryStat {
  memstats, _ := mem.VirtualMemory()
  return memstats
}
