package database

import "samsamoohooh-api/internal/core/domain"

func AutoMigrate(d *Database) error {
	return d.AutoMigrate(&domain.User{})
}
