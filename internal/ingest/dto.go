package ingest

import "time"

type HeartbeatDTO struct {
	UptimeSeconds   *int64     `json:"uptime_seconds"`
	RSSI            *int32     `json:"rssi"`
	FreeHeapBytes   *int32     `json:"free_heap_bytes"`
	FirmwareVersion string     `json:"firmware_version"`
	DeviceTS        *time.Time `json:"device_ts"`
}

type DetectionDTO struct {
	WasteType  string     `json:"waste_type"`
	Confidence float64    `json:"confidence"`
	DeviceTS   *time.Time `json:"device_ts"`
}

type SensorReadingDTO struct {
	BinType     string     `json:"bin_type"`
	DistanceCM  float64    `json:"distance_cm"`
	FillPercent *float64   `json:"fill_percent"`
	DeviceTS    *time.Time `json:"device_ts"`
}

type SensorBatchDTO struct {
	Readings []SensorReadingDTO `json:"readings"`
}
