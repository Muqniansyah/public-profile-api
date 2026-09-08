package config

import "os"

// getEnv mengambil nilai Environment Variable dan jika tidak tersedia, menggunakan nilai default.
func GetEnv(key string, fallback string) string {

    value := os.Getenv(key)

    if value == "" {
        return fallback
    }

    return value
}