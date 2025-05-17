package extractor

import "log"

type Extracted struct {
	Result        map[Identifier]string
	DatasNotFound []Identifier
}

func (e *Extractor) Result() Extracted {
	extracted := Extracted{
		Result:        e.result,
		DatasNotFound: nil,
	}

	for _, t := range e.tokens {
		if !t.IsExtracted {
			log.Printf("token not found - token: '%s'", identifiersNames[t.Identifier])
			extracted.DatasNotFound = append(extracted.DatasNotFound, t.Identifier)
		}

		// This if and logs still for debuggin porpuses now
		// but we can use something related to say may we extracted wrong bro
		// because we have a idea of how much long is the data that we want to extract
		// for each token.
		if len(t.Value) >= 55 {
			log.Printf("maybe token value is incorrect - token: '%s'", identifiersNames[t.Identifier])
			log.Printf("token value: '%s'", t.Value)
		}
	}

	return extracted
}
