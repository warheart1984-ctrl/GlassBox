package ledger

import (
	"encoding/json"
	"os"
)

type FileLedger struct {
	Path string
}

func (l FileLedger) Append(record Record) error {
	f, err := os.OpenFile(l.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.Marshal(record)
	if err != nil {
		return err
	}
	if _, err := f.Write(append(data, '\n')); err != nil {
		return err
	}
	return nil
}
