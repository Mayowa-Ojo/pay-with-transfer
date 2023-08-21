package healthcheck

import (
	"context"
	"os"
	"runtime"
)

type Service struct {
	tag      string
	commit   string
	hostname string
}

func New() HealthService {
	hostname, _ := os.Hostname()

	return &Service{
		hostname: hostname,
	}
}

func (h *Service) Status(ctx context.Context) HealthCheckResponse {
	resp := HealthCheckResponse{
		Status:   true,
		Hostname: h.hostname,
	}

	resp.Version.Tag = h.tag
	resp.Version.Commit = h.commit

	resp.System.Version = runtime.Version()
	resp.System.NumCPU = runtime.NumCPU()
	resp.System.NumGoroutines = runtime.NumGoroutine()

	mem := &runtime.MemStats{}
	runtime.ReadMemStats(mem)
	resp.System.AllocBytes = mem.HeapAlloc
	resp.System.TotalAllocBytes = mem.TotalAlloc
	resp.System.NumHeapObjects = mem.HeapObjects

	return resp
}
