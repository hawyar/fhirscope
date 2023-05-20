package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	examples := map[string]Scope{
		"patient/Observation.rs?param=true&anotha=1": {
			Context:    PatientContext,
			Resource:   "Observation",
			Operations: []Operation{ReadOperation, SearchOperation},
			Params: map[string]string{
				"param":  "true",
				"anotha": "1",
			},
		},
		"patient/*.cud": {
			Context:    PatientContext,
			Resource:   Wildcard,
			Operations: []Operation{CreateOperation, UpdateOperation, DeleteOperation},
		},
		"system/Encounter.crud": {
			Context:    SystemContext,
			Resource:   "Encounter",
			Operations: []Operation{CreateOperation, ReadOperation, UpdateOperation, DeleteOperation},
		},
		"user/Observation.read?foo=bar": {
			Context:    UserContext,
			Resource:   "Observation",
			Operations: []Operation{ReadOperation, SearchOperation},
			Params: map[string]string{
				"foo": "bar",
			},
		},
		"patient/*.*": {
			Context:    PatientContext,
			Resource:   Wildcard,
			Operations: []Operation{CreateOperation, ReadOperation, UpdateOperation, DeleteOperation, SearchOperation},
		},
	}

	for input, expected := range examples {
		scope, err := Parse(input)

		if err != nil {
			t.Error(err)
		}

		if scope.Context != expected.Context {
			t.Errorf("Expected context %v, got %v", expected.Context, scope.Context)
		}

		if scope.Resource != expected.Resource {
			t.Errorf("Expected resource %v, got %v", expected.Resource, scope.Resource)
		}

		if len(scope.Operations) != len(expected.Operations) {
			t.Errorf("Expected %v operations, got %v", len(expected.Operations), len(scope.Operations))
		}

		if len(scope.Params) != len(expected.Params) {
			t.Errorf("Expected %v params, got %v", len(expected.Params), len(scope.Params))
		}

		for k, v := range scope.Params {
			if expected.Params[k] != v {
				t.Errorf("Expected param %v = %v, got %v", k, expected.Params[k], v)
			}
		}

		for i, op := range scope.Operations {
			if op != expected.Operations[i] {
				t.Errorf("Expected operation %v, got %v", expected.Operations[i], op)
			}
		}
	}
}
