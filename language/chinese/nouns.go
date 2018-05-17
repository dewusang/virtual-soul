package chinese

type Nouns struct {
	Base
	noun string
}

//	Nouns's functions
func (nouns *Nouns) Noun() string {
	return nouns.noun
}

func (nouns *Nouns) SetNoun(noun string) {
	nouns.noun = noun
}

//	Base's functions override in Nouns
func (nouns Nouns) Show() {
	nouns.Noun()
}