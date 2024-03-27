package lever

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/corbaltcode/lever-data-api-go/internal/testclient"
	"github.com/corbaltcode/lever-data-api-go/model"
	"github.com/stretchr/testify/assert"
)

var zeroTime int64 = 0
var endTime int64 = int64(9223372036854775807)

func TestOpportunities(t *testing.T) {
	ta := assert.New(t)

	candidates := []map[string]any{shaneSmith, chaofanWest, robertaEaston}

	s := testclient.NewExpectManyHandler(
		testclient.NewExpectHandler(
			http.StatusOK,
			toJSON(map[string]any{"data": expandCandidates(candidates)}),
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/opportunities"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			toJSON(map[string]any{"data": expandCandidates([]map[string]any{shaneSmith}, "followers", "owner", "sourcedBy", "stage")}),
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/opportunities"),
			testclient.ExpectQuery("expand", "followers"),
			testclient.ExpectQuery("expand", "owner"),
			testclient.ExpectQuery("expand", "sourcedBy"),
			testclient.ExpectQuery("expand", "stage"),
		),
		testclient.NewExpectHandler( // Make sure expanded JSON form works.
			http.StatusOK,
			toJSONIndent(map[string]any{"data": expandCandidates([]map[string]any{shaneSmith}, "followers", "owner", "sourcedBy", "stage")}),
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/opportunities"),
			testclient.ExpectQuery("expand", "followers"),
			testclient.ExpectQuery("expand", "owner"),
			testclient.ExpectQuery("expand", "sourcedBy"),
			testclient.ExpectQuery("expand", "stage"),
			testclient.ExpectQuery("tag", "San Francisco"),
			testclient.ExpectQuery("email", "shane@exampleq3.com"),
			testclient.ExpectQuery("origin", "sourced"),
			testclient.ExpectQuery("source", "linkedin"),
			testclient.ExpectQuery("confidentiality", "non-confidential"),
			testclient.ExpectQuery("stage_id", "00922a60-7c15-422b-b086-f62000824fd7"),
			testclient.ExpectQuery("posting_id", "cdb4ff13-f7aa-49b0-b6ec-eb4617009cfa"),
			testclient.ExpectQuery("archived_posting_id", "cdb4ff13-f7aa-49b0-b6ec-eb4617009cfa"),
			testclient.ExpectQuery("created_at_start", "0"),
			testclient.ExpectQuery("created_at_end", "9223372036854775807"),
			testclient.ExpectQuery("updated_at_start", "0"),
			testclient.ExpectQuery("updated_at_end", "9223372036854775807"),
			testclient.ExpectQuery("advanced_at_start", "0"),
			testclient.ExpectQuery("advanced_at_end", "9223372036854775807"),
			testclient.ExpectQuery("archived_at_start", "0"),
			testclient.ExpectQuery("archived_at_end", "9223372036854775807"),
			testclient.ExpectQuery("archived", "true"),
			testclient.ExpectQuery("archive_reason_id", "cdb4ff13-f7aa-49b0-b6ec-eb4617009cfa"),
			testclient.ExpectQuery("snoozed", "true"),
			testclient.ExpectQuery("contact_id", "cdb4ff13-f7aa-49b0-b6ec-eb4617009cfa"),
			testclient.ExpectQuery("location", "San Francisco"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			toJSON(map[string]any{"data": expandCandidates([]map[string]any{shaneSmith})}),
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/opportunities/deleted"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			toJSONIndent(map[string]any{"data": expandCandidate(shaneSmith, "followers", "owner", "sourcedBy", "stage")}),
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/opportunities/250d8f03-738a-4bba-a671-8a3d73477145"),
			testclient.ExpectQuery("expand", "followers"),
			testclient.ExpectQuery("expand", "owner"),
			testclient.ExpectQuery("expand", "sourcedBy"),
			testclient.ExpectQuery("expand", "stage"),
		),
	)

	httpClient := http.Client{
		Transport: s,
	}

	c := NewClient(WithHTTPClient(&httpClient))
	ctx := context.Background()

	// List opportunities
	listReq := NewListOpportunitiesRequest()
	listResp, err := c.ListOpportunities(ctx, listReq)

	if ta.NoError(err) {
		ta.Len(listResp.Opportunities, 3)
		for _, opportunity := range listResp.Opportunities {
			ta.NotEmpty(opportunity.ID)
			ta.NotEmpty(opportunity.FollowerIDs)
			ta.Empty(opportunity.Followers)
		}
	}

	// Expand all fields
	listReq.Expand = []string{"followers", "owner", "sourcedBy", "stage"}
	listReq.Limit = 1
	listResp, err = c.ListOpportunities(ctx, listReq)

	if ta.NoError(err) {
		ta.Len(listResp.Opportunities, 1)
		opportunity := listResp.Opportunities[0]
		ta.NotEmpty(opportunity.ID)
		ta.NotEmpty(opportunity.FollowerIDs)
		ta.Len(opportunity.Followers, 3)

		ta.NotNil(opportunity.Owner)
		ta.Equal(opportunity.OwnerID, opportunity.Owner.ID)

		ta.NotNil(opportunity.SourcedBy)
		ta.Equal(opportunity.SourcedByID, opportunity.SourcedBy.ID)

		ta.NotNil(opportunity.Stage)
		ta.Equal(opportunity.StageID, opportunity.Stage.ID)
	}

	var isArchived bool = true

	// Make sure expansion works with JSON whitespace; add filters
	listReq.Tags = []string{"San Francisco"}
	listReq.Emails = []string{"shane@exampleq3.com"}
	listReq.Origins = []string{"sourced"}
	listReq.Sources = []string{"linkedin"}
	listReq.Confidentiality = []string{"non-confidential"}
	listReq.StageIDs = []string{"00922a60-7c15-422b-b086-f62000824fd7"}
	listReq.PostingIDs = []string{"cdb4ff13-f7aa-49b0-b6ec-eb4617009cfa"}
	listReq.ArchivedPostingIDs = []string{"cdb4ff13-f7aa-49b0-b6ec-eb4617009cfa"}
	listReq.CreatedAtStart = &zeroTime
	listReq.CreatedAtEnd = &endTime
	listReq.UpdatedAtStart = &zeroTime
	listReq.UpdatedAtEnd = &endTime
	listReq.AdvancedAtStart = &zeroTime
	listReq.AdvancedAtEnd = &endTime
	listReq.ArchivedAtStart = &zeroTime
	listReq.ArchivedAtEnd = &endTime
	listReq.Archived = &isArchived
	listReq.ArchiveReasonIDs = []string{"cdb4ff13-f7aa-49b0-b6ec-eb4617009cfa"}
	listReq.Snoozed = &isArchived
	listReq.ContactIDs = []string{"cdb4ff13-f7aa-49b0-b6ec-eb4617009cfa"}
	listReq.Locations = []string{"San Francisco"}
	listResp, err = c.ListOpportunities(ctx, listReq)

	if ta.NoError(err) {
		ta.Len(listResp.Opportunities, 1)
		opportunity := listResp.Opportunities[0]
		ta.NotEmpty(opportunity.ID)
		ta.NotEmpty(opportunity.FollowerIDs)
		ta.Len(opportunity.Followers, 3)

		ta.NotNil(opportunity.Owner)
		ta.Equal(opportunity.OwnerID, opportunity.Owner.ID)

		ta.NotNil(opportunity.SourcedBy)
		ta.Equal(opportunity.SourcedByID, opportunity.SourcedBy.ID)

		ta.NotNil(opportunity.Stage)
		ta.Equal(opportunity.StageID, opportunity.Stage.ID)
	}

	// List deleted opportunities
	listDelReq := NewListDeletedOpportunitiesRequest()
	listDelResp, err := c.ListDeletedOpportunities(ctx, listDelReq)

	if ta.NoError(err) {
		ta.Len(listDelResp.Opportunities, 1)
		opportunity := listDelResp.Opportunities[0]
		ta.NotEmpty(opportunity.ID)
		ta.NotEmpty(opportunity.FollowerIDs)
		ta.Empty(opportunity.Followers)
	}

	// Get an opportunity
	getReq := NewGetOpportunityRequest("250d8f03-738a-4bba-a671-8a3d73477145")
	getReq.Expand = []string{"stage", "sourcedBy", "owner"}
	getResp, err := c.GetOpportunity(ctx, getReq)

	if ta.NoError(err) {
		if ta.NotNil(getResp.Opportunity) {
			opportunity := getResp.Opportunity
			ta.Equal("250d8f03-738a-4bba-a671-8a3d73477145", opportunity.ID)
			ta.NotEmpty(opportunity.FollowerIDs)

			ta.NotNil(opportunity.Stage)
			ta.NotNil(opportunity.Owner)
			ta.NotNil(opportunity.SourcedBy)
		}
	}
}

// TestCreateCandidate tests creating a candidate.
func TestCreateCandidate(t *testing.T) {
	ta := assert.New(t)

	s := testclient.NewExpectManyHandler(
		testclient.NewExpectHandler(
			http.StatusOK,
			`{
	"deduped": false,
	"data": {
		"id": "cb45668d-38b6-43dd-9ba5-bd325d50dbfc",
		"name": "Dave Test Candidate DONOTUSE",
		"contact": "caa42efa-d432-4d22-bbcc-4c8e3751bb3e",
		"headline": "Brickly LLC, Vandelay Industries, Inc, Central Perk",
		"stage": "applicant-new",
		"confidentiality": "non-confidential",
		"location": "Seattle",
		"phones": [
			{
				"type": "other",
				"value": "(123) 456-7890"
			},
			{
				"type": "other",
				"value": "(234) 567-8901"
			}
		],
		"emails": [
			"testuser@example.com"
		],
		"links": [
			"http://indeed.com/r/Shane-Smith/0b7c87f6b246d2bc",
			"http://github.com/shanesmith"
		],
		"archived": null,
		"tags": [],
		"sources": [],
		"stageChanges": [
			{
				"toStageId": "applicant-new",
				"toStageIndex": 4,
				"updatedAt": 1711572497188,
				"userId": "68d28c32-e972-4f0f-81c7-83f2e9430293"
			}
		],
		"origin": "internal",
		"sourcedBy": null,
		"owner": "68d28c32-e972-4f0f-81c7-83f2e9430293",
		"followers": [
			"68d28c32-e972-4f0f-81c7-83f2e9430293",
			"113fa8ae-f6ea-4316-8267-7ed692fc9471"
		],
		"applications": [],
		"createdAt": 1711572497188,
		"updatedAt": 1711572497191,
		"lastInteractionAt": 1711572497189,
		"lastAdvancedAt": 1711572497188,
		"snoozedUntil": null,
		"urls": {
			"list": "https://hire.lever.co/candidates",
			"show": "https://hire.lever.co/candidates/cb45668d-38b6-43dd-9ba5-bd325d50dbfc"
		},
		"isAnonymized": false,
		"dataProtection": null
	}
}`,
			testclient.ExpectMethod(http.MethodPost),
			testclient.ExpectPath("/v1/opportunities"),
		),
	)

	httpClient := http.Client{
		Transport: s,
	}

	c := NewClient(WithHTTPClient(&httpClient))
	ctx := context.Background()

	createReq := NewCreateOpportunityRequest("68d28c32-e972-4f0f-81c7-83f2e9430293")
	createReq.Name = "Dave Test Candidate DONOTUSE"
	createReq.Headline = "Brickly LLC, Vandelay Industries, Inc, Central Perk"
	createReq.StageID = "cd4ff13-f7aa-49b0-b6ec-eb4617009cfa"
	createReq.Location = "Seattle"
	createReq.Phones = []model.Phone{
		{
			Type:  "mobile",
			Value: "(123) 456-7890",
		},
		{
			Type:  "mobile",
			Value: "(234) 567-8901",
		},
	}
	createReq.Emails = []string{"testcandidate@example.com"}
	createReq.Links = []string{"indeed.com/r/Shane-Smith/0b7c87f6b246d2bc", "github.com/shanesmith"}
	createReq.Sources = []string{"linkedin"}
	createReq.Origin = "internal"
	createReq.PostingID = "cdb4ff13-f7aa-49b0-b6ec-eb4617009cfa"
	createReq.CreatedAt = &zeroTime
	createReq.Archived = &model.Archived{
		ArchivedAt: &zeroTime,
		ReasonID:   "cdb4ff13-f7aa-49b0-b6ec-eb4617009cfa",
	}
	createReq.ContactID = "caa42efa-d432-4d22-bbcc-4c8e3751bb3e"

	createResp, err := c.CreateOpportunity(ctx, createReq)
	if ta.NoError(err) {
		ta.False(createResp.Deduped)

		if ta.NotNil(createResp.Opportunity) {
			ta.Equal(createResp.Opportunity.ID, "cb45668d-38b6-43dd-9ba5-bd325d50dbfc")
			ta.Equal(createResp.Opportunity.Name, "Dave Test Candidate DONOTUSE")
		}
	}

}

// expandCandidates expands the specified fields in the array of candidate data.
func expandCandidates(orig []map[string]any, fields ...string) []map[string]any {
	expanded := make([]map[string]any, 0, len(orig))
	for _, candidate := range orig {
		expanded = append(expanded, expandCandidate(candidate, fields...))
	}

	return expanded
}

// expandCandidate expands the specified fields in the candidate data.
func expandCandidate(orig map[string]any, fields ...string) map[string]any {
	expanded := make(map[string]any)
	for k, v := range orig {
		expanded[k] = v
	}

	for _, field := range fields {
		switch field {
		case "followers":
			expanded["followers"] = expandUsers(expanded["followers"].([]string))

		case "owner":
			expanded["owner"] = expandUser(expanded["owner"].(string))

		case "sourcedBy":
			sourcedBy := expanded["sourcedBy"]
			if sourcedBy != nil {
				expanded["sourcedBy"] = expandUser(sourcedBy.(string))
			}

		case "stage":
			expanded["stage"] = expandStage(expanded["stage"].(string))
		}
	}

	return expanded
}

// toJSON converts the given value to a compact JSON string, panicking on error.
func toJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(b)
}

