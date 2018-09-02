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

package sessionkeys

import (
	// stdlib
	"strings"
	"time"

	// local
	"github.com/welltrainedfolks/magister/internal/config"
	"github.com/welltrainedfolks/magister/internal/database"
	"github.com/welltrainedfolks/magister/internal/helpers"

	// other
	"github.com/rs/zerolog/log"
)

type SessionKey struct {
	ID     int       `db:"id"`
	Key    string    `db:"key"`
	Issued time.Time `db:"issued"`
}

func DeleteSessionKey(key string) {
	_ = database.DB.MustExec("DELETE FROM `sessions` WHERE `key`=?", key)
}

func NewSessionKey(uid int) string {
	log.Debug().Msg("Generating new session key")
	sk := SessionKey{
		ID:     uid,
		Issued: time.Now(),
	}

	// Generate key.
	sk.Key = helpers.GenerateRandomString(32)
	// Put in database.
	_, err := database.DB.NamedExec("INSERT INTO `sessions` (`id`, `key`, `issued`) VALUES (:id, :key, :issued)", &sk)
	if err != nil {
		log.Error().Msgf("Failed to save session data in database: %s", err.Error())
		return ""
	}

	return sk.Key
}

func CheckSessionKey(key string) (int, bool) {
	// Check if session key exist.
	sk := &SessionKey{}
	err := database.DB.Get(sk, "SELECT * FROM `sessions` WHERE `key`=?", key)
	if err != nil {
		if !strings.Contains(err.Error(), "sql: no rows") {
			log.Error().Msgf("Failed to check session key validity: %s", err.Error())
			return -1, false
		}
	}

	// Check key validity.
	curTime := time.Now().UTC()
	log.Debug().Msgf("Current time: %v <> Token issuance time: %v <> Sub: %v <> Valid for: %v", curTime, sk.Issued, curTime.Sub(sk.Issued), time.Hour*time.Duration(24*config.Config.HTTP.SessionValidityDays))
	if curTime.Sub(sk.Issued) > time.Hour*time.Duration(24*config.Config.HTTP.SessionValidityDays) {
		log.Debug().Msgf("Session expired.")
		return -1, false
	}

	return sk.ID, true
}
