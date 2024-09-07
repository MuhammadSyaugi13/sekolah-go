package service

import (
	"context"
	"database/sql"

	"github.com/MuhammadSyaugi13/go-project-1/helper"
	"github.com/MuhammadSyaugi13/go-project-1/model/domain"
	"github.com/MuhammadSyaugi13/go-project-1/model/web"
	"github.com/MuhammadSyaugi13/go-project-1/repository"
	"github.com/go-playground/validator/v10"
)

type KelasServiceImpl struct {
	KelasRepository repository.KelasRepository
	DB              *sql.DB
	Validate        validator.Validate
}

func (service *KelasServiceImpl) Create(ctx context.Context, request web.KelasCreateRequest) web.KelasCreateResponse {

	// validasi request inputan
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// membuat koneksi transaksi db
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	kelas := domain.Kelas{
		Ruang: request.Ruang,
	}

	kelas = service.KelasRepository.Create(ctx, tx, kelas)

	return helper.ToCategoryresponse(kelas)

}
