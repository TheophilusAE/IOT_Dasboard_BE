package ingest

import (
	"fmt"
	"time"

	"iot-backend/internal/models"
)

func normalizeConfidence(raw float64) (float64, error) {
	v := raw
	if v > 1 {
		v = v / 100
	}
	if v < 0 || v > 1 {
		return 0, fmt.Errorf("confidence out of range after normalization: %v (raw: %v)", v, raw)
	}
	return v, nil
}

func mapDetection(deviceID int64, dto DetectionDTO, rawPayload []byte) (detectionInsert, error) {
	binType, ok := models.NormalizeBinType(dto.WasteType)
	if !ok {
		return detectionInsert{}, fmt.Errorf("unrecognized waste_type: %q", dto.WasteType)
	}

	confidence, err := normalizeConfidence(dto.Confidence)
	if err != nil {
		return detectionInsert{}, err
	}

	return detectionInsert{
		DeviceID:   deviceID,
		WasteType:  binType,
		Confidence: confidence,
		DeviceTS:   dto.DeviceTS,
		RawPayload: rawPayload,
	}, nil
}

func mapSensorReading(deviceID int64, dto SensorReadingDTO, rawPayload []byte) (sensorInsert, error) {
	binType, ok := models.NormalizeBinType(dto.BinType)
	if !ok {
		return sensorInsert{}, fmt.Errorf("unrecognized bin_type: %q", dto.BinType)
	}

	const maxSaneDistanceCM = 500
	if dto.DistanceCM < 0 || dto.DistanceCM > maxSaneDistanceCM {
		return sensorInsert{}, fmt.Errorf("distance_cm out of sane range: %v", dto.DistanceCM)
	}

	if dto.FillPercent != nil && (*dto.FillPercent < 0 || *dto.FillPercent > 100) {
		return sensorInsert{}, fmt.Errorf("fill_percent out of range: %v", *dto.FillPercent)
	}

	return sensorInsert{
		DeviceID:    deviceID,
		BinType:     binType,
		DistanceCM:  dto.DistanceCM,
		FillPercent: dto.FillPercent,
		DeviceTS:    dto.DeviceTS,
		RawPayload:  rawPayload,
	}, nil
}

type detectionInsert struct {
	DeviceID   int64
	WasteType  models.BinType
	Confidence float64
	DeviceTS   *time.Time
	RawPayload []byte
}

type sensorInsert struct {
	DeviceID    int64
	BinType     models.BinType
	DistanceCM  float64
	FillPercent *float64
	DeviceTS    *time.Time
	RawPayload  []byte
}
