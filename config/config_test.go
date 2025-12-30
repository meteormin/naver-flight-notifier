package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg, err := New("../config.toml")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "https://flight-api.naver.com", cfg.Naver.URL)
}
