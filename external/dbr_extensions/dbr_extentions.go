package dbr_extensions

import "github.com/gocraft/dbr"

func CreateTx(session *dbr.Session, sessFunc func(runner dbr.SessionRunner) error) error {
	tx, err := session.Begin()
	if err != nil {
		return err
	}

	defer tx.RollbackUnlessCommitted()

	if err := sessFunc(tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
