package images

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Images interface {
	RecordSavingImageDuration(duration float64)
	IncreaseTotalSavedImages()
	IncreaseTotalCreatedArchives()
	RecordUploadingDuration(duration float64)
	List() []prometheus.Collector
}

type imagesImpl struct{}

func New(serviceName string) Images { return &imagesImpl{} }

func (i *imagesImpl) List() []prometheus.Collector               { return []prometheus.Collector{} }
func (i *imagesImpl) RecordSavingImageDuration(duration float64) {}
func (i *imagesImpl) IncreaseTotalSavedImages()                  {}
func (i *imagesImpl) IncreaseTotalCreatedArchives()              {}
func (i *imagesImpl) RecordUploadingDuration(duration float64)   {}
