package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadSingleFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		// want     *Paifu
	}{
		{"scc20240710", "../data/tenhou_html/scc20240710.html"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadSingleFile(tt.filePath)
			if err != nil {
				t.Error(err)
			}
			assert.Nil(t, err)
			assert.NotNil(t, got)
			t.Log(got)
		})
	}
}
