package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	examples := map[string]Scope{
		"patient/Observation.rs?param=true": {
			Context:    PatientContext,
			Resource:   "Observation",
			Operations: []Operation{ReadOperation, SearchOperation},
		},
	}

	for input, expected := range examples {
		scope, err := Parse(input)

		if err != nil {
			t.Error(err)
		}

		fmt.Println(expected.Operations)
		fmt.Println(scope.Operations)
	}
}
