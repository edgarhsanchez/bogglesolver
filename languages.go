package main

// SupportedLanguage encapsulates the code and the language name for each language supported
type SupportedLanguage struct {
	//Lang is the code of the language en_US, es_ES, etc...
	Lang string
	//Language is the pretty name in each relative language of each language
	Language string
}

// SupportedLanguageResponse encapsulates the response for the supported list of languages
type SupportedLanguageResponse struct {
	SupportedLanguages []SupportedLanguage `json:"supportedLanguages"`
}
