package metadata

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func UpdateMetadataDate(path string, info os.FileInfo, fileDateTime *time.Time) {
	if strings.HasSuffix(strings.ToLower(path), ".jpg") ||
		strings.HasSuffix(strings.ToLower(path), ".jpeg") ||
		strings.HasSuffix(strings.ToLower(path), ".gif") {
		updateMetadataDateJpg(path, info, fileDateTime)
	}
}

func updateMetadataDateJpg(filepath string, info os.FileInfo, fileDateTime *time.Time) {
	fmt.Println("    - file: ", filepath)
	var existingFileDateTime = extractExifMetadataDate(filepath, info)
	fmt.Println("      newDateTime: ", fileDateTime)
	if existingFileDateTime != nil {
		fmt.Println("      existingDateTime: ", existingFileDateTime)
	}
	setExifMetadataDate(filepath, *fileDateTime)
	err := os.Chtimes(filepath, *fileDateTime, *fileDateTime)
	if err != nil {
		fmt.Println(err)
	}
}
