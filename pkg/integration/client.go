package integration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Rustixir/go-challenge/pkg/config"
)

func SendRequest(method string, path string, model any, response any) error {
	url := fmt.Sprintf("http://localhost:%d/api%s", config.Config.Port, path)
	var body io.Reader
	if model != nil {
		raw, _ := json.Marshal(model)
		body = bytes.NewBuffer(raw)
	}
	req, _ := http.NewRequest(method, url, body)
	if model != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		raw, _ := io.ReadAll(resp.Body)
		return errors.New(string(raw))
	}

	if response == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(response)
}
