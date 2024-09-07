package repository

import (
	"context"
	"database/sql"

	"github.com/MuhammadSyaugi13/go-project-1/helper"
	"github.com/MuhammadSyaugi13/go-project-1/model/domain"
)

type KelasRepositoryImpl struct{}

func (repository *KelasRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, kelas domain.Kelas) domain.Kelas {
	SQL := "insert into kelas(ruang) values(?)"
	result, err := tx.ExecContext(ctx, SQL, kelas.Ruang) // insert ke database
	helper.PanicIfError(err)

	//ambil id data yang diinsertkan
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	kelas.Id = int(id)
	return kelas
}
