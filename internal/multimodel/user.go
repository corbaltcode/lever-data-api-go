package multimodel

import (
	"encoding/json"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// Unmarshal a user ID or user.
//   - If the raw message is empty, returns ("", nil, nil).
//   - If the raw message is a string, returns (id, nil, nil).
//   - If the raw message is a user, returns (user.ID, &user, nil).
func unmarshalUserOrID(raw json.RawMessage) (string, *model.User, error) {
	if len(raw) == 0 {
		return "", nil, nil
	}

	var user model.User

	// Try unmarshalling as a user first
	if err := json.Unmarshal(raw, &user); err != nil {
		// Can't unmarshal as a user; try unmarshalling as an ID.
		var id string
		if err := json.Unmarshal(raw, &id); err != nil {
			return "", nil, err
		}

		return id, nil, nil
	}

	return user.ID, &user, nil
}

// Unmarshal an array of user IDs or an array of users.
//   - If the raw message is empty, returns (nil, nil, nil).
//   - If the raw message is an array of strings, returns (ids, nil, nil).
//   - If the raw message is an array of users, returns (ids, users, nil).
func unmarshalArrayOfUsersOrIDs(raw json.RawMessage) ([]string, []model.User, error) {
	if len(raw) == 0 {
		return nil, nil, nil
	}

	var ids []string
	var users []model.User

	// Try unmarshalling as an array of users first
	if err := json.Unmarshal(raw, &users); err != nil {
		// Can't unmarshal as users; try unmarshalling as IDs.
		if err := json.Unmarshal(raw, &ids); err != nil {
			return nil, nil, err
		}

		return ids, nil, nil
	}

	for _, user := range users {
		ids = append(ids, user.ID)
	}

	return ids, users, nil
}
