package junit_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/pierreprinetti/go-junit"
)

func TestSuites_Unmarshal(t *testing.T) {
	tests := []struct {
		desc        string
		suites      junit.TestSuites
		noXMLHeader bool
		goVersion   string
	}{
		{
			desc: "Suites should marshal back and forth",
			suites: junit.TestSuites{
				Suites: []junit.TestSuite{
					{
						Name: "suite1",
						TestCases: []junit.TestCase{
							{Name: "test1-1"},
							{Name: "test1-2"},
						},
					},
					{
						Name: "suite2",
						TestCases: []junit.TestCase{
							{Name: "test2-1"},
							{Name: "test2-2"},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Logf("Test case: %v", test.desc)
		initialBytes, err := xml.Marshal(test.suites)
		if err != nil {
			t.Fatalf("Expected no failure when generating xml; got %v", err)
		}

		var suites junit.TestSuites
		err = xml.Unmarshal(initialBytes, &suites)
		if err != nil {
			t.Fatalf("Expected no failure when unmarshaling; got %v", err)
		}

		newBytes, err := xml.Marshal(suites)
		if err != nil {
			t.Fatalf("Expected no failure when generating xml again; got %v", err)
		}

		if !bytes.Equal(newBytes, initialBytes) {
			t.Errorf("Expected the same result when marshal/unmarshal/marshal. Expected\n%v\n\t but got\n%v",
				string(initialBytes),
				string(newBytes),
			)
		}
	}
}
