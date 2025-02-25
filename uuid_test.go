package jsonschema

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestUUIDFuncTest(t *testing.T) {
	r := Reflector{
		UuidFunc: func(parentUuid uuid.UUID, field reflect.StructField) uuid.UUID {
			pUuid := uuid.NameSpaceOID
			if parentUuid != uuid.Nil {
				pUuid = parentUuid
			}

			return uuid.NewSHA1(pUuid, []byte(field.Name))
		},
	}

	schema := r.Reflect(struct {
		Test  string
		Dummy int
	}{})

	s, _ := json.Marshal(schema)
	fmt.Println(string(s))
}
