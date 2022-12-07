package main

import (
	"flag"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/joshfinley/memscan/memscan"
)

func main() {
	name := flag.String("n", "", "Name of the process")
	pid := flag.Int("p", 0, "ID of the process")
	flag.Parse()

	if *name == "" && *pid == 0 {
		fmt.Println("Please specify either a process name or process id")
		flag.PrintDefaults()
		return
	} else if *name != "" && *pid != 0 {
		fmt.Println("Please specify either a process name or process id, not both")
		flag.PrintDefaults()
		return
	}

	if *name != "" {
		pid, err := memscan.FindPidByName(*name)
		if err != nil {
			return
		}

		printMemMap(pid)

	}
	if *pid != 0 {
		printMemMap(uint32(*pid))
	}

}

func printMemMap(pid uint32) error {
	allocs, err := memscan.GetAllocs(pid)
	if err != nil {
		return err
	}

	colAddr := "Address"
	colSize := "Size"
	colType := "Type"
	colProt := "Protect"
	colMod := "Module Path"

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{colAddr, colSize, colType, colProt, colMod})
	tw.SetTitle(fmt.Sprintf("Memory Map for Process %d", pid))

	for _, alloc := range *allocs {
		row := table.Row{
			fmt.Sprintf("0x%02x", uint64(alloc.Addr)),
			fmt.Sprintf("0x%x", alloc.Size),
			alloc.AllocType,
			alloc.Protection,
			alloc.ModuleName}
		tw.AppendRows([]table.Row{row})

	}

	fmt.Println(tw.Render())
	return nil
}
