package messages

import (
	"github.com/vube/i18n"
)

func getTranslator(lang_key string) *i18n.Translator {
	f, _ := i18n.NewTranslatorFactory(
    	[]string{"./messages"},
    	[]string{"./messages"},
    	lang_key,
	)

	t, _ := f.GetTranslator(lang_key)	

	return t
}

func GetLocaleMessage(lang_key string, message string) string {
	t := getTranslator(lang_key)	

	localeString, _ := t.Translate(message, map[string]string{})	

	return localeString
}