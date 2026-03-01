package service

import (
	"datagps/internal/models"
	"datagps/internal/repository"
	"log"
	"strconv"
)

type AppService struct {
	dataRepo       repository.DataRepository
	groupRepo      repository.GroupRepository
	serviceStellar *AppServiceStellar
}

func NewAppService(dt repository.DataRepository, gr repository.GroupRepository, ss *AppServiceStellar) *AppService {
	return &AppService{dt, gr, ss}
}

func (s *AppService) CrearteData(d *models.Data) error {
	return s.dataRepo.Create(d)
}

func (s *AppService) PendingGroups(limit int) ([]models.Group, error) {
	return s.groupRepo.PendingGroups(limit)
}

func (s *AppService) ProcessPendingGroups(limit int) ([]models.Group, error) {
	groups, err := s.groupRepo.PendingGroups(limit)
	if err != nil {
		log.Fatal(err)
	}

	var groupsNew []models.Group
	for _, group := range groups {
		key := strconv.FormatUint(uint64(group.Id), 10)
		group.HashBlkc = s.serviceStellar.SaveData(key, group.HashGroup)
		s.dataRepo.UpdateBlck(int(group.IdStart), int(group.IdFinish), int(group.ID))
		s.groupRepo.Save(&group)
		groupsNew = append(groupsNew, group)
	}

	return groupsNew, nil
}

func (s *AppService) ExecSetGroup() error {
	return s.groupRepo.ExecSetGroup()
}

func (s *AppService) Groups(page int, totalPage int, records int) ([]models.Group, int, error) {
	return s.groupRepo.Groups(page, totalPage, records)
}
