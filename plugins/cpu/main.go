package cpu

import (
  "time"
  "github.com/shirou/gopsutil/cpu"
)

func Get() []cpu.TimesStat {
  cpustats, _ := cpu.Times(false)
  time.Sleep(1*time.Second)
  cpustats, _ = cpu.Times(false)
  return cpustats
}
