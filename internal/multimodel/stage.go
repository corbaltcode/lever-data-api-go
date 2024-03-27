package multimodel

import (
	"encoding/json"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// Unmarshal a stage ID or stage.
//
// If the raw message is empty, returns ("", nil, nil).
// If the raw message is a string, returns (id, nil, nil).
// If the raw message is a stage, returns (stage.ID, &stage, nil).
func unmarshalStageOrID(raw json.RawMessage) (string, *model.Stage, error) {
	if len(raw) == 0 {
		return "", nil, nil
	}

	// Try unmarshalling as a stage first
	var stage model.Stage
	if err := json.Unmarshal(raw, &stage); err != nil {
		// Can't unmarshal as a stage; try unmarshalling as an ID.
		var id string
		if err := json.Unmarshal(raw, &id); err != nil {
			return "", nil, err
		}

		return id, nil, nil
	}

	return stage.ID, &stage, nil
}
