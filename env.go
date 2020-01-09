package env

import (
	"os"
	"strconv"
	"time"
)

type Value struct {
	v string
}

func Get(key string) *Value {
	return &Value{
		v: os.Getenv(key),
	}
}

func (v *Value) String(def string) string {
	if v.v == "" {
		return def
	}

	return v.v
}

func (v *Value) Int(def int) int {
	val, err := strconv.Atoi(v.v)

	if err != nil || val == 0 {
		return def
	}

	return val
}

func (v *Value) Bool(def bool) bool {
	val, err := strconv.ParseBool(v.v)

	if err != nil {
		return def
	}

	return val
}

func (v *Value) Duration(def time.Duration) time.Duration {
	val, err := time.ParseDuration(v.v)

	if err != nil || val == 0 {
		return def
	}

	return val
}
