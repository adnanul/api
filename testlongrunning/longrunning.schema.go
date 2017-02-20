package testlongrunning

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var longRunningRequestSchema *gojsonschema.Schema

func init() {
	var err error
	longRunningRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "duration": {
      "type": "integer"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *LongRunningRequest) IsValid() (*gojsonschema.Result, error) {
	return longRunningRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *LongRunningRequest) IsRequest() {}

