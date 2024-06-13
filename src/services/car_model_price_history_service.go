package services

import (
	"context"

	"github.com/nazarow/golang-clean/api/dto"
	"github.com/nazarow/golang-clean/config"
	"github.com/nazarow/golang-clean/data/db"
	"github.com/nazarow/golang-clean/data/models"
	"github.com/nazarow/golang-clean/pkg/logging"
)

type CarModelPriceHistoryService struct {
	base *BaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistoryRequest, dto.UpdateCarModelPriceHistoryRequest, dto.CarModelPriceHistoryResponse]
}

func NewCarModelPriceHistoryService(cfg *config.Config) *CarModelPriceHistoryService {
	return &CarModelPriceHistoryService{
		base: &BaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistoryRequest, dto.UpdateCarModelPriceHistoryRequest, dto.CarModelPriceHistoryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *CarModelPriceHistoryService) Create(ctx context.Context, req *dto.CreateCarModelPriceHistoryRequest) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelPriceHistoryService) Update(ctx context.Context, id int, req *dto.UpdateCarModelPriceHistoryRequest) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelPriceHistoryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelPriceHistoryService) GetById(ctx context.Context, id int) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelPriceHistoryService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelPriceHistoryResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
