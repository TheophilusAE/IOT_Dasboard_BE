package models

import "time"

type DetectionEvent struct {
	ID         int64
	DeviceID   int64
	WasteType  BinType
	Confidence float64
	ReceivedAt time.Time
	DeviceTS   *time.Time
	RawPayload []byte
}

type BinReadingStatus string

const (
	BinStatusHasData BinReadingStatus = "has_data"
	BinStatusNoData  BinReadingStatus = "no_data"
)
