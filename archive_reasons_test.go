package lever

import (
	"context"
	"net/http"
	"testing"

	"github.com/corbaltcode/lever-data-api-go/internal/testclient"
	"github.com/corbaltcode/lever-data-api-go/model"
	"github.com/stretchr/testify/assert"
)

func TestListGetArchiveReasons(t *testing.T) {
	s := testclient.NewExpectManyHandler(
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"63dd55b2-a99f-4e7b-985f-22c7bf80ab42","text":"Underqualified","status":"active","type":"non-hired"}],"hasNext":true,"next":"1"}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/archive_reasons"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"41f98875-06c7-4cb1-b3d0-07f7ae192c0c","text":"Culture Fit","status":"active","type":"non-hired"}],"hasNext":true,"next":"2"}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/archive_reasons"),
			testclient.ExpectQuery("offset", "1"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"737b8899-7f32-42a5-954f-e199b0306fcb","text":"Timing","status":"active","type":"non-hired"}],"hasNext":true,"next":"3"}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/archive_reasons"),
			testclient.ExpectQuery("offset", "2"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"3274b963-c37c-4465-abeb-1113896e2aa3","text":"Withdrew","status":"active","type":"non-hired"}],"hasNext":true,"next":"4"}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/archive_reasons"),
			testclient.ExpectQuery("offset", "3"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"1fdbfaac-4c73-45f7-af32-369054426364","text":"Offer declined","status":"active","type":"non-hired"}],"hasNext":true,"next":"5"}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/archive_reasons"),
			testclient.ExpectQuery("offset", "4"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"16308515-dc77-4aa7-b3af-96161b2b4e9b","text":"Hired","status":"active","type":"hired"}],"hasNext":true,"next":"6"}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/archive_reasons"),
			testclient.ExpectQuery("offset", "5"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"9b6f482e-0399-4ba4-9a2f-1e1d9e3da15b","text":"Position filled","status":"inactive","type":"non-hired"}],"hasNext":false}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/archive_reasons"),
			testclient.ExpectQuery("offset", "6"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":{"id":"63dd55b2-a99f-4e7b-985f-22c7bf80ab42","text":"Underqualified","status":"active","type":"non-hired"}}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/archive_reasons/63dd55b2-a99f-4e7b-985f-22c7bf80ab42"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":{"id":"63dd55b2-a99f-4e7b-985f-22c7bf80ab42","text":"Underqualified"}}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/archive_reasons/63dd55b2-a99f-4e7b-985f-22c7bf80ab42"),
			testclient.ExpectQuery("include", "text"),
		),
	)

	httpClient := http.Client{
		Transport: s,
	}

	c := NewClient(WithHTTPClient(&httpClient))
	ctx := context.Background()

	// List all archive reasons, but limit the response to 1 item.
	var archiveReasons []model.ArchiveReason

	listReq := NewListArchiveReasonsRequest()
	listReq.Limit = 1

	ta := assert.New(t)

	for {
		listResp, err := c.ListArchiveReasons(ctx, listReq)
		if !ta.NoError(err) {
			t.Fatal("Failed to list archive reasons:", err)
		}

		archiveReasons = append(archiveReasons, listResp.ArchiveReasons...)
		if !listResp.HasNext {
			break
		}

		listReq.Offset = listResp.Next
	}

	if !ta.NotEmpty(archiveReasons) {
		t.Fatal("No archive reasons found")
	}

	// Get the first archive reason and make sure it matches the first archive reason in the list.
	reason := archiveReasons[0]
	getReq := NewGetArchiveReasonRequest(reason.ID)
	getResp, err := c.GetArchiveReason(ctx, getReq)

	if ta.NoError(err) {
		if ta.NotNil(getResp.ArchiveReason) {
			ta.Equal(getResp.ArchiveReason.ID, reason.ID)
			ta.Equal(getResp.ArchiveReason.Text, reason.Text)
			ta.Equal(getResp.ArchiveReason.Status, reason.Status)
			ta.Equal(getResp.ArchiveReason.Type, reason.Type)
		}
	}

	// Get the first archive reason again, but only include the text field.
	getReq = NewGetArchiveReasonRequest(reason.ID)
	getReq.Include = []string{"text"}
	getResp, err = c.GetArchiveReason(ctx, getReq)

	if ta.NoError(err) {
		if ta.NotNil(getResp.ArchiveReason) {
			ta.Equal(getResp.ArchiveReason.ID, reason.ID)
			ta.Equal(getResp.ArchiveReason.Text, reason.Text)
			ta.Empty(getResp.ArchiveReason.Status)
			ta.Empty(getResp.ArchiveReason.Type)
		}
	}
}
