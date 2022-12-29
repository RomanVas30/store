package storage

import (
	"fmt"
	"github.com/RomanVas30/store/external/dbr_extensions"
	"github.com/RomanVas30/store/internal/entities"
	"github.com/gocraft/dbr"
	"time"
)

type Staff interface {
	CreateStaffer(staffer *entities.Staffer) error
	GetStaff() (*[]entities.ShortStaffer, error)
	DeleteStaffer(snils string) error
	SearchStaff(fio string, snils string) (*[]entities.Staffer, error)
	UpdateStaffer(staffer *entities.Staffer, addPosts []entities.StafferPost, deletePosts []entities.StafferPost) error
}

type StaffStorage struct {
	db *dbr.Connection
}

func NewStaffStorage(db *dbr.Connection) *StaffStorage {
	return &StaffStorage{
		db: db,
	}
}

func addStrafferPosts(runner dbr.SessionRunner, stafferId int, addPosts []entities.StafferPost) error {
	for _, post := range addPosts {
		result, err := runner.Update("post_rate").
			Set("staffer_id", stafferId).
			Where(dbr.Eq("id", runner.Select("id").
				From("post_rate").
				Where(
					dbr.And(
						dbr.Eq("staffer_id", nil),
						dbr.Eq("org_unit_id",
							runner.Select("id AS org_unit_id").
								From("org_unit").
								Where(dbr.Eq("name", post.OrgUnitName)),
						),
						dbr.Eq("post", post.PostName),
						dbr.Eq("rate", post.Rate),
					),
				).
				Limit(1),
			)).
			Exec()
		if err != nil {
			return fmt.Errorf("failed to add post: %v", err)
		}
		if rows, _ := result.RowsAffected(); rows != 1 {
			return fmt.Errorf("failed to add post: %s", post.PostName)
		}
	}

	return nil
}

func (e *StaffStorage) CreateStaffer(staffer *entities.Staffer) error {
	newSession := e.db.NewSession(nil)

	employmentDate := time.Now()

	staffer.EmploymentDate = employmentDate.UTC().Format("02.01.2006 15:04:05")

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		err := runner.InsertInto("staff").
			Pair("fio", staffer.FIO).
			Pair("birth_year", staffer.BirthYear).
			Pair("snils", staffer.SNILS).
			Pair("employment_date", employmentDate).
			Returning("id").
			Load(&staffer.Id)
		if err != nil {
			return err
		}

		if staffer.Posts != nil {
			if err := addStrafferPosts(runner, staffer.Id, staffer.Posts); err != nil {
				return err
			}
		}

		return nil
	})
	if sessionError != nil {
		return sessionError
	}

	return nil
}

func (e *StaffStorage) GetStaff() (*[]entities.ShortStaffer, error) {
	newSession := e.db.NewSession(nil)

	var staff []entities.ShortStaffer

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		_, err := runner.Select("fio", "birth_year", "snils").
			From("staff").
			Load(&staff)
		if err != nil {
			return err
		}
		return nil
	})
	if sessionError != nil {
		return nil, sessionError
	}

	return &staff, nil
}

func (e *StaffStorage) DeleteStaffer(snils string) error {
	newSession := e.db.NewSession(nil)

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		result, err := runner.DeleteFrom("staff").
			Where(dbr.Eq("snils", snils)).
			Exec()
		if err != nil {
			return err
		}
		if rows, _ := result.RowsAffected(); rows != 1 {
			return fmt.Errorf("failed to delete staffer with SNILS: %s", snils)
		}
		return nil
	})
	if sessionError != nil {
		return sessionError
	}

	return nil
}

func (e *StaffStorage) SearchStaff(fio string, snils string) (*[]entities.Staffer, error) {
	newSession := e.db.NewSession(nil)

	var staff []entities.Staffer
	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		if fio != "" {
			_, err := runner.Select("id", "fio", "birth_year", "employment_date", "snils").
				From("staff").
				Where(dbr.Eq("fio", fio)).
				Load(&staff)
			if err != nil {
				return err
			}
		} else if snils != "" {
			_, err := runner.Select("id", "fio", "birth_year", "employment_date", "snils").
				From("staff").
				Where(dbr.Eq("snils", snils)).
				Load(&staff)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("incorrect data for search")
		}

		for i := 0; i < len(staff); i++ {
			_, err := runner.Select("ou.name", "pr.post", "pr.rate").
				From(dbr.I("post_rate").As("pr")).
				Join(dbr.I("org_unit").As("ou"), dbr.Expr("pr.org_unit_id = ou.id")).
				Where(dbr.Eq("pr.staffer_id", staff[i].Id)).
				Load(&staff[i].Posts)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if sessionError != nil {
		return nil, sessionError
	}

	return &staff, nil
}

func (e *StaffStorage) UpdateStaffer(
	staffer *entities.Staffer,
	addPosts []entities.StafferPost,
	deletePosts []entities.StafferPost,
) error {
	newSession := e.db.NewSession(nil)

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		updateStaffer := runner.Update("staff")
		if staffer.FIO != "" {
			updateStaffer.Set("fio", staffer.FIO)
		}
		err := updateStaffer.Where(dbr.Eq("snils", staffer.SNILS)).
			Returning("id").
			Load(&staffer.Id)
		if err != nil {
			return err
		}
		if staffer.Id == 0 {
			return fmt.Errorf("staffer not found")
		}

		if addPosts != nil {
			if err := addStrafferPosts(runner, staffer.Id, addPosts); err != nil {
				return err
			}
		}

		if deletePosts != nil {
			for _, post := range deletePosts {
				result, err := runner.Update("post_rate").
					Set("staffer_id", nil).
					Where(dbr.And(
						dbr.Neq("staffer_id", nil),
						dbr.Eq("org_unit_id", runner.Select("id").
							From("org_unit").
							Where(dbr.Eq("name", post.OrgUnitName))),
						dbr.Eq("post", post.PostName),
						dbr.Eq("rate", post.Rate),
					)).Exec()
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
