package memlimit

import (
	"github.com/yonomesh/automemlimit/memory"
)

// FromSystem returns the total memory of the system.
func FromSystem() (uint64, error) {
	limit := memory.TotalMemory()
	if limit == 0 {
		return 0, ErrNoLimit
	}
	return limit, nil
}
