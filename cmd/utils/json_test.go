package utils_test

import (
	"testing"

	"github.com/peterldowns/testy/assert"
	"github.com/ryands17/go-bytes/cmd/utils"
)

func TestCustomMarshalStruct(t *testing.T) {
	t.Run("Test the function with struct fields in a different order", func(t *testing.T) {
		type TestStruct struct {
			Field1 string `json:"field1"`
			Field2 int    `json:"field2"`
		}

		type TestStruct2 struct {
			Field2 int    `json:"field2"`
			Field1 string `json:"field1"`
		}

		teststructPayload, err := utils.MarshalStruct(TestStruct{
			Field1: "value1",
			Field2: 42,
		})

		assert.NoError(t, err)

		teststructPayload2, err := utils.MarshalStruct(TestStruct2{
			Field1: "value1",
			Field2: 42,
		})
		assert.NoError(t, err)

		assert.Equal(t, teststructPayload, teststructPayload2)
	})

	t.Run("Test struct with omitempty fields", func(t *testing.T) {
		type TestStruct struct {
			Required string `json:"required"`
			Empty    string `json:"empty,omitempty"`
			ZeroInt  int    `json:"zeroint"`
		}

		test := TestStruct{
			Required: "value",
			ZeroInt:  0, // this should not be omitted
		}

		result, err := utils.MarshalStruct(test)
		assert.NoError(t, err)
		assert.Equal(t, `{"required":"value","zeroint":0}`, string(result))
	})
}
