package ledger

import (
	"time"

	"glass-box/internal/util"
)

type RecordType string

const (
	RecordTransition RecordType = "transition"
	RecordClassifier RecordType = "classifier"
	RecordConstraint RecordType = "constraint"
)

type Record struct {
	ID        string     `json:"id"`
	Type      RecordType `json:"type"`
	Timestamp time.Time  `json:"timestamp"`
	Payload   any        `json:"payload"`
	Signature []byte     `json:"signature"`
}

func NewRecord(recordType RecordType, payload any) Record {
	return Record{
		ID:        util.NewID("rec"),
		Type:      recordType,
		Timestamp: time.Now().UTC(),
		Payload:   payload,
	}
}
