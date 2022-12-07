package memscan

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type Alloc struct {
	Addr       uintptr
	Size       uint32
	AllocType  string
	Protection string
	ModuleName string
}

func GetAllocs(pid uint32) (*map[uintptr]*Alloc, error) {

	mbimap, err := getMemInfoMap(pid)
	if mbimap == nil {
		return nil, err
	}

	// get allocations
	allocs := make(map[uintptr]*Alloc)
	for p, a := range mbimap {
		var ma Alloc
		ma.Addr = p

		switch a.Type {
		case MEM_IMAGE:
			ma.AllocType = "Image"
		case MEM_MAPPED:
			ma.AllocType = "Mapped"
		case MEM_PRIVATE:
			ma.AllocType = "Private"
		}

		ma.Protection = protectDwordToString(a.Protect)
		ma.Size = uint32(a.RegionSize)

		allocs[p] = &ma
	}

	// get the module names
	modules, err := getProcessModulesByPid(pid)
	if err != nil {
		return nil, err
	}

	for ptr, _ := range allocs {
		if val, ok := modules[ptr]; ok {
			a := allocs[ptr]
			a.ModuleName = val
			allocs[ptr] = a
		}
	}

	return &allocs, nil
}

func protectDwordToString(dword uint32) string {
	switch dword {
	case windows.PAGE_READONLY:
		return "PAGE_READONLY"
	case windows.PAGE_EXECUTE:
		return "PAGE_EXECUTE"
	case windows.PAGE_READWRITE:
		return "PAGE_READWRITE"
	case windows.PAGE_EXECUTE_READ:
		return "PAGE_EXECUTE_READ"
	case windows.PAGE_EXECUTE_READWRITE:
		return "PAGE_EXECUTE_READWRITE"
	case windows.PAGE_NOACCESS:
		return "PAGE_NOACCESS"
	case windows.PAGE_WRITECOMBINE:
		return "PAGE_WRITECOMBINE"
	case windows.PAGE_GUARD:
		return "PAGE_GUARD"
	case windows.PAGE_TARGETS_NO_UPDATE:
		return "PAGE_TARGETS_NO_UPDATE"
	}

	return ""
}

func getMemInfoMap(pid uint32) (map[uintptr]*windows.MemoryBasicInformation, error) {
	h, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION, false, pid)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(h)

	allocs := make(map[uintptr]*windows.MemoryBasicInformation)
	for addr := uintptr(0); ; {
		var mbi windows.MemoryBasicInformation
		err := windows.VirtualQueryEx(h, addr, &mbi, unsafe.Sizeof(mbi))
		if err != nil {
			break
		}
		allocs[mbi.BaseAddress] = &mbi
		addr += mbi.RegionSize
	}

	return allocs, nil
}
