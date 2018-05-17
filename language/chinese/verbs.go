package chinese

type Verbs struct {
	Base
	verb string
	subject Base
	object Base
}

//	Verbs's functions
func (verbs *Verbs) Verb() string{
	return verbs.verb
}

func (verbs *Verbs) SetVerb(verb string) {
	verbs.verb = verb
}

func (verbs *Verbs) Subject() Base {
	return verbs.subject
}

func (verbs *Verbs) SetSubject(subject Base) {
	verbs.subject = subject
}

func (verbs *Verbs) Object() Base {
	return verbs.object
}

func (verbs *Verbs) SetObject(object Base) {
	verbs.object = object
}

//	Base's functions override in Verbs
func (verbs Verbs) Show() {
	verbs.Verb()
}