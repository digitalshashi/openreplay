package storage

import (
	"github.com/prometheus/client_golang/prometheus"

	"openreplay/backend/pkg/metrics/common"
)

type Storage interface {
	RecordSessionSize(fileSize float64, fileType string)
	IncreaseStorageTotalSessions(fileType string)
	RecordSessionUploadDuration(durMillis float64, fileType, mode string)
	List() []prometheus.Collector
}

type storageImpl struct {
	sessionSize           *prometheus.HistogramVec
	totalSessions         *prometheus.CounterVec
	sessionUploadDuration *prometheus.HistogramVec
}

func New(serviceName string) Storage {
	return &storageImpl{
		sessionSize:           newSessionSize(serviceName),
		totalSessions:         newTotalSessions(serviceName),
		sessionUploadDuration: newSessionUploadDuration(serviceName),
	}
}

func (s *storageImpl) List() []prometheus.Collector {
	return []prometheus.Collector{
		s.sessionSize,
		s.totalSessions,
		s.sessionUploadDuration,
	}
}

func newSessionSize(serviceName string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: serviceName,
			Name:      "session_size_bytes",
			Help:      "A histogram displaying the size of each session file in bytes prior to any manipulation.",
			Buckets:   common.DefaultSizeBuckets,
		},
		[]string{"file_type"},
	)
}

func (s *storageImpl) RecordSessionSize(fileSize float64, fileType string) {
	s.sessionSize.WithLabelValues(fileType).Observe(fileSize)
}

func newTotalSessions(serviceName string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: serviceName,
			Name:      "sessions_total",
			Help:      "A counter displaying the total number of session files uploaded, partitioned by file_type.",
		},
		[]string{"file_type"},
	)
}

func (s *storageImpl) IncreaseStorageTotalSessions(fileType string) {
	s.totalSessions.WithLabelValues(fileType).Inc()
}

func newSessionUploadDuration(serviceName string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: serviceName,
			Name:      "upload_duration_seconds",
			Help:      "A histogram displaying the wall-clock duration of the streaming upload pipeline (read+compress/encrypt+S3 PUT) in seconds.",
			Buckets:   common.DefaultDurationBuckets,
		},
		[]string{"file_type", "mode"},
	)
}

func (s *storageImpl) RecordSessionUploadDuration(durMillis float64, fileType, mode string) {
	s.sessionUploadDuration.WithLabelValues(fileType, mode).Observe(durMillis / 1000.0)
}
