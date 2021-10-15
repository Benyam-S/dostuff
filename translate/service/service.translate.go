package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Benyam-S/dostuff/log"
	"github.com/Benyam-S/dostuff/translate"
)

// TranslateService is a type that defines a translation service
type TranslateService struct {
	store  map[string]string
	logger *log.Logger
}

// NewTranslateService is a function that returns a new translation service
func NewTranslateService(store map[string]string, logger *log.Logger) translate.IService {
	return &TranslateService{store: store, logger: logger}
}

// TranslateStuff is a function that translates a 'stuff' to certain location language's equivalent
func (service *TranslateService) TranslateStuff(location string) (string, error) {

	/* ---------------------------- Logging ---------------------------- */
	service.logger.Log(fmt.Sprintf("Started translating 'stuff' to client language { Client Location : %s }",
		location), service.logger.Logs.DebugLogFile)

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
		/* ---------------------------- Logging ---------------------------- */
		service.logger.Log(fmt.Sprintf("Error: unable to translate 'stuff' to client language "+
			"{ Client Location : %s }", location), service.logger.Logs.ErrorLogFile)

		return "", errors.New("unable to translate stuff")
	}

	/* ---------------------------- Logging ---------------------------- */
	service.logger.Log(fmt.Sprintf("Finished translating 'stuff' to client language, "+
		"{ Client Location : %s, Translation : %s }", location, translation), service.logger.Logs.DebugLogFile)

	return translation, nil
}
