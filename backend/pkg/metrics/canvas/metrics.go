package canvas

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Canvas interface {
	RecordCanvasImageSize(size float64)
	IncreaseTotalSavedImages()
	RecordCanvasesPerSession(number float64)
	RecordPreparingDuration(duration float64)
	IncreaseTotalCreatedArchives()
	RecordUploadingDuration(duration float64)
	List() []prometheus.Collector
}

type canvasImpl struct{}

func New(serviceName string) Canvas { return &canvasImpl{} }

func (c *canvasImpl) List() []prometheus.Collector             { return []prometheus.Collector{} }
func (c *canvasImpl) RecordCanvasImageSize(size float64)       {}
func (c *canvasImpl) IncreaseTotalSavedImages()                {}
func (c *canvasImpl) RecordCanvasesPerSession(number float64)  {}
func (c *canvasImpl) RecordPreparingDuration(duration float64) {}
func (c *canvasImpl) IncreaseTotalCreatedArchives()            {}
func (c *canvasImpl) RecordUploadingDuration(duration float64) {}
