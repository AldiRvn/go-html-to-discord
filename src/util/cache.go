package util

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
)

var idCache map[string][]string

func AddCache(web string, id string) {
	idCache[web] = append(idCache[web], id)
}

func CheckIdExists(web string, id string) bool {
	for _, existing := range idCache[web] {
		if existing == id {
			return true
		}
	}
	return false
}

func LoadCache(path string) {
	// Pastikan folder-nya ada
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		slog.Error("❌ Gagal buat folder cache", "err", err)
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		slog.Error(err.Error())
		idCache = map[string][]string{}
		return
	}
	_ = json.Unmarshal(data, &idCache)
	if idCache == nil {
		idCache = map[string][]string{}
	}
}

func SaveCache(path string) {
	data, _ := json.MarshalIndent(idCache, "", "  ")
	_ = os.WriteFile(path, data, os.ModePerm)
}
