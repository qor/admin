package admin_test

import (
	"encoding/xml"
	"errors"
	"testing"

	admin "github.com/TykTechnologies/raava-admin"
)

func TestXMLTransformerEncode(t *testing.T) {
	t.Skip()
	xmlResult := admin.XMLStruct{
		Result: map[string]interface{}{"error": errors.New("error message"), "status": map[string]int{"code": 200}},
	}
	result := "<response>\n\t<error>error message</error>\n\t<status>\n\t\t<code>200</code>\n\t</status>\n</response>"

	if xmlMarshalResult, err := xml.MarshalIndent(xmlResult, "", "\t"); err != nil {
		t.Errorf("no error should happen, but got %v", err)
	} else if string(xmlMarshalResult) != result {
		t.Errorf("Generated XML got %v, but should be %v", string(xmlMarshalResult), result)
	}
}
