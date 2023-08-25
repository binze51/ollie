package service

import (
	"context"
	"log"

	"google.golang.org/grpc/health/grpc_health_v1"
)

// Check 重新实现grpc health check
// https://github.com/grpc/grpc/blob/master/doc/health-checking.md
func (h *ServiceImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (resp *grpc_health_v1.HealthCheckResponse, err error) {
	log.Default().Println("serving health")
	resp.Status = grpc_health_v1.HealthCheckResponse_SERVING
	return resp, nil
}
