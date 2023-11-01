package repository

import (
	"motivation-app/kutipan_motivasi/pkg/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DatabaseRepository adalah repository yang berinteraksi dengan basis data.
type DatabaseRepository struct {
	db *gorm.DB
}

// NewDatabaseRepository digunakan untuk membuat instansiasi DatabaseRepository.
func NewDatabaseRepository() (*DatabaseRepository, error) {
	// Membuka koneksi ke basis data SQLite (gantilah dengan basis data yang sesuai)
	db, err := gorm.Open(sqlite.Open("motivation.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrasi model-model ke basis data (tabel-tabel)
	db.AutoMigrate(&model.Quote{}, &model.User{})

	return &DatabaseRepository{db: db}, nil
}

// SaveQuote digunakan untuk menyimpan kutipan motivasi ke basis data.
func (dr *DatabaseRepository) SaveQuote(quote *model.Quote) error {
	return dr.db.Create(quote).Error
}

// GetUserByID digunakan untuk mendapatkan pengguna berdasarkan ID.
func (dr *DatabaseRepository) GetUserByID(userID int) (*model.User, error) {
	var user model.User
	if err := dr.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
