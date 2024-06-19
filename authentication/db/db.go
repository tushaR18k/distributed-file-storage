package db

import (
	"authentication/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresStore struct {
	DB *gorm.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		log.Fatal("Environment variable POSTGRES_PASSWORD is not set")
	}
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		log.Fatal("Environment variable POSTGRES_HOST is not set")
	}

	connStr := fmt.Sprintf("user=postgres dbname=postgres sslmode=disable password=%s host=%s", password, host)
	DB, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		DB: DB,
	}, nil
}

func (s *PostgresStore) StoreUser(user models.User) error {
	if err := s.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) FindUser(username string) (*models.User, error) {
	var user models.User
	s.DB.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}
