package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
)

const filePath = "/app/static/tasks/_index.md"

func TimeWorker() {
	ticker := time.NewTicker(5 * time.Second)
	counter := 0

	for {
		select {
		case <-ticker.C:
			currentTime := time.Now().Format("2006-01-02 15:04:05")

			currentContent, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Println("Ошибка чтения файла:", err)
				continue
			}

			updatedContent := replaceText(currentContent, `Текущее время:.*`, fmt.Sprintf("Текущее время: %s", currentTime))
			updatedContent = replaceText(updatedContent, `Счетчик:.*`, fmt.Sprintf("Счетчик: %d", counter))

			err = os.WriteFile(filePath, updatedContent, 0644)
			if err != nil {
				log.Println("Ошибка записи файла:", err)
				continue
			}

			counter++
		}
	}
}

func replaceText(content []byte, pattern string, replacement string) []byte {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAll(content, []byte(replacement))
}
