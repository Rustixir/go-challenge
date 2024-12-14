package localization

import (
	"encoding/json"
	"log/slog"
)

type schema struct {
	dict map[string]map[string]string
}

func newSchema() *schema {
	return &schema{
		dict: make(map[string]map[string]string),
	}
}

func (s *schema) load(raw []byte) {
	var data map[string]map[string]string
	err := json.Unmarshal(raw, &data)
	if err != nil {
		slog.Error("failed to load localization file", "error", err)
		return
	}
	s.dict = data
}
