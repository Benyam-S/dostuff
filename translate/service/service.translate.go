package service

import (
	"errors"
	"strings"

	"github.com/Benyam-S/dostaff/translate"
)

// TranslateService is a type that defines a translation service
type TranslateService struct {
	store map[string]string
}

// NewTranslateService is a function that returns a new translation service
func NewTranslateService(store map[string]string) translate.IService {
	return &TranslateService{store: store}
}

// TranslateStuff is a function that translates a 'stuff' to certain location language's equivalent
func (service *TranslateService) TranslateStuff(location string) (string, error) {

	var canBeTranslated bool
	var translation string

	for key, value := range service.store {
		if strings.ToLower(key) == strings.ToLower(location) {
			canBeTranslated = true
			translation = value
			break
		}
	}

	if !canBeTranslated {
		return "", errors.New("unable to translate stuff")
	}

	return translation, nil
}
