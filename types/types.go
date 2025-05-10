package types

import "time"

type SeverityLevel string

const (
	SeverityInfo    SeverityLevel = "info"
	SeverityWarning SeverityLevel = "warning"
	SeverityError   SeverityLevel = "error"
	SeverityFatal   SeverityLevel = "fatal"
)

// Event represents a system event
type Event struct {
	ID        string        `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Severity  SeverityLevel `gorm:"type:varchar(10);not null" json:"severity"`
	Title     string        `gorm:"type:varchar(255);not null" json:"title"`
	Message   string        `gorm:"type:text;not null" json:"message"`
	CreatedAt time.Time     `gorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
	Source    string        `gorm:"type:varchar(50);not null" json:"source"`
	Tags      interface{}   `gorm:"type:jsonb;not null" json:"tags"`
	EventTime time.Time     `gorm:"type:timestamp;not null" json:"event_time"`
}
