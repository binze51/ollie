package service

import (
	"context"
	"flag"
	"fmt"
	"strings"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	authv3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	typev3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

const (
	authorization     = "x-authorization"
	checkHeader       = "x-ext-authz"
	allowedValue      = "allow"
	resultHeader      = "x-ext-authz-check-result"
	receivedHeader    = "x-ext-authz-check-received"
	overrideHeader    = "x-ext-authz-additional-header-override"
	overrideGRPCValue = "grpc-additional-header-override-value"
	resultAllowed     = "allowed"
	resultDenied      = "denied"
)

var (
	serviceAccount = flag.String("allow_service_account", "a",
		"allowed service account, matched against the service account in the source principal from the client certificate")
	denyBody = fmt.Sprintf("denied by ext_authz for not found header `%s: %s` in the request", checkHeader, allowedValue)
)

// Check implements the Envoy grpc check interface.
func (s *ServiceImpl) Check(_ context.Context, request *authv3.CheckRequest) (*authv3.CheckResponse, error) {
	attrs := request.GetAttributes()
	httpAttrs := attrs.GetRequest().GetHttp()
	// Determine whether to allow or deny the request.
	allow := false
	tokenValue, contains := httpAttrs.GetHeaders()[authorization]
	if contains {
		// 使用jwt是合法有效
		jwtInfo, err := s.ParseToken(tokenValue)
		if err != nil {
			allow = false
		}
		// 使用casbin操作权限判断，roles rpc-methodpath 用户id
		// roles /ceres.enterprise.EnterpriseService/GetGroup uid
		// 这里需要使用jwt里的uid查次库(redis或者内存里)，返回一些用户信息：roles info啥的，并判断是否拉黑字段
		// user:=s.GetAccountInfo(jwt.ID) jwt.ID=uid
		// allow = s.batchCheck(user.Roles, httpAttrs.Path, jwt.ID)
		allow = s.batchCheck("jwt.RolesTODO", httpAttrs.Path, jwtInfo.ID)
	} else {
		allow = attrs.Source != nil && strings.HasSuffix(attrs.Source.Principal, "/sa/"+*serviceAccount)
	}

	if allow {
		return s.allow(request), nil
	}

	return s.deny(request), nil
}

func (s *ServiceImpl) allow(request *authv3.CheckRequest) *authv3.CheckResponse {
	// s.logRequest("allowed", request)
	return &authv3.CheckResponse{
		HttpResponse: &authv3.CheckResponse_OkResponse{
			OkResponse: &authv3.OkHttpResponse{
				Headers: []*corev3.HeaderValueOption{
					{
						Header: &corev3.HeaderValue{
							Key:   resultHeader,
							Value: resultAllowed,
						},
					},
					{
						Header: &corev3.HeaderValue{
							Key:   receivedHeader,
							Value: request.GetAttributes().String(),
						},
					},
					{
						Header: &corev3.HeaderValue{
							Key:   overrideHeader,
							Value: overrideGRPCValue,
						},
					},
				},
			},
		},
		Status: &status.Status{Code: int32(codes.OK)},
	}
}

func (s *ServiceImpl) deny(request *authv3.CheckRequest) *authv3.CheckResponse {
	// s.logRequest("denied", request)
	return &authv3.CheckResponse{
		HttpResponse: &authv3.CheckResponse_DeniedResponse{
			DeniedResponse: &authv3.DeniedHttpResponse{
				Status: &typev3.HttpStatus{Code: typev3.StatusCode_Forbidden},
				Body:   denyBody,
				Headers: []*corev3.HeaderValueOption{
					{
						Header: &corev3.HeaderValue{
							Key:   resultHeader,
							Value: resultDenied,
						},
					},
					{
						Header: &corev3.HeaderValue{
							Key:   receivedHeader,
							Value: request.GetAttributes().String(),
						},
					},
					{
						Header: &corev3.HeaderValue{
							Key:   overrideHeader,
							Value: overrideGRPCValue,
						},
					},
				},
			},
		},
		Status: &status.Status{Code: int32(codes.PermissionDenied)},
	}
}
