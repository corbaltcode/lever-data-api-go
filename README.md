# lever-data-api-go
Golang interface to the [Lever data API](https://data.lever.co/)

## Using the API

To use this library, first add the required modules to your Go project:
```sh
go mod add github.com/corbaltcode/lever-data-api-go
go mod add github.com/corbaltcode/lever-data-api-go/model
```

To create a client handle, use `NewClient` and specify any options required. You will almost
certainly need to specify `WithAPIKey`.

```go
const apiKey = "EXAMPLE_API_KEY"

c := lever.NewClient(lever.WithAPIKey(apiKey))
```

Options include:
- `WithAPIKey`: Specify the API key to use in calls to the Lever API.
- `WithBaseURL`: Override the default base URL for the Lever API (default: `https://api.lever.co/v1`).
- `WithHeader`: Add headers to each request.
- `WithHTTPClient`: Use the specified HTTP client instead of creating a default.
- `WithUserAgent`: Override the default user agent (default: `lever-data-api-go/0.0.1`).

This library uses request and response objects for each API call. Required parameters are specified
in the `NewXxxRequest` function. Optional parameters can be set on the request before performing
the call.

```go
const userID = "526e1010-b7f8-48e1-aba8-ecce327775ca"

func testGetUser(c *lever.Client) error {
    getUserReq := lever.NewGetUserRequest(userId)
    getUserReq.Include = []string{"username", "accessRole"}
    
    getUserResp, err := c.GetUser(getUserReq)
    if err != nil {
        return err
    }

    fmt.Printf("Retrieved user: %v\n", getUserResp.User)
    return nil
}
```

All request objects inherit (via composition) from `lever.BaseRequest`, which includes the
following optional parameters:
```go
// Base type for all requests.
// This adds the includes and expands parameters.
type BaseRequest struct {
	// Parameters to include the the response. This is optional.
	Include []string

	// Parameters to expand in the response. This is optional.
	Expand []string
}
```

All responses give access to the original HTTP response object (though the body will have been
closed):
```go
type BaseResponse struct {
	HTTPResponse *http.Response `json:"-"`
}
```

### Pagination

List requests add two additional optional parameters from `lever.BaseListRequest`:
```go
// Base type for all list requests.
// This builds on BaseRequest by adding limit and offset parameters.
type BaseListRequest struct {
	BaseRequest

	// The number of items to include in the response. This is optional.
	Limit int

	// The pagination offset. This is optional.
	Offset string
}
```

All list responses include a `Next` and `HasNext` parameter for pagination:
```go
// Base type for all list responses.
// This adds the next and hasNext parameters for pagination.
type BaseListResponse struct {
	BaseResponse

	// The next pagination offset.
	Next string `json:"next,omitempty"`

	// Whether there is a next page.
	HasNext bool `json:"hasNext,omitempty"`
}
```

You may not be able to retrieve all items in a single List API call, so you will have to paginate
through the results by passing the `Next` field from the response into the `Offset` field for the
next request:

```go
func testListAllUsers(c *lever.Client) error {
    listReq := lever.NewListUsersRequest()
    for {
        listResp, err := c.ListUsers(listReq)
        if err != nil {
            return err
        }

        for _, user := range listResp.Users {
            fmt.Printf("Found user %v", user.Name)
        }

        if !listResp.HasNext {
            break
        }

        listReq.Offset = listResp.Next
    }

    return nil
}
```

## API support status

The following APIs have been implemented.

- [Applications](https://hire.lever.co/developer/documentation#applications)
- [Archive Reasons](https://hire.lever.co/developer/documentation#archive-reasons)
- [Contacts](https://hire.lever.co/developer/documentation#contacts)
- [Opportunities](https://hire.lever.co/developer/documentation#opportunities)
- [Sources](https://hire.lever.co/developer/documentation#sources)
- [Stages](https://hire.lever.co/developer/documentation#stages)
- [Tags](https://hire.lever.co/developer/documentation#tags)
- [Users](https://hire.lever.co/developer/documentation#users)


The following APIs are in progress.
- [Resumes](https://hire.lever.co/developer/documentation#resumes)

The following APIs are not yet implemented.

- [Audit Events](https://hire.lever.co/developer/documentation#audit-events)
- [EEO Questions](https://hire.lever.co/developer/documentation#eeo)
- [Feedback Forms](https://hire.lever.co/developer/documentation#feedback)
- [Feedback Templates](https://hire.lever.co/developer/documentation#feedback-templates)
- [Files](https://hire.lever.co/developer/documentation#files)
- [Form Fields](https://hire.lever.co/developer/documentation#form-fields)
- [Interviews](https://hire.lever.co/developer/documentation#interviews)
- [Notes](https://hire.lever.co/developer/documentation#notes)
- [Offers](https://hire.lever.co/developer/documentation#offers)
- [Panels](https://hire.lever.co/developer/documentation#panels)
- [Postings](https://hire.lever.co/developer/documentation#postings)
- [Posting Forms](https://hire.lever.co/developer/documentation#posting-forms)
- [Profile Forms](https://hire.lever.co/developer/documentation#profile-forms)
- [Profile Form Templates](https://hire.lever.co/developer/documentation#profile-form-templates)
- [Referrals](https://hire.lever.co/developer/documentation#referrals)
- [Requisitions](https://hire.lever.co/developer/documentation#requisitions)
- [Requisition Fields](https://hire.lever.co/developer/documentation#requisition-fields)
- [Uploads](https://hire.lever.co/developer/documentation#uploads)
- [Webhooks](https://hire.lever.co/developer/documentation#webhooks-via-the-api)

The following are deprecated APIs and will not be implemented.

- [Candidates](https://hire.lever.co/developer/documentation#candidates)