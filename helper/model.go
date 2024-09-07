package helper

import (
	"github.com/MuhammadSyaugi13/go-project-1/model/domain"
	"github.com/MuhammadSyaugi13/go-project-1/model/web"
)

func ToCategoryresponse(category domain.Kelas) web.KelasCreateResponse {
	return web.KelasCreateResponse{
		Id:    category.Id,
		Ruang: category.Ruang,
	}
}
