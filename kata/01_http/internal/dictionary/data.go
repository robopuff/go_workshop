package dictionary

type Entries []struct {
	Word      string      `json:"word"`
	Phonetic  string      `json:"phonetic,omitempty"`
	Phonetics []Phonetics `json:"phonetics"`
	Meanings  []Meaning   `json:"meanings"`
}

type Phonetics struct {
	Text  string `json:"text,omitempty"`
	Audio string `json:"audio,omitempty"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech,omitempty"`
	Definitions  []Definition `json:"definitions,omitempty"`
}

type Definition struct {
	Definition string   `json:"definition,omitempty"`
	Example    string   `json:"example,omitempty"`
	Synonyms   []string `json:"synonyms,omitempty"`
}
