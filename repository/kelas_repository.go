package repository

import (
	"context"
	"database/sql"

	"github.com/MuhammadSyaugi13/go-project-1/model/domain"
)

type KelasRepository interface {
	Create(ctx context.Context, tx *sql.Tx, kelas domain.Kelas) domain.Kelas
}
