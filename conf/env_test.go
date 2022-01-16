package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv_Env(t *testing.T) {
	t.Parallel()

	tests := []struct {
		env    string
		expect string
	}{
		{env: EnvProd, expect: EnvProd},
		{env: EnvDev, expect: EnvDev},
		{env: EnvTest, expect: EnvTest},
		{env: "", expect: EnvDev},
	}

	for _, test := range tests {
		t.Setenv("GO_ENV", test.env)
		assert.Equal(t, Env(), test.expect)
	}
}

func TestEnv_IsProd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		env    string
		expect bool
	}{
		{env: EnvProd, expect: true},
		{env: EnvDev, expect: false},
		{env: EnvTest, expect: false},
		{env: "", expect: false},
	}

	for _, test := range tests {
		t.Setenv("GO_ENV", test.env)
		assert.Equal(t, IsProd(), test.expect)
	}
}

func TestEnv_IsDev(t *testing.T) {
	t.Parallel()

	tests := []struct {
		env    string
		expect bool
	}{
		{env: EnvProd, expect: false},
		{env: EnvDev, expect: true},
		{env: EnvTest, expect: false},
		{env: "", expect: true},
	}

	for _, test := range tests {
		t.Setenv("GO_ENV", test.env)
		assert.Equal(t, IsDev(), test.expect)
	}
}

func TestEnv_IsTest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		env    string
		expect bool
	}{
		{env: EnvProd, expect: false},
		{env: EnvDev, expect: false},
		{env: EnvTest, expect: true},
		{env: "", expect: false},
	}

	for _, test := range tests {
		t.Setenv("GO_ENV", test.env)
		assert.Equal(t, IsTest(), test.expect)
	}
}
