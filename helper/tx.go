package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback) // jika error pada rollback
		panic(err)
	} else {

		//commit transaksi
		errorCommit := tx.Commit() //commit Transaksi jika tidak ada error
		PanicIfError(errorCommit)  // jika terjadi error saat commit
	}
}
