package storage

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Storage interface {
	RecordSessionSize(fileSize float64, fileType string)
	IncreaseStorageTotalSessions(fileType string)
	RecordSessionUploadDuration(durMillis float64, fileType, mode string)
	List() []prometheus.Collector
}

type storageImpl struct{}

func New(serviceName string) Storage { return &storageImpl{} }

func (s *storageImpl) List() []prometheus.Collector                                         { return []prometheus.Collector{} }
func (s *storageImpl) RecordSessionSize(fileSize float64, fileType string)                  {}
func (s *storageImpl) IncreaseStorageTotalSessions(fileType string)                         {}
func (s *storageImpl) RecordSessionUploadDuration(durMillis float64, fileType, mode string) {}
