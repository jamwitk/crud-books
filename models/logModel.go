package models

type RequestLog struct {
	ID        uint   `gorm:"primarykey"`
	UserAgent string `json:"userAgent" validate:"required"`
	SessionID string `json:"sessionID" validate:"required"`
	DateTime  string `json:"dateTime" validate:"required" validate:"datetime=2006-01-02 15:04:05"`
	Request   string `json:"request" validate:"required"`
	Response  string `json:"response" validate:"required"`
}
type EventLog struct {
	ID          uint   `gorm:"primarykey"`
	EventName   string `json:"eventName" validate:"required"`
	Source      string `json:"source" validate:"required"`
	Tags        string `json:"tags" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserAgent   string `json:"userAgent" validate:"required"`
	SessionID   string `json:"sessionID" validate:"required"`
	DateTime    string `json:"dateTime" validate:"required" validate:"datetime=2006-01-02 15:04:05"`
}
type DataBaseLog struct {
	ID          uint   `gorm:"primarykey"`
	DateTime    string `json:"dateTime" validate:"required" validate:"datetime=2006-01-02 15:04:05"`
	TableAction string `json:"tableAction" validate:"required"`
	Table       string `json:"table" validate:"required"`
	Query       string `json:"query" validate:"required"`
	Response    string `json:"response" validate:"required"`
	UserAgent   string `json:"userAgent" validate:"required"`
}
