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

package migrations

import (
	// stdlib
	"database/sql"
)

func InitialUp(tx *sql.Tx) error {
	if _, err := tx.Exec("CREATE TABLE `users` (`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'User ID', `login` varchar(191) NOT NULL COMMENT 'User login', `email` varchar(191) NOT NULL COMMENT 'User email', `password` varchar(64) NOT NULL COMMENT 'User password', `password_salt` varchar(64) NOT NULL COMMENT 'Password salt', `is_active` BOOL NOT NULL COMMENT 'Is user active? 1 - yes, 0 - banned', `created_at` datetime NOT NULL COMMENT 'User creation timestamp', `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User last update time', PRIMARY KEY (`id`), UNIQUE KEY `id` (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Users';"); err != nil {
		return err
	}

	if _, err1 := tx.Exec("CREATE TABLE `sessions` (`id` int(11) NOT NULL COMMENT 'User ID', `key` varchar(64) NOT NULL COMMENT 'Session key', `issued` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'When session key was issued') ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Sessions keys'"); err1 != nil {
		return err1
	}

	if _, err2 := tx.Exec("CREATE TABLE `packages` (`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Package ID', `name` text NOT NULL COMMENT 'Package name to show on HTML', `original_package_url` text NOT NULL COMMENT 'Original package URL as in import line', `created_at` datetime NOT NULL COMMENT 'Timestamp when package was created', `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp when package was last updated', PRIMARY KEY (`id`), UNIQUE KEY `id` (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='Packages basic info'"); err2 != nil {
		return err2
	}

	if _, err3 := tx.Exec("CREATE TABLE `packages_urls` (`package_id` int(11) NOT NULL COMMENT 'Package ID', `url` text NOT NULL COMMENT 'Sources URL', `enabled` boolean NOT NULL DEFAULT true COMMENT 'Is redirection enabled?') ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Packages URLs'"); err3 != nil {
		return err3
	}

	return nil
}

func InitialDown(tx *sql.Tx) error {
	if _, err := tx.Exec("DROP TABLE `users`;"); err != nil {
		return err
	}

	if _, err1 := tx.Exec("DROP TABLE `sessions`;"); err1 != nil {
		return err1
	}

	if _, err2 := tx.Exec("DROP TABLE `packages`;"); err2 != nil {
		return err2
	}

	if _, err3 := tx.Exec("DROP TABLE `packages`;"); err3 != nil {
		return err3
	}

	return nil
}
