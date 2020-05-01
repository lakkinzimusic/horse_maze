package logger

import (
	"fmt"
	"runtime"
)

var memory runtime.MemStats

// GetAvailableMemory func
func GetAvailableMemory() (availableMemory uint64) {
	runtime.ReadMemStats(&memory)
	fmt.Printf("Alloc = %v MiB", bToMb(memory.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(memory.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(memory.Sys))
	fmt.Printf("\tNumGC = %v\n", memory.NumGC)
	return memory.Sys
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// func WriteNewRecord(maxLen, thisLen int) {
// 	if maxLen > thisLen {
// 		maxLen = len(newBranch)
// 		fmt.Println("New RecordLen: ", MaxLenBranch)
// 	}
// }
