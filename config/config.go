package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

type BaseConfig struct {
	Port           string `env:"PORT,8080"`
	Env            string `env:"ENV,dev"`
	Namespace      string `env:"NAMESPACE,default"`
	KubeConfigPath string `env:"KUBECONFIG_PATH,config.yml"`
}

// LoadConfig loads environment variables from .env and populates a struct
func LoadConfig(c interface{}) error {
	// Load from .env file if present
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	// Reflect on the struct
	s := reflect.ValueOf(c)
	t := s.Type()

	// Loop through struct fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("env")

		// Skip fields without env tag
		if tag == "" {
			continue
		}

		// Extract default value from tag
		parts := strings.SplitN(tag, ",", 2)
		defaultValue := ""
		if len(parts) > 1 {
			defaultValue = parts[1]
		}

		// Get environment variable value
		value := os.Getenv(parts[0])

		// Use default value if env var is not set
		if value == "" {
			value = defaultValue
		}

		// Set field value
		f := s.Field(i)
		if !f.CanSet() {
			return fmt.Errorf("field %s is not settable", field.Name)
		}
		switch f.Kind() {
		case reflect.String:
			f.SetString(value)
		case reflect.Int:
			var intValue int
			if _, err := fmt.Sscan(value, &intValue); err != nil {
				return fmt.Errorf("error converting %s to int: %w", value, err)
			}
			f.SetInt(int64(intValue))
		default:
			return fmt.Errorf("unsupported field type %s for env var %s", f.Kind(), tag)
		}
	}

	return nil
}