// toJSONIndent converts the given value to an indented JSON string, panicking on error.
func toJSONIndent(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

// expandStage expands the stage ID into stage data.
func expandStage(stageID string) map[string]any {
	stage, ok := stageIDToStage[stageID]
	if ok {
		return stage
	}

	panic(fmt.Sprintf("No stage found for ID %s", stageID))
}

// expandUsers expands the user IDs into user data.
func expandUsers(userIDs []string) []map[string]any {
	result := make([]map[string]any, 0, len(userIDs))

	for _, userID := range userIDs {
		result = append(result, expandUser(userID))
	}

	return result
}

// expandUser expands the user ID into user data.
func expandUser(userID string) map[string]any {
	user, ok := userIDToUser[userID]
	if ok {
		return user
	}

	panic(fmt.Sprintf("No user found for ID %s", userID))
}

var shaneSmith = map[string]any{
	"id":       "250d8f03-738a-4bba-a671-8a3d73477145",
	"name":     "Shane Smith",
	"headline": "Brickly LLC, Vandelay Industries, Inc, Central Perk",
	"contact":  "7f23e772-f2cb-4ebb-b33f-54b872999992",
	"emails":   []string{"shane@exampleq3.com"},
	"phones": []map[string]any{
		{
			"value": "(123) 456-7891",
		},
	},
	"confidentiality":   "non-confidential",
	"location":          "Oakland",
	"links":             []string{"indeed.com/r/Shane-Smith/0b7c87f6b246d2bc"},
	"createdAt":         1407460071043,
	"updatedAt":         1407460080914,
	"lastInteractionAt": 1417588008760,
	"lastAdvancedAt":    1417587916150,
	"snoozedUntil":      1505971500000,
	"archivedAt":        nil,
	"archiveReason":     nil,
	"stage":             "00922a60-7c15-422b-b086-f62000824fd7",
	"stageChanges": []map[string]any{
		{
			"toStageId":    "00922a60-7c15-422b-b086-f62000824fd7",
			"toStageIndex": 1,
			"userId":       "df0adaa6-172c-4cd6-8520-49b203660fe1",
			"updatedAt":    1407460071043,
		},
	},
	"owner": "df0adaa6-172c-4cd6-8520-49b203660fe1",
	"tags": []string{
		"San Francisco",
		"Full-time",
		"Support",
		"Customer Success",
		"Customer Success Manager",
	},
	"sources":      []string{"linkedin"},
	"origin":       "sourced",
	"sourcedBy":    "df0adaa6-172c-4cd6-8520-49b203660fe1",
	"applications": []string{"cdb4ff13-f7aa-49b0-b6ec-eb4617009cfa"},
	"resume":       nil,
	"followers": []string{
		"df0adaa6-172c-4cd6-8520-49b203660fe1",
		"ecdb6670-d9f3-4b87-8267-1cde26d1bc42",
		"022d6639-1333-419b-9635-31f93015335f",
	},
	"urls": map[string]any{
		"list": "https://hire.lever.co/candidates",
		"show": "https://hire.lever.co/candidates/250d8f03-738a-4bba-a671-8a3d73477145",
	},
	"dataProtection": map[string]any{
		"store": map[string]any{
			"allowed":   true,
			"expiresAt": 1522540800000,
		},
		"contact": map[string]any{
			"allowed":   false,
			"expiresAt": nil,
		},
	},
	"isAnonymized": false,
}

var chaofanWest = map[string]any{
	"id":       "5c86dcd8-6cf1-40da-9ae3-5e7ea91079f5",
	"name":     "Chaofan West",
	"headline": "Grunnings, Inc., Coffee Bean, Ltd, Betelgeuse Commercial Service Co., Ltd, Double C Private Co., Ltd",
	"contact":  "bd4d81c8-7858-4624-be98-552dfb9ca850",
	"emails":   []string{"chaofan@example.com"},
	"phones": []map[string]any{
		{
			"value": "(123) 456-7891",
		},
	},
	"location":          "San Francisco",
	"links":             []string{"indeed.com/r/Chaofan-West/4f2c7523b0edefbb"},
	"createdAt":         1407778275799,
	"lastInteractionAt": 1417587990376,
	"lastAdvancedAt":    1417587903121,
	"snoozedUntil":      1499577840000,
	"archivedAt":        nil,
	"archiveReason":     nil,
	"stage":             "00922a60-7c15-422b-b086-f62000824fd7",
	"owner":             "ecdb6670-d9f3-4b87-8267-1cde26d1bc42",
	"tags": []string{
		"San Francisco",
		"Marketing",
		"Customer Success",
		"Customer Success Manager",
		"Full-time",
	},
	"sources":      []string{"Job site"},
	"origin":       "applied",
	"sourcedBy":    nil,
	"applications": []string{"e326d6e6-e3f6-46eb-9c14-6f90b88aacac"},
	"resume":       nil,
	"followers": []string{
		"df0adaa6-172c-4cd6-8520-49b203660fe1",
		"ecdb6670-d9f3-4b87-8267-1cde26d1bc42",
		"022d6639-1333-419b-9635-31f93015335f",
	},
	"dataProtection": nil,
}

var robertaEaston = map[string]any{
	"id":       "37caee03-bd3f-487d-b32e-a296ce05aa6b",
	"name":     "Roberta Easton",
	"headline": "Useful Information Access",
	"contact":  "853af6b1-71a2-46d7-a430-c067e28b08f9",
	"emails":   []string{"roberta@exampleq3.com"},
	"phones": []map[string]any{
		{
			"value": "(123) 456-7891",
		},
	},
	"location":          "San Jose",
	"links":             []string{"https://linkedin.com/in/roberta-e"},
	"createdAt":         1407778277088,
	"lastInteractionAt": 1417587981210,
	"lastAdvancedAt":    1417587891220,
	"snoozedUntil":      1420266291216,
	"archivedAt":        nil,
	"archiveReason":     nil,
	"stage":             "00922a60-7c15-422b-b086-f62000824fd7",
	"owner":             "ecdb6670-d9f3-4b87-8267-1cde26d1bc42",
	"tags": []string{
		"San Francisco",
		"Marketing",
		"Customer Success",
		"Customer Success Manager",
		"Full-time",
	},
	"sources":      []string{"Job site"},
	"origin":       "applied",
	"sourcedBy":    nil,
	"applications": []string{"eb91c63f-3511-4e9d-a805-aaa92f0c80c9"},
	"resume":       nil,
	"followers": []string{
		"df0adaa6-172c-4cd6-8520-49b203660fe1",
		"ecdb6670-d9f3-4b87-8267-1cde26d1bc42",
		"022d6639-1333-419b-9635-31f93015335f",
	},
	"dataProtection": nil,
}

var chandlerBing = map[string]any{
	"id":                  "df0adaa6-172c-4cd6-8520-49b203660fe1",
	"name":                "Chandler Bing",
	"username":            "chandler",
	"email":               "chandler@example.com",
	"createdAt":           1407357447018,
	"deactivatedAt":       nil,
	"externalDirectoryId": "2277399",
	"accessRole":          "super admin",
	"photo":               "https://gravatar.com/avatar/gp781413e3bb44143bddf43589b03038?s=26&d=404",
	"linkedContactIds":    []string{"38f608d5-9a60-4960-83c1-99d18f40c428"},
}

var rachelGreen = map[string]any{
	"id":                  "ecdb6670-d9f3-4b87-8267-1cde26d1bc42",
	"name":                "Rachel Green",
	"username":            "rachel",
	"email":               "rachel@example.com",
	"createdAt":           1407357447018,
	"deactivatedAt":       nil,
	"externalDirectoryId": "2277400",
	"accessRole":          "interviewer",
	"photo":               "https://gravatar.com/avatar/gp781413e3bb44143bddf43589b03038?s=26&d=404",
	"linkedContactIds":    []string{"38f608d5-9a60-4960-83c1-99d18f40c428"},
}

var monicaGeller = map[string]any{
	"id":                  "022d6639-1333-419b-9635-31f93015335f",
	"name":                "Monica Geller",
	"username":            "monica",
	"email":               "monica@example.com",
	"createdAt":           1407357447018,
	"deactivatedAt":       nil,
	"externalDirectoryId": "2277401",
	"accessRole":          "admin",
	"photo":               "https://gravatar.com/avatar/gp781413e3bb44143bddf43589b03038?s=26&d=404",
	"linkedContactIds":    []string{"38f608d5-9a60-4960-83c1-99d18f40c428"},
}

var userIDToUser = map[string]map[string]any{
	"df0adaa6-172c-4cd6-8520-49b203660fe1": chandlerBing,
	"ecdb6670-d9f3-4b87-8267-1cde26d1bc42": rachelGreen,
	"022d6639-1333-419b-9635-31f93015335f": monicaGeller,
}

var stageRecruiterScreen = map[string]any{
	"id":   "00922a60-7c15-422b-b086-f62000824fd7",
	"text": "Recruiter Screen",
}

var stageIDToStage = map[string]map[string]any{
	"00922a60-7c15-422b-b086-f62000824fd7": stageRecruiterScreen,
}
