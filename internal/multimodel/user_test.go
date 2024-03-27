package multimodel

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyUserUnmarshalling(t *testing.T) {
	ta := assert.New(t)

	id, user, err := unmarshalUserOrID(json.RawMessage(""))
	if ta.NoError(err) {
		ta.Empty(id)
		ta.Nil(user)
	}
}

func TestUserUnmarshalling(t *testing.T) {
	ta := assert.New(t)

	id, user, err := unmarshalUserOrID(json.RawMessage(`{
	"id": "8d49b010-cc6a-4f40-ace5-e86061c677ed",
	"name": "Chandler Bing",
	"username": "chandler",
	"email": "chandler@example.com",
	"createdAt": 1407357447018,
	"deactivatedAt": 1409556487918,
	"externalDirectoryId": "2277399",
	"accessRole": "super admin",
	"photo": "https://gravatar.com/avatar/gp781413e3bb44143bddf43589b03038?s=26&d=404",
	"linkedContactIds": [
		"38f608d5-9a60-4960-83c1-99d18f40c428"
	]
}`))
	ta.NoError(err)
	ta.Equal(id, "8d49b010-cc6a-4f40-ace5-e86061c677ed")
	ta.NotNil(user)
	ta.Equal(user.ID, "8d49b010-cc6a-4f40-ace5-e86061c677ed")

	id, user, err = unmarshalUserOrID(json.RawMessage(`  "8d49b010-cc6a-4f40-ace5-e86061c677ed"  `))
	ta.NoError(err)
	ta.Equal(id, "8d49b010-cc6a-4f40-ace5-e86061c677ed")
	ta.Nil(user)
}

func TestEmptyUserArrayUnmarshalling(t *testing.T) {
	ta := assert.New(t)

	ids, users, err := unmarshalArrayOfUsersOrIDs(json.RawMessage(""))
	if ta.NoError(err) {
		ta.Empty(ids)
		ta.Empty(users)
	}
}

func TestInvalidUserUnmarshalling(t *testing.T) {
	ta := assert.New(t)

	id, user, err := unmarshalUserOrID(json.RawMessage(`[[]]`))
	if ta.Error(err) {
		ta.Empty(id)
		ta.Nil(user)
	}
}

func TestUserArrayUnmarshalling(t *testing.T) {
	ta := assert.New(t)

	ids, users, err := unmarshalArrayOfUsersOrIDs(json.RawMessage(`[
	{
		"id": "8d49b010-cc6a-4f40-ace5-e86061c677ed",
		"name": "Chandler Bing",
		"username": "chandler",
		"email": "chandler@example.com",
		"createdAt": 1407357447018,
		"deactivatedAt": 1409556487918,
		"externalDirectoryId": "2277399",
		"accessRole": "super admin",
		"photo": "https://gravatar.com/avatar/gp781413e3bb44143bddf43589b03038?s=26&d=404",
		"linkedContactIds": [
			"38f608d5-9a60-4960-83c1-99d18f40c428"
		]
	}
]`))

	if ta.NoError(err) {
		ta.Len(ids, 1)
		ta.Len(users, 1)
		ta.Equal(ids[0], "8d49b010-cc6a-4f40-ace5-e86061c677ed")
		ta.Equal(users[0].ID, "8d49b010-cc6a-4f40-ace5-e86061c677ed")
	}

	ids, users, err = unmarshalArrayOfUsersOrIDs(json.RawMessage(`[
	"8d49b010-cc6a-4f40-ace5-e86061c677ed"
]`))

	if ta.NoError(err) {
		ta.Len(ids, 1)
		ta.Empty(users)
		ta.Equal(ids[0], "8d49b010-cc6a-4f40-ace5-e86061c677ed")
	}
}

func TestInvalidUserArrayUnmarshalling(t *testing.T) {
	ta := assert.New(t)

	ids, users, err := unmarshalArrayOfUsersOrIDs(json.RawMessage(`[[]]`))
	if ta.Error(err) {
		ta.Empty(ids)
		ta.Nil(users)
	}
}
