package repository

import (
	"crave/hub/cmd/model"
	"crave/shared/database"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	mysql *database.MysqlWrapper
}

func NewRepository(mysql *database.MysqlWrapper) *Repository {
	return &Repository{mysql: mysql}
}

func (r *Repository) GetLastIndex() (uint, error) {
	var lastIndex uint
	result := r.mysql.Driver.Table("work").Select("COALESCE(MAX(id), 0) as last_index").Scan(&lastIndex)
	if result.Error != nil {
		return 0, result.Error
	}
	return lastIndex, nil
}

func (r *Repository) Create(work *model.Work) (*model.Work, error) {
	var createdWork model.Work
	err := r.mysql.Driver.Transaction(func(tx *gorm.DB) error {
		// Create the work with a lock
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Create(work).Error; err != nil {
			return err
		}

		// Fetch the last created work with a lock
		if err := tx.Clauses(clause.Locking{Strength: "SHARE"}).Last(&createdWork).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &createdWork, nil
}

func (r *Repository) FindById(id uint16) (*model.Work, error) {
	var work model.Work
	result := r.mysql.Driver.First(&work, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("🛢️work with id %d not found", id)
		}
		return nil, fmt.Errorf("error fetching work: %w", result.Error)
	}
	return &work, nil
}

func (r *Repository) Update(work *model.Work) (*model.Work, error) {
	result := r.mysql.Driver.Model(&model.Work{}).
		Where("id = ?", work.Id).
		Updates(work)
	if result.Error != nil {
		return nil, result.Error
	}
	return work, nil
}
