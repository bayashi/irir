package main

import (
	"encoding/json"

	"github.com/invopop/jsonschema"
)

func dumpJSONSchema() string {
	s := jsonschema.Reflect(map[string][]*Rule{"": []*Rule{&Rule{}}})
	data, _ := json.MarshalIndent(s, "", "  ")

	return string(data)
}
