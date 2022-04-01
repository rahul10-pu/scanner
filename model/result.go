package model

import "time"

type ScannerResult struct {
	Id             string
	Status         string
	RepositoryName string
	RepositoryUrl  string
	Findings       string
	QueuedAt       time.Time
	ScanningAt     time.Time
	FinishedAt     time.Time
}

func (sr *ScannerResult) Perform() {

}
func (sr *ScannerResult) StampQueuedAt() {
	sr.QueuedAt = time.Now()
}
func (sr *ScannerResult) StampScanningAt() {
	sr.ScanningAt = time.Now()
}
func (sr *ScannerResult) StampFinishedAt() {
	sr.FinishedAt = time.Now()
}
