package storage

import (
	"fmt"
	"github.com/RomanVas30/store/external/dbr_extensions"
	"github.com/RomanVas30/store/internal/entities"
	"github.com/gocraft/dbr"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GetUser(username, password string) (entities.User, error)
	ChangePassword(changePass entities.ChangePassword) error
}

type Auth struct {
	db *dbr.Connection
}

func NewAuth(db *dbr.Connection) *Auth {
	return &Auth{
		db: db,
	}
}

func (r *Auth) CreateUser(user entities.User) (int, error) {
	newSession := r.db.NewSession(nil)

	var id int
	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		err := runner.InsertInto("users").
			Pair("name", user.Name).
			Pair("username", user.Username).
			Pair("password_hash", user.Password).
			Returning("id").
			Load(&id)
		if err != nil {
			return err
		}
		return nil
	})
	if sessionError != nil {
		return 0, sessionError
	}

	return id, nil
}

func (r *Auth) GetUser(username, password string) (entities.User, error) {
	newSession := r.db.NewSession(nil)

	var user entities.User
	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		err := runner.Select("id", "role").
			From("users").
			Where(dbr.Eq("username", username)).
			Where(dbr.Eq("password_hash", password)).
			LoadOne(&user)
		if err == dbr.ErrNotFound {
			return fmt.Errorf("user with this credentials was not found")
		}
		if err != nil {
			return err
		}
		return nil
	})
	if sessionError != nil {
		return entities.User{}, sessionError
	}

	return user, nil
}

func (r *Auth) ChangePassword(changePass entities.ChangePassword) error {
	newSession := r.db.NewSession(nil)

	sessionError := dbr_extensions.CreateTx(newSession, func(runner dbr.SessionRunner) error {
		result, err := runner.Update("users").
			Set("password_hash", changePass.NewPassword).
			Where(dbr.Eq("username", changePass.Username)).
			Where(dbr.Eq("password_hash", changePass.OldPassword)).
			Exec()
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return fmt.Errorf("user with this credentials was not found")
		}

		return nil
	})
	if sessionError != nil {
		return sessionError
	}

	return nil
}
