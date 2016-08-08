package thesaurus

//Thesaurus defines the interface for searching for synonyms
type Thesaurus interface {
	Synonyms(term string) ([]error, error)
}
