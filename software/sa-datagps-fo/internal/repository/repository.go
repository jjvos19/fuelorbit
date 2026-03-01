package repository

import (
	"datagps/internal/models"
	"math"

	"github.com/google/uuid"
	"github.com/stellar/go-stellar-sdk/support/log"
	"gorm.io/gorm"
)

type DataRepository interface {
	Create(data *models.Data) error
	UpdateBlck(idStart int, idFinish int, idHashBlck int) (records int, err error)
}
type GroupRepository interface {
	Create(group *models.Group) error
	PendingGroups(limit int) ([]models.Group, error)
	ExecSetGroup() error
	Save(*models.Group) error
	Groups(page int, totalPages int, records int) ([]models.Group, int, error)
}

type dataRepo struct {
	db *gorm.DB
}

type groupRepo struct {
	db     *gorm.DB
	dtRepo dataRepo
}

func NewRepository(db *gorm.DB) (DataRepository, GroupRepository) {
	return &dataRepo{db}, &groupRepo{db, dataRepo{db}}
}

func (r *dataRepo) Create(data *models.Data) error {
	return r.db.Create(&data).Error
}

func (r *dataRepo) UpdateBlck(idStart int, idFinish int, idHashBlck int) (records int, err error) {
	result := r.db.Where("id between ? and ?", idStart, idFinish).Updates(models.Data{GroupBlck: idHashBlck})
	records = int(result.RowsAffected)
	err = result.Error
	return records, err
}

func (r *groupRepo) Create(group *models.Group) error {
	return r.db.Create(&group).Error
}

func (r *groupRepo) PendingGroups(limit int) ([]models.Group, error) {
	var groups []models.Group
	err := r.db.Where("hash_blkc IS NULL").Order("id ASC").Limit(limit).Find(&groups).Error
	return groups, err
}

func (r *groupRepo) Save(group *models.Group) error {
	return r.db.Save(&group).Error
}

func (r *groupRepo) ExecSetGroup() error {
	result := r.db.Exec("call sp_set_group()")
	return result.Error
}

func (r *groupRepo) Groups(page int, totalPages int, records int) ([]models.Group, int, error) {
	var groups []models.Group
	var count int64
	if totalPages == 0 {
		r.db.Table("tkr_group").Count(&count)
		pagesCount := float64(count) / float64(records)
		totalPages = int(math.Ceil(pagesCount))
		log.Infof("page: %d totalPages: %d records: %d PagesCount: %f", page, totalPages, records, pagesCount)
	}
	err := r.db.Order("id ASC").Limit(records).Offset((page - 1) * records).Find(&groups).Error
	return groups, totalPages, err

}

func getUUID() string {
	return uuid.NewString()
}
