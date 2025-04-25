package env

import (
	"os"
	"strconv"
)

func GetEnvAsString(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

func GetEnvAsBoolean(key string, defaultValue bool) (bool, error) {
	if v := os.Getenv(key); v != "" {
		value, err := strconv.ParseBool(v)
		if err != nil {
			return defaultValue, err
		}
		return value, nil
	}
	return defaultValue, nil
}

func GetEnvAsInt(key string, defaultValue int) (int, error) {
	if v := os.Getenv(key); v != "" {
		value, err := strconv.Atoi(v)
		if err != nil {
			return defaultValue, err
		}
		return value, nil
	}
	return defaultValue, nil
}

func GetEnvAsInt64(key string, defaultValue int64) (int64, error) {
	if v := os.Getenv(key); v != "" {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return defaultValue, err
		}
		return value, nil
	}
	return defaultValue, nil
}

func GetEnvAsFloat64(key string, defaultValue float64) (float64, error) {
	if v := os.Getenv(key); v != "" {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return defaultValue, err
		}
		return value, nil
	}
	return defaultValue, nil
}
