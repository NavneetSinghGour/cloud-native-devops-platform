package models

type RuntimeInfo struct {
	GoVersion  string `json:"goVersion"`
	OS          string `json:"os"`
	Arch        string `json:"arch"`
	CPU         int    `json:"cpu"`
	Goroutines  int    `json:"goroutines"`
	MemoryAlloc uint64 `json:"memoryAlloc"`
}
