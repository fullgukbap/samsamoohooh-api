package database

import "samsamoohooh-mini-api/internal/core/domain"

func AutoMigration(d *Database) error {
	return d.AutoMigrate(&domain.User{})
}
