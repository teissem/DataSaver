package configuration

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

func CalculateDestination(destination string) string {
	re := regexp.MustCompile(`\$\{(.*?)\}`)
	matches := re.FindStringSubmatch(destination)
	currentTime := time.Now()
	for _, match := range matches {
		switch match {
		case "date":
			date := currentTime.Format("20060102")
			destination = strings.ReplaceAll(destination, fmt.Sprintf("${%s}", match), date)
		case "time":
			time := currentTime.Format("150405")
			destination = strings.ReplaceAll(destination, fmt.Sprintf("${%s}", match), time)
		case "datetime":
			datetime := currentTime.Format("20060102150405")
			destination = strings.ReplaceAll(destination, fmt.Sprintf("${%s}", match), datetime)
		default:
			log.Printf("Unknown match for destination : %s", match)
		}
	}
	return destination
}
