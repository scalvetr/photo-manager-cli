package metadata

import (
	"fmt"
	"strings"
	"time"
)

func CheckAlbumDateMismatch(filePath string, year int, month int, dateMismatch func(filePath string, fileDate *time.Time)) {
	if strings.HasSuffix(strings.ToLower(filePath), ".jpg") ||
		strings.HasSuffix(strings.ToLower(filePath), ".jpeg") ||
		strings.HasSuffix(strings.ToLower(filePath), ".gif") {
		checkAlbumDateMismatchJpg(filePath, year, month, dateMismatch)
	} else if strings.HasSuffix(strings.ToLower(filePath), ".mpg") ||
		strings.HasSuffix(strings.ToLower(filePath), ".mpeg") ||
		strings.HasSuffix(strings.ToLower(filePath), ".mp4") {
		checkAlbumDateMismatchMpeg(filePath, year, month, dateMismatch)
	}
}
func checkAlbumDateMismatchMpeg(filePath string, year int, month int, dateMismatch func(filePath string, fileDate *time.Time)) {
	fmt.Println("    - file: ", filePath)
	var existingFileDateTime = getModTime(filePath)
	fmt.Printf("year => \nalbum: %s \nfile: %s", year, existingFileDateTime.Year())
	fmt.Printf("month => \nalbum: %s \nfile: %s", month, existingFileDateTime.Month())
	if year != existingFileDateTime.Year() || month != int(existingFileDateTime.Month()) {
		dateMismatch(filePath, existingFileDateTime)
	}
}
func checkAlbumDateMismatchJpg(filePath string, year int, month int, dateMismatch func(filePath string, fileDate *time.Time)) {
	fmt.Println("    - file: ", filePath)
	var existingFileDateTime = extractExifMetadataDate(filePath)
	if existingFileDateTime == nil {
		fmt.Println("no date specified")
		dateMismatch(filePath, nil)
	} else {
		fmt.Printf("year => \n  album: %s \n  file: %s\n", year, existingFileDateTime.Year())
		fmt.Printf("month => \n  album: %s \n  file: %s\n", month, int(existingFileDateTime.Month()))

		if year != existingFileDateTime.Year() || month != int(existingFileDateTime.Month()) {
			dateMismatch(filePath, existingFileDateTime)
		}
	}
}
