package ledger

import (
	"os"
	"testing"
)

func TestSignAndVerifyRecord(t *testing.T) {
	secret := []byte("test-secret")
	record := NewRecord(RecordTransition, map[string]any{"input": "hello"})

	signed, err := SignRecord(record, secret)
	if err != nil {
		t.Fatalf("SignRecord returned error: %v", err)
	}

	if len(signed.Signature) == 0 {
		t.Fatal("expected signature to be populated")
	}
	if !VerifyRecord(signed, secret) {
		t.Fatal("expected signed record to verify")
	}
}

func TestVerifyRecordRejectsTamperedPayload(t *testing.T) {
	secret := []byte("test-secret")
	record := NewRecord(RecordTransition, map[string]any{"input": "hello"})
	signed, err := SignRecord(record, secret)
	if err != nil {
		t.Fatalf("SignRecord returned error: %v", err)
	}

	signed.Payload = map[string]any{"input": "changed"}

	if VerifyRecord(signed, secret) {
		t.Fatal("expected tampered record to fail verification")
	}
}

func TestFileLedgerAppendsJSONLRecord(t *testing.T) {
	path := t.TempDir() + "/ledger.jsonl"
	record := NewRecord(RecordTransition, map[string]any{"input": "hello"})
	signed, err := SignRecord(record, []byte("test-secret"))
	if err != nil {
		t.Fatalf("SignRecord returned error: %v", err)
	}

	if err := (FileLedger{Path: path}).Append(signed); err != nil {
		t.Fatalf("Append returned error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}
	if len(data) == 0 || data[len(data)-1] != '\n' {
		t.Fatal("expected appended JSONL record ending in newline")
	}
}
