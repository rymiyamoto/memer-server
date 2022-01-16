package conf_test

import (
	"testing"

	"github.com/rymiyamoto/memer-server/conf"
	"github.com/stretchr/testify/assert"
)

func TestEnv_Env(t *testing.T) {
	t.Parallel()

	tests := []struct {
		env    string
		expect string
	}{
		{env: conf.EnvProd, expect: conf.EnvProd},
		{env: conf.EnvDev, expect: conf.EnvDev},
		{env: conf.EnvTest, expect: conf.EnvTest},
		{env: "", expect: conf.EnvDev},
	}

	for _, test := range tests {
		t.Setenv("GO_ENV", test.env)
		assert.Equal(t, conf.Env(), test.expect)
	}
}

func TestEnv_IsProd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		env    string
		expect bool
	}{
		{env: conf.EnvProd, expect: true},
		{env: conf.EnvDev, expect: false},
		{env: conf.EnvTest, expect: false},
		{env: "", expect: false},
	}

	for _, test := range tests {
		t.Setenv("GO_ENV", test.env)
		assert.Equal(t, conf.IsProd(), test.expect)
	}
}

func TestEnv_IsDev(t *testing.T) {
	t.Parallel()

	tests := []struct {
		env    string
		expect bool
	}{
		{env: conf.EnvProd, expect: false},
		{env: conf.EnvDev, expect: true},
		{env: conf.EnvTest, expect: false},
		{env: "", expect: true},
	}

	for _, test := range tests {
		t.Setenv("GO_ENV", test.env)
		assert.Equal(t, conf.IsDev(), test.expect)
	}
}

func TestEnv_IsTest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		env    string
		expect bool
	}{
		{env: conf.EnvProd, expect: false},
		{env: conf.EnvDev, expect: false},
		{env: conf.EnvTest, expect: true},
		{env: "", expect: false},
	}

	for _, test := range tests {
		t.Setenv("GO_ENV", test.env)
		assert.Equal(t, conf.IsTest(), test.expect)
	}
}
