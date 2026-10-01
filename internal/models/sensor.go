package models

import "time"

type SensorReading struct {
	ID          int64
	DeviceID    int64
	BinType     BinType
	DistanceCM  float64
	FillPercent *float64
	ReceivedAt  time.Time
	DeviceTS    *time.Time
	RawPayload  []byte
}

type LatestBinReading struct {
	BinType BinType
	Status  BinReadingStatus
	Reading *SensorReading
}
