package helper

import "database/sql"

func DBTransaction(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		HelperError(errRollback, "db transaction rollback failed")
		panic(err)
	} else {
		errCommit := tx.Commit()
		HelperError(errCommit, "db transaction commit failed")
	}
}
