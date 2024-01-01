package helpers

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


// GetAsString mengembalikan nilai string dari variabel lingkungan dengan nama tertentu.
// Jika variabel tidak ada, kembalikan nilai default.
func GetAsString(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

// GetAsInt mengembalikan nilai integer dari variabel lingkungan dengan nama tertentu.
// Jika variabel tidak ada atau tidak dapat dikonversi menjadi integer, kembalikan nilai default.
func GetAsInt(name string, defaultValue int) int {
	valueStr := GetAsString(name, "")
	value, err := strconv.Atoi(valueStr)
	if err == nil {
		return value
	}
	return defaultValue
}

	// LoadEnv digunakan untuk memuat variabel lingkungan dari file .env
	func LoadEnv(path string) error {
		err := godotenv.Load(path)
		if err != nil {
			return err
		}
		return nil
	}