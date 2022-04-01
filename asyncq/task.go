package asyncq

type Task interface {
	Perform()
	StampQueuedAt()
	StampScanningAt()
	StampFinishedAt()
}
