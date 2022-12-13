package configuration

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

func CalculatePath(path string) string {
	re := regexp.MustCompile(`\$\{(.*?)\}`)
	matches := re.FindStringSubmatch(path)
	currentTime := time.Now()
	for _, match := range matches {
		switch match {
		case "date":
			date := currentTime.Format("20060102")
			path = strings.ReplaceAll(path, fmt.Sprintf("${%s}", match), date)
		case "time":
			time := currentTime.Format("150405")
			path = strings.ReplaceAll(path, fmt.Sprintf("${%s}", match), time)
		case "datetime":
			datetime := currentTime.Format("20060102150405")
			path = strings.ReplaceAll(path, fmt.Sprintf("${%s}", match), datetime)
		default:
			log.Printf("Unknown match for path : %s", match)
		}
	}
	return path
}
