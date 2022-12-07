package memscan_test

import (
	"testing"

	"github.com/joshfinley/memscan/memscan"
)

func TestGetAllocs(t *testing.T) {
	pid, _ := memscan.FindPidByName("explorer.exe")
	allocs, err := memscan.GetAllocs(pid)
	if err != nil {
		t.Fail()
	}

	t.Log(allocs)
}
