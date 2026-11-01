package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertMorse(s string) (string, error) {
	if s == "" {
		return "", errors.New("пустая строка")
	}
	if strings.TrimLeft(s, ".- ") == "" {
		return morse.ToText(s), nil
	}

	return morse.ToMorse(s), nil
}
