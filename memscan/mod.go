package memscan

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func getProcessModules(procName string) (map[uintptr]string, error) {
	pid, err := FindPidByName(procName)
	if err != nil {
		return nil, err
	}
	return getProcessModulesByPid(pid)
}

func getProcessModulesByPid(pid uint32) (map[uintptr]string, error) {
	// get the handle of the remote process
	handle, err := windows.OpenProcess(
		windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ, false, pid)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(handle)

	list := make([]windows.Handle, 1024)
	var cbNeeded uint32
	// get the list of modules in the process
	err = windows.EnumProcessModulesEx(handle, &list[0], 1024, &cbNeeded, 0)
	if err != nil {
		fmt.Printf("Failed to enumerate modules: %s\n", err)
		return nil, err
	}

	modules := make(map[uintptr]string)

	for _, m := range list {
		var name [256]uint16
		err := windows.GetModuleFileNameEx(handle, m, &name[0], 256)
		if err != nil {
			continue
		}

		exeFile := windows.UTF16ToString(name[:])
		modules[uintptr(m)] = exeFile
	}

	return modules, err
}
