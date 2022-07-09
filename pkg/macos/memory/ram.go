package memory

import (
	"fmt"

	"github.com/pbnjay/memory"
)

func getRam() float64 {
	return float64(memory.TotalMemory())
}

func getRamString() string {
	mem := getRam()
	return fmt.Sprintf("Total system memory: %.2fGB", mem/1073741824)
}
