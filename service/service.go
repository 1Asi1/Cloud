package service

import (
	"log"
	"solution/model"
	"solution/store"
)

type Service struct {
	conf    *model.Config
	version float64
	db      store.Store
}

func NewService(s store.Store) *Service {
	ser := &Service{conf: model.NewConfig(), db: s}
	return ser
}

func (s *Service) CreateConfig(h *model.Config) {
	s.db.Create(h)
}

func (s *Service) ReadConfig(h *model.Config, ser string) {
	err := s.db.Read(h, ser, &s.version)
	if err != nil {
		log.Println("was not found config", err)
		return
	}
}

func (s *Service) UpdateConfig(h *model.Config) {
	err := s.db.Update(h)
	if err != nil {
		log.Println("has not been update config", err)
	}
}

func (s *Service) DeleteConfig(ser string, ver float64) {
	err := s.db.Delete(s.version, ser, ver)
	if err != nil {
		log.Println("has not been delete config", err)
	}
}

func (s *Service) GetVersion() float64 {
	return s.version
}

func (s *Service) SetLastVersion(ser string) {
	err := s.db.Read(s.conf, ser, &s.version)
	if err != nil {
		log.Println("was not found config")
		return
	}
}
