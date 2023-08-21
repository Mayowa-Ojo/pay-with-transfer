package healthcheck

import (
	"context"
)

type HealthService interface {
	Status(ctx context.Context) HealthCheckResponse
}

type HealthCheckResponse struct {
	Version struct {
		Tag    string `json:"tag"`
		Commit string `json:"commit"`
	} `json:"version"`
	Status   bool   `json:"status"`
	Hostname string `json:"hostname"`
	System   struct {
		Version         string `json:"version"`
		NumCPU          int    `json:"num_cpu"`
		NumGoroutines   int    `json:"num_goroutines"`
		NumHeapObjects  uint64 `json:"num_heap_objects"`
		TotalAllocBytes uint64 `json:"total_alloc_bytes"`
		AllocBytes      uint64 `json:"alloc_bytes"`
	} `json:"system,omitempty"`
}
