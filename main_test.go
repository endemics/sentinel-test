package main

import (
	"os"
	"testing"

	"github.com/hashicorp/sentinel-sdk"
	plugintesting "github.com/hashicorp/sentinel-sdk/testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	plugintesting.Clean()
	os.Exit(exitCode)
}

func TestImport(t *testing.T) {
	cases := []struct {
		Name   string
		Data   map[string]interface{}
		Source string
	}{
		{
			"month",
			map[string]interface{}{"timestamp": 1495483674},
			`main = subject.month.string == "September"`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			plugintesting.TestImport(t, plugintesting.TestImportCase{
				Config: tc.Data,
				Source: tc.Source,
			})
		})
	}
}
