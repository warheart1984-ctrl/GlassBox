package ledger

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

func SignRecord(record Record, secret []byte) (Record, error) {
	data, err := canonicalRecordPayload(record)
	if err != nil {
		return record, err
	}

	mac := hmac.New(sha256.New, secret)
	_, _ = mac.Write(data)
	record.Signature = mac.Sum(nil)
	return record, nil
}

func VerifyRecord(record Record, secret []byte) bool {
	expected, err := SignRecord(Record{
		ID:        record.ID,
		Type:      record.Type,
		Timestamp: record.Timestamp,
		Payload:   record.Payload,
	}, secret)
	if err != nil {
		return false
	}
	return hmac.Equal(record.Signature, expected.Signature)
}

func EncodeSignature(signature []byte) string {
	return base64.StdEncoding.EncodeToString(signature)
}

func canonicalRecordPayload(record Record) ([]byte, error) {
	unsigned := struct {
		ID        string     `json:"id"`
		Type      RecordType `json:"type"`
		Timestamp string     `json:"timestamp"`
		Payload   any        `json:"payload"`
	}{
		ID:        record.ID,
		Type:      record.Type,
		Timestamp: record.Timestamp.UTC().Format("2006-01-02T15:04:05.000000000Z07:00"),
		Payload:   record.Payload,
	}
	return json.Marshal(unsigned)
}
