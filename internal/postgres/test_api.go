package postgres

import (
	"context"
	"try-standard-layout/internal/domain"
)

type TestApi interface {
	List(ctx context.Context) ([]*domain.TestApi, error)
}

type testApi struct {
}

func (p *testApi) List(ctx context.Context) ([]*domain.TestApi, error) {
	var data []*domain.TestApi
	err := DBFromCtx(ctx).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
