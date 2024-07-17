package db

import (
	"em/internal/config"
	"em/internal/db/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repository struct {
	config *config.PSQLConnection
	logger *log.Logger
}

func NewRepository(psql *config.PSQLConnection, logger *log.Logger) *repository {
	return &repository{
		config: psql,
		logger: logger,
	}
}

func (r *repository) Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		r.config.Host, r.config.Port, r.config.Username, r.config.Password, r.config.Database)), &gorm.Config{})

	if err != nil {
		r.logger.Fatal(err)
	}

	return db
}

func ApplyMigrations(db *gorm.DB) {
	db.Migrator().DropTable(&models.TaskModel{})
	db.Migrator().DropTable(&models.UserModel{})

	db.Migrator().AutoMigrate(&models.UserModel{}, &models.TaskModel{})
}
