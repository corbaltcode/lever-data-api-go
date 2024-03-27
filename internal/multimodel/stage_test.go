package multimodel

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyStage(t *testing.T) {
	ta := assert.New(t)

	id, stage, err := unmarshalStageOrID(json.RawMessage(""))
	if ta.NoError(err) {
		ta.Empty(id)
		ta.Nil(stage)
	}
}

func TestStageUnmarshalling(t *testing.T) {
	ta := assert.New(t)

	id, stage, err := unmarshalStageOrID(json.RawMessage(`{
	"id": "8d49b010-cc6a-4f40-ace5-e86061c677ed",
	"text": "Interview"
}`))

	ta.NoError(err)
	ta.Equal(id, "8d49b010-cc6a-4f40-ace5-e86061c677ed")
	ta.NotNil(stage)
	ta.Equal(stage.ID, "8d49b010-cc6a-4f40-ace5-e86061c677ed")

	id, stage, err = unmarshalStageOrID(json.RawMessage(`  "8d49b010-cc6a-4f40-ace5-e86061c677ed"  `))
	ta.NoError(err)
	ta.Equal(id, "8d49b010-cc6a-4f40-ace5-e86061c677ed")
	ta.Nil(stage)
}

func TestInvalidStageUnmarshalling(t *testing.T) {
	ta := assert.New(t)

	_, _, err := unmarshalStageOrID(json.RawMessage(`[]`))
	ta.Error(err)
}
