package services

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

type VoiceManager struct{}

func NewVoiceManager() *VoiceManager {
	return &VoiceManager{}
}

func GetRandomVoiceBytes(rootDir string) []byte {
	var files []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	randomIndex := rand.Intn(len(files))
	randomFile := files[randomIndex]

	fileContent, err := os.ReadFile(randomFile)
	if err != nil {
		log.Fatal(err)
	}

	return fileContent
}
