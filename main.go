package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
)

func main() {
	info, _ := host.Info()
	printIndent(info)
	printIndentE(load.Avg())
	partitions, _ := disk.Partitions(true)
	for _, partition := range partitions {
		printIndentE(disk.Usage(partition.Mountpoint))
	}
}

func printIndent(v any) {
	printIndentE(v, nil)
}

func printIndentE(v any, e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
	b, _ := json.MarshalIndent(v, "", "\t")
	fmt.Printf("%s\n", b)
}
