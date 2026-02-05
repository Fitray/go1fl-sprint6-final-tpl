package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertText(text string) string {
	for _, ch := range text {
		if !strings.ContainsRune(".- ", ch) {
			return morse.ToMorse(text)
		}
	}
	return morse.ToText(text)
}
