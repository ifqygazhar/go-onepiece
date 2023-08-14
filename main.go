package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

)

const (
	baseURL      = "https://onepieceberwarna.com/komik/"
	minVolume    = 5
	maxVolume    = 108
	minChapter   = 37
	maxChapter   = 1090
	maxImage     = 69
)

func main() {
	for numvol := minVolume; numvol <= maxVolume; numvol++ {
		for chapter := minChapter; chapter <= maxChapter; chapter++ {
			chapterFailed := false

			for urutanGambar := 1; urutanGambar <= maxImage; urutanGambar++ {
				imageURL := fmt.Sprintf("%sVOL%%20%d/%d/%02d.jpg", baseURL, numvol, chapter, urutanGambar)
				response, err := http.Get(imageURL)

				if err == nil && response.StatusCode == http.StatusOK {
					chapterFailed = false
					folderPath := fmt.Sprintf("VOL_%d/Chapter_%d", numvol, chapter)
					err := os.MkdirAll(folderPath, os.ModePerm)

					if err != nil {
						fmt.Printf("Error creating folder: %v\n", err)
						continue
					}

					filePath := path.Join(folderPath, fmt.Sprintf("Image_%02d.jpg", urutanGambar))
					file, err := os.Create(filePath)

					if err != nil {
						fmt.Printf("Error creating file: %v\n", err)
						continue
					}

					_, err = io.Copy(file, response.Body)
					if err != nil {
						fmt.Printf("Error copying image data: %v\n", err)
					}

					file.Close()
					response.Body.Close()
					fmt.Printf("Downloaded: %s\n", imageURL)
				} else {
					fmt.Printf("Failed to download: %s\n", imageURL)
					chapterFailed = true
				}

				if chapterFailed && urutanGambar == 1 {
					break
				}
			}
		}
	}
}
