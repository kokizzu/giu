// Package main shows using of translations in giu.
package main

import "github.com/AllenDang/giu"

var (
	// here we define our multi-language dictionary.
	// You could also do that e.g. with JSON.
	languageDefs = map[string]map[string]string{
		"en": {}, // as we write our UI in english, the default value of the text will be fine (see (*BasicTranslator).Translate for more)
		"pl": {
			"Hello world!": "Witaj świecie",
		},
		"de": {
			"Hello world!": "Hallo Welt!",
		},
	}

	languageCodes = []string{"en", "pl", "de"}
	currentLang   int32
)

func loop() {
	giu.SingleWindow().Layout(
		giu.Combo("Select language", languageCodes[currentLang], languageCodes, &currentLang).OnChange(func() {
			if err := giu.Context.Translator.SetLanguage(languageCodes[currentLang]); err != nil {
				panic(err)
			}
		}),
		giu.Label("Hello world!"),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Hello world", 800, 600, giu.MasterWindowFlagsNotResizable)
	// initialize translation. Do that before loop.
	translator := giu.NewBasicTranslator()
	for k, v := range languageDefs {
		translator.AddLanguage(k, v)
	}

	if err := translator.SetLanguage(languageCodes[currentLang]); err != nil {
		panic(err)
	}

	giu.Context.SetTranslator(translator)

	wnd.Run(loop)
}
