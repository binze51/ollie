package service

import (
	"context"

	releasesvc "ollie/kitex_gen/release"
)

func (h *ServiceImpl) FetchServiceNamesByPrefix(ctx context.Context, req *releasesvc.FetchServiceNamesByPrefixRequest) (resp *releasesvc.FetchServiceNamesByPrefixResponse, err error) {
	return nil, nil
}

func (h *ServiceImpl) FetchServiceVersionsByPrefix(ctx context.Context, req *releasesvc.FetchServiceVersionsByPrefixRequest) (resp *releasesvc.FetchServiceVersionsByPrefixResponse, err error) {
	return nil, nil
}

func (h *ServiceImpl) CreateRelease(ctx context.Context, req *releasesvc.CreateReleaseRequest) (resp *releasesvc.CreateReleaseResponse, err error) {
	return nil, nil
}

func (h *ServiceImpl) FetchReleases(ctx context.Context, req *releasesvc.FetchReleasesRequest) (resp *releasesvc.FetchReleasesResponse, err error) {
	return nil, nil
}

func (h *ServiceImpl) UpdateReleaseByID(context.Context, *releasesvc.UpdateReleaseByIDRequest) (*releasesvc.UpdateReleaseByIDResponse, error) {
	return nil, nil
}

func (h *ServiceImpl) RemoveReleaseByID(context.Context, *releasesvc.RemoveReleaseByIDRequest) (*releasesvc.RemoveReleaseByIDResponse, error) {
	return nil, nil
}

func (h *ServiceImpl) RetryReleaseByID(context.Context, *releasesvc.RetryReleaseByIDRequest) (*releasesvc.RetryReleaseByIDResponse, error) {
	return nil, nil
}

func (h *ServiceImpl) PrimaryReleaseByID(context.Context, *releasesvc.PrimaryReleaseByIDRequest) (*releasesvc.PrimaryReleaseByIDResponse, error) {
	return nil, nil
}

func (h *ServiceImpl) RollbackPrimaryReleaseByHistoryID(context.Context, *releasesvc.RollbackPrimaryReleaseByHistoryIDRequest) (*releasesvc.RollbackPrimaryReleaseByHistoryIDResponse, error) {
	return nil, nil
}
