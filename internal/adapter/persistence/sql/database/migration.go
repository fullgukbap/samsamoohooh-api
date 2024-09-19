package database

import "samsamoohooh-mini-api/internal/core/domain"

func AutoMigrate(d *Database) error {
	return d.AutoMigrate(&domain.User{})
}
