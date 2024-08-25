package pkg

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/adjust/rmq/v5"
)

type RMQueues = map[string]*rmq.Queue
type Targets map[string]string

type QueueItem struct {
	Header   http.Header `json:"headers"`
	Path     string      `json:"path"`
	Query    url.Values  `json:"q"`
	Endpoint string      `json:"endpoint"`
}

// MarshalBinary encodes the struct into a binary blob
func (u *QueueItem) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalBinary decodes the struct into a QueueItem
func (u *QueueItem) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &u); err != nil {
		return err
	}
	return nil
}
