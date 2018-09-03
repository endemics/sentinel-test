package main

import (
	"fmt"
	"github.com/hashicorp/sentinel-sdk"
	"github.com/hashicorp/sentinel-sdk/framework"
	"github.com/hashicorp/sentinel-sdk/rpc"
	"time"
)

func main() {
	rpc.Serve(&rpc.ServeOpts{
		ImportFunc: func() sdk.Import {
			return &framework.Import{Root: &root{}}
		},
	})
}

type root struct {
	time time.Time
}

// framework.Root impl.
func (m *root) Configure(raw map[string]interface{}) error {
	if _, ok := raw["timestamp"]; !ok {
		raw["timestamp"] = time.Now().Unix()
	}

	v := raw["timestamp"]
	timestamp, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid timestamp type %T", v)
	}

	m.time = time.Unix(timestamp, 0).UTC()
	return nil
}

// framework.Namespace impl.
func (m *root) Get(key string) (interface{}, error) {
	switch key {
	case "minute":
		return m.time.Minute(), nil
	}

	return nil, nil
}
