package ws

type OutgoingMessage struct {
	Type string `json:"type"`
	Data any    `json:"data,omitempty"`
}

type SubscribeMessage struct {
	Type      string  `json:"type"`
	DeviceIDs []int64 `json:"device_ids"`
}

const (
	TypeSubscribed    = "subscribed"
	TypeDetection     = "detection_event"
	TypeSensorReading = "sensor_reading"
	TypeDeviceStatus  = "device_status"
	TypePing          = "ping"
)
