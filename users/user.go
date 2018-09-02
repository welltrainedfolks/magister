// Copyright (c) 2018, Well Trained Folks and contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package users

import (
	// stdlib
	"crypto/sha256"
	"fmt"
	"time"

	// local
	"github.com/welltrainedfolks/magister/internal/database"
	"github.com/welltrainedfolks/magister/internal/helpers"

	// other
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/scrypt"
)

// User represents single user in system.
type User struct {
	ID           int       `db:"id"`
	Login        string    `db:"login"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	PasswordSalt string    `db:"password_salt"`
	IsActive     bool      `db:"is_active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// GetCurrentlyLoggedInUser returns user data based on echo's Context
// information.
func GetCurrentlyLoggedInUser(ec echo.Context) *User {
	if !ec.Get("AUTHORIZED").(bool) {
		return nil
	}
	user := &User{}
	err := database.DB.Get(user, database.DB.Rebind("SELECT * FROM `users` WHERE id=?"), ec.Get("UID").(int))
	if err != nil {
		log.Error().Msgf("Failed to get user with ID '%s': %s", ec.Get("UID").(int), err.Error())
		return nil
	}

	log.Debug().Msgf("Got user: %+v", user)
	return user
}

// GetCurrentlyLoggedInUserName returns user name for currently logged
// in user.
func GetCurrentlyLoggedInUserName(ec echo.Context) string {
	u := GetCurrentlyLoggedInUser(ec)
	if u == nil {
		return ""
	}
	return u.Email
}

// GetUser returns user data from database.
func GetUser(email string) *User {
	user := &User{}
	err := database.DB.Get(user, database.DB.Rebind("SELECT * FROM `users` WHERE email=?"), email)
	if err != nil {
		log.Error().Msgf("Failed to get user with email '%s': %s", email, err.Error())
		return nil
	}

	log.Debug().Msgf("Got user: %+v", user)
	return user
}

// GetUserByID returns user by ID.
func GetUserByID(id int) *User {
	user := &User{}
	err := database.DB.Get(user, database.DB.Rebind("SELECT * FROM `users` WHERE id=?"), id)
	if err != nil {
		log.Error().Msgf("Failed to get user with id '%d': %s", id, err.Error())
		return nil
	}

	log.Debug().Msgf("Got user: %+v", user)
	return user
}

// GetUserByLogin returns user by login.
func GetUserByLogin(login string) *User {
	user := &User{}
	err := database.DB.Get(user, database.DB.Rebind("SELECT * FROM `users` WHERE login=?"), login)
	if err != nil {
		log.Error().Msgf("Failed to get user with login '%s': %s", login, err.Error())
		return nil
	}

	log.Debug().Msgf("Got user: %+v", user)
	return user
}

// NewUser creates user in database.
func NewUser(login, email, password string) *User {
	u := &User{}
	u.Login = login
	u.Email = email
	u.CreatePassword(password)
	u.IsActive = false
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()

	// Create user in database.
	res, err := database.DB.NamedExec("INSERT INTO `users` (login, email, password, password_salt, is_active, created_at, updated_at) VALUES (:login, :email, :password, :password_salt, :is_active, :created_at, :updated_at)", u)
	if err != nil {
		log.Error().Msgf("Failed to create new user: %s", err.Error())
		return nil
	}

	lastInsertedID, err1 := res.LastInsertId()
	if err1 != nil {
		log.Error().Msgf("Failed to get last inserted ID for user insertion: %s", err1.Error())
		return nil
	}

	u.ID = int(lastInsertedID)
	return u
}

// CheckPassword checks password validity.
func (u *User) CheckPassword(password string) bool {
	if u.Password == "" || u.PasswordSalt == "" {
		return false
	}

	pass := u.hashPassword(password, u.PasswordSalt)

	if pass == u.Password {
		return true
	}

	return false
}

// CreatePassword creates password and hash.
func (u *User) CreatePassword(password string) {
	log.Debug().Msg("Generating password hash...")

	// Generate random string.
	randomString := helpers.GenerateRandomString(64)
	// Generate scrypt'd password salt.
	passwordSaltBytes, _ := scrypt.Key([]byte(randomString), []byte(randomString), 32768, 8, 1, 32)
	// Sha256 it!
	passwordSaltSha256Bytes := sha256.Sum256(passwordSaltBytes)
	// Stringify it.
	u.PasswordSalt = fmt.Sprintf("%x", passwordSaltSha256Bytes[:])

	u.Password = u.hashPassword(password, u.PasswordSalt)

}

// Delete deletes current user from database.
func (u *User) Delete() error {
	_, err := database.DB.NamedExec("DELETE FROM users WHERE login=:login", u)
	return err
}

// Hashes provided password and returns it's hashed and crypted value.
func (u *User) hashPassword(password string, salt string) string {
	// Generate scrypt'd password itself.
	passwordBytes, _ := scrypt.Key([]byte(password), []byte(u.PasswordSalt), 32768, 8, 1, 32)
	// Sha256 it and stringify.
	passwordBytesSha256 := sha256.Sum256(passwordBytes)
	passwordHashString := fmt.Sprintf("%x", passwordBytesSha256[:])
	return passwordHashString
}

// Save saves user.
func (u *User) Save() {
	u.UpdatedAt = time.Now().UTC()
	_, err := database.DB.NamedExec("UPDATE `users` SET password=:password, password_salt=:password_salt, is_active=:is_active, updated_at=:updated_at WHERE id=:id", u)
	if err != nil {
		log.Error().Msgf("Failed to update user's data in database: %s", err.Error())
	}
}

// SetActive sets user's active status.
func (u *User) SetActive() {
	u.IsActive = true
	_, err := database.DB.NamedExec("UPDATE `users` SET is_active=:is_active WHERE id=:id", u)
	if err != nil {
		log.Error().Msgf("Failed to set user's active status in database: %s", err.Error())
	}
}
