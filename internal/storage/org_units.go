package storage

import (
	"fmt"
	"github.com/RomanVas30/store/external/dbr_extensions"
	"github.com/RomanVas30/store/internal/entities"
	"github.com/gocraft/dbr"
)

type OrgUnits interface {
	CreateOrgUnit(unit *entities.OrgUnit) error
	GetOrgUnits() (*[]string, error)
	DeleteOrgUnit(name string) error
	UpdateOrgUnit(unit *entities.OrgUnit) error
}

type OrgUnitsStorage struct {
	db *dbr.Connection
}

func NewOrgUnitsStorage(db *dbr.Connection) *OrgUnitsStorage {
	return &OrgUnitsStorage{
		db: db,
	}
}

func (e *OrgUnitsStorage) CreateOrgUnit(unit *entities.OrgUnit) error {
	newSession := e.db.NewSession(nil)

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		err := runner.InsertInto("org_unit").
			Pair("name", unit.Name).
			Pair("description", unit.Description).
			Returning("id").
			Load(&unit.Id)
		if err != nil {
			return err
		}

		if unit.Posts != nil {
			for _, post := range unit.Posts {
				result, err := runner.InsertInto("post_rate").
					Pair("org_unit_id", unit.Id).
					Pair("post", post.PostName).
					Pair("rate", post.Rate).
					Exec()
				if err != nil {
					return err
				}
				if rows, _ := result.RowsAffected(); rows != 1 {
					return fmt.Errorf("failed to add post: %s", post.PostName)
				}
			}
		}

		return nil
	})
	if sessionError != nil {
		return sessionError
	}

	return nil
}

func (e *OrgUnitsStorage) GetOrgUnits() (*[]string, error) {
	newSession := e.db.NewSession(nil)

	var units []string

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		_, err := runner.Select("name").
			From("org_unit").
			Load(&units)
		if err != nil {
			return err
		}
		return nil
	})
	if sessionError != nil {
		return nil, sessionError
	}

	return &units, nil
}

func (e *OrgUnitsStorage) DeleteOrgUnit(name string) error {
	newSession := e.db.NewSession(nil)

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		result, err := runner.DeleteFrom("org_unit").
			Where(dbr.Eq("name", name)).
			Exec()
		if err != nil {
			return err
		}
		if rows, _ := result.RowsAffected(); rows != 1 {
			return fmt.Errorf("failed to delete organisation unit: %s", name)
		}
		return nil
	})
	if sessionError != nil {
		return sessionError
	}

	return nil
}

func (e *OrgUnitsStorage) UpdateOrgUnit(unit *entities.OrgUnit) error {
	newSession := e.db.NewSession(nil)

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		err := runner.Update("org_unit").
			Set("description", unit.Description).
			Where(dbr.Eq("name", unit.Name)).
			Returning("id").
			Load(&unit.Id)
		if err != nil {
			return err
		}

		if unit.AddPosts != nil {
			for _, post := range unit.AddPosts {
				result, err := runner.InsertInto("post_rate").
					Pair("org_unit_id", unit.Id).
					Pair("post", post.PostName).
					Pair("rate", post.Rate).
					Exec()
				if err != nil {
					return err
				}
				if rows, _ := result.RowsAffected(); rows != 1 {
					return fmt.Errorf("failed to add post: %s", post.PostName)
				}
			}
		}

		if unit.DeletePosts != nil {
			for _, post := range unit.DeletePosts {
				result, err := runner.DeleteFrom("post_rate").
					Where(dbr.Eq("id", runner.Select("id").
						From("post_rate").
						Where(
							dbr.And(
								dbr.Eq("org_unit_id", unit.Id),
								dbr.Eq("post", post.PostName),
								dbr.Eq("rate", post.Rate),
							),
						).
						Limit(1),
					)).
					Exec()
				if err != nil {
					return err
				}
				if rows, _ := result.RowsAffected(); rows != 1 {
					return fmt.Errorf("failed to delete post: %s", post.PostName)
				}
			}
		}

		return nil
	})
	if sessionError != nil {
		return sessionError
	}

	return nil
}
