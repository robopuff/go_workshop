package internal

type Entries []struct {
	Word      string      `json:"word"`
	Phonetics []Phonetics `json:"phonetics"`
	Meanings  []Meaning   `json:"meanings"`
}

type Phonetics struct {
	Text  string `json:"text"`
	Audio string `json:"audio"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type Definition struct {
	Definition string   `json:"definition"`
	Example    string   `json:"example"`
	Synonyms   []string `json:"synonyms"`
}
