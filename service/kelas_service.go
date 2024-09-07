package service

import (
	"context"

	"github.com/MuhammadSyaugi13/go-project-1/model/web"
)

type KelasService interface {
	Create(ctx context.Context, request web.KelasCreateRequest) web.KelasCreateResponse
}
