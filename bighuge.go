package thesaurus

import (
	"encoding/json"
	"errors"
	"net/http"
)

//BigHuge contains the return value of the api from words.bughugelabs.com
type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

//Synonyms returns the synonyms of the passed in term
func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string

	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	if err != nil {
		return syns, errors.New("bighuge:  Failed when looking for synonyms for \"" + term + "\"" + err.Error())
	}

	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}

	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)

	return syns, nil
}
