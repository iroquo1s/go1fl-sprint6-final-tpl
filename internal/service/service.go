package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DetectContent(data []byte) (string, error) {

	detectMorse := strings.ContainsAny(string(data), ".- ")

	if detectMorse {
		return morse.ToText(string(data)), nil
	} else {
		return morse.ToMorse(string(data)), nil
	}

}
