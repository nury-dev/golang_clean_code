package services

import (
	"context"

	"github.com/nazarow/golang-clean/api/dto"
	"github.com/nazarow/golang-clean/config"
	"github.com/nazarow/golang-clean/data/db"
	"github.com/nazarow/golang-clean/data/models"
	"github.com/nazarow/golang-clean/pkg/logging"
)

type CompanyService struct {
	base *BaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]
}

func NewCompanyService(cfg *config.Config) *CompanyService {
	return &CompanyService{
		base: &BaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Country"},
			},
		},
	}
}

// Create
func (s *CompanyService) Create(ctx context.Context, req *dto.CreateCompanyRequest) (*dto.CompanyResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CompanyService) Update(ctx context.Context, id int, req *dto.UpdateCompanyRequest) (*dto.CompanyResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CompanyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CompanyService) GetById(ctx context.Context, id int) (*dto.CompanyResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CompanyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CompanyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
