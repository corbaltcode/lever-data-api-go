package lever

import (
	"context"
	"net/http"
	"testing"

	"github.com/corbaltcode/lever-data-api-go/internal/testclient"
	"github.com/corbaltcode/lever-data-api-go/model"
	"github.com/stretchr/testify/assert"
)

func TestStages(t *testing.T) {
	s := testclient.NewExpectManyHandler(
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"fff60592-31dd-4ebe-ba8e-e7a397c30f8e","text":"New applicant"},{"id":"51adb2bb-1e24-4135-9950-cb96e3886226","text":"New lead"},{"id":"a42482ff-00ac-4cb2-a698-4b4436885b0c","text":"Recruiter Screen"},{"id":"fe763d80-e612-4787-98bc-686679c6ac9b","text":"Phone Interview"},{"id":"fbfb4473-38d2-4acf-943b-28cc0ed7ba87","text":"On-Site Interview"},{"id":"e7c6f8eb-9239-46b8-8777-ef4215cc8a7d","text":"Background Check"},{"id":"f48aad6f-91f7-4e87-b3f3-5e7a9207e54b","text":"Offer"}],"hasNext":false}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/stages"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":{"id":"fff60592-31dd-4ebe-ba8e-e7a397c30f8e","text":"New applicant"}}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/stages/fff60592-31dd-4ebe-ba8e-e7a397c30f8e"),
		),
		testclient.NewExpectHandler(
			http.StatusNotFound,
			`{"code":"ResourceNotFound","message":"Stage was not found"}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/stages/00000000-0000-0000-0000-000000000000"),
		),
	)

	httpClient := http.Client{
		Transport: s,
	}

	ta := assert.New(t)

	c := NewClient(WithHTTPClient(&httpClient))
	ctx := context.Background()
	var leverError *model.LeverError

	// List stages
	listReq := NewListStagesRequest()
	listResp, err := c.ListStages(ctx, listReq)

	if ta.NoError(err) {
		ta.Len(listResp.Stages, 7)
	}

	// Get a single stage
	getReq := NewGetStageRequest("fff60592-31dd-4ebe-ba8e-e7a397c30f8e")
	getResp, err := c.GetStage(ctx, getReq)

	if ta.NoErrorf(err, "Failed to get stage: %v", err) {
		if ta.NotNil(getResp.Stage, "Expected a stage") {
			ta.Equal(getResp.Stage.Text, "New applicant")
		}
	}

	// Get a single stage that does not exist
	getReq = NewGetStageRequest("00000000-0000-0000-0000-000000000000")
	getResp, err = c.GetStage(ctx, getReq)

	if ta.Error(err) {
		ta.Nil(getResp)
		if ta.ErrorAs(err, &leverError) {
			ta.Equal("ResourceNotFound", leverError.Code)
			ta.Equal("Stage was not found", leverError.Message)
			if ta.NotNil(leverError.HTTPResponse) {
				ta.Equal(http.StatusNotFound, leverError.HTTPResponse.StatusCode)
			}
		}
	}
}
