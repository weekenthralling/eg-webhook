package storage

import (
	"gorm.io/gorm"
	"zjuici.com/tablegpt/eg-webhook/models"
)

type KernelSessionStore struct {
	db *gorm.DB
}

func NewKernelSessionStore(db *gorm.DB) *KernelSessionStore {
	return &KernelSessionStore{
		db: db,
	}
}

func (s KernelSessionStore) SaveSession(kernel *models.KernelSession) error {
	result := s.db.Save(kernel)
	return result.Error
}

// GetSessionByID retrieve the kernel session by kernel id
func (s KernelSessionStore) GetSessionByID(id string) (*models.KernelSession, error) {
	var kernel models.KernelSession
	result := s.db.First(&kernel, "id = ?", id)
	return &kernel, result.Error
}

// DeleteSessionsByID batch delete kernel session by id
func (s KernelSessionStore) DeleteSessionsByID(id []string) error {
	result := s.db.Unscoped().Delete(&models.KernelSession{}, "id IN ?", id)
	return result.Error
}

// ListSessions retrieve a list of all kernel sessions from a database
func (s KernelSessionStore) ListSessions() ([]models.KernelSession, error) {
	var kernels []models.KernelSession
	result := s.db.Find(&kernels)
	return kernels, result.Error
}
