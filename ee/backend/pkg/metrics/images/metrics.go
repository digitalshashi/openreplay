package images

import (
	"github.com/prometheus/client_golang/prometheus"

	"openreplay/backend/pkg/metrics/common"
)

type Images interface {
	RecordSavingImageDuration(duration float64)
	IncreaseTotalSavedImages()
	IncreaseTotalCreatedArchives()
	RecordUploadingDuration(duration float64)
	List() []prometheus.Collector
}

type imagesImpl struct {
	savingImageDuration  prometheus.Histogram
	totalSavedImages     prometheus.Counter
	totalCreatedArchives prometheus.Counter
	uploadingDuration    prometheus.Histogram
}

func New(serviceName string) Images {
	return &imagesImpl{
		savingImageDuration:  newSavingImageDuration(serviceName),
		totalSavedImages:     newTotalSavedImages(serviceName),
		totalCreatedArchives: newTotalCreatedArchives(serviceName),
		uploadingDuration:    newUploadingDuration(serviceName),
	}
}

func (i *imagesImpl) List() []prometheus.Collector {
	return []prometheus.Collector{
		i.savingImageDuration,
		i.totalSavedImages,
		i.totalCreatedArchives,
		i.uploadingDuration,
	}
}

func newSavingImageDuration(serviceName string) prometheus.Histogram {
	return prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: serviceName,
			Name:      "saving_image_duration_seconds",
			Help:      "A histogram displaying the duration of saving each image in seconds.",
			Buckets:   common.DefaultDurationBuckets,
		},
	)
}

func (i *imagesImpl) RecordSavingImageDuration(duration float64) {
	i.savingImageDuration.Observe(duration)
}

func newTotalSavedImages(serviceName string) prometheus.Counter {
	return prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: serviceName,
			Name:      "total_saved_images",
			Help:      "A counter displaying the total number of saved images.",
		},
	)
}

func (i *imagesImpl) IncreaseTotalSavedImages() {
	i.totalSavedImages.Inc()
}

func newTotalCreatedArchives(serviceName string) prometheus.Counter {
	return prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: serviceName,
			Name:      "total_created_archives",
			Help:      "A counter displaying the total number of created archives.",
		},
	)
}

func (i *imagesImpl) IncreaseTotalCreatedArchives() {
	i.totalCreatedArchives.Inc()
}

func newUploadingDuration(serviceName string) prometheus.Histogram {
	return prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: serviceName,
			Name:      "uploading_duration_seconds",
			Help:      "A histogram displaying the wall-clock duration of streaming-archive upload to S3 in seconds.",
			Buckets:   common.DefaultDurationBuckets,
		},
	)
}

func (i *imagesImpl) RecordUploadingDuration(duration float64) {
	i.uploadingDuration.Observe(duration)
}
