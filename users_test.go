package lever

import (
	"context"
	"net/http"
	"testing"

	"github.com/corbaltcode/lever-data-api-go/internal/testclient"
	"github.com/corbaltcode/lever-data-api-go/model"
	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {
	s := testclient.NewExpectManyHandler(
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"8d49b010-cc6a-4f40-ace5-e86061c677ed","name":"Chandler Bing","username": "chandler","email":"chandler@example.com","createdAt":1407357447018,"deactivatedAt":1409556487918,"externalDirectoryId":"2277399","accessRole":"super admin","photo":"https://gravatar.com/avatar/gp781413e3bb44143bddf43589b03038?s=26&d=404","linkedContactIds":["38f608d5-9a60-4960-83c1-99d18f40c428"]}],"hasNext":false}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/users"),
		),
		testclient.NewExpectHandler(
			http.StatusBadRequest,
			`{"code":"BadRequestError","message":"A user already exists with that email address. at /email"}`,
			testclient.ExpectMethod(http.MethodPost),
			testclient.ExpectPath("/v1/users"),
		),
		testclient.NewExpectHandler(
			http.StatusBadRequest,
			`{"code":"BadRequestError","message":"Email domain must be one of: example.com at /email"}`,
			testclient.ExpectMethod(http.MethodPost),
			testclient.ExpectPath("/v1/users"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":{"id":"00d29867-34e2-4100-99a0-2f9f10f9b93d","name":"Test User1","username":"testuser","email":"testuser@example.com","accessRole":"interviewer","photo":null,"createdAt":1711510651144,"deactivatedAt":null,"externalDirectoryId":null,"linkedContactIds":null,"jobTitle":null,"managerId":null}}`,
			testclient.ExpectMethod(http.MethodPost),
			testclient.ExpectPath("/v1/users"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":{"id":"00d29867-34e2-4100-99a0-2f9f10f9b93d","name":"Test User1","username":"testuser","email":"testuser@example.com","accessRole":"interviewer","photo":null,"createdAt":1711510651144,"deactivatedAt":null,"externalDirectoryId":null,"linkedContactIds":null,"jobTitle":null,"managerId":null}}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/users/00d29867-34e2-4100-99a0-2f9f10f9b93d"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":{"id":"00d29867-34e2-4100-99a0-2f9f10f9b93d","name":"Test User1","username":"testuser","email":"testuser@example.com","accessRole":"interviewer","photo":null,"createdAt":1711510651144,"deactivatedAt":1711510789761,"externalDirectoryId":null,"linkedContactIds":null,"jobTitle":null,"managerId":null}}`,
			testclient.ExpectMethod(http.MethodPost),
			testclient.ExpectPath("/v1/users/00d29867-34e2-4100-99a0-2f9f10f9b93d/deactivate"),
		),
		testclient.NewExpectHandler(
			http.StatusNotFound,
			`{"code":"ResourceNotFound","message":"user was not found"}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/users/00d29867-34e2-4100-99a0-2f9f10f9b93d"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"8d49b010-cc6a-4f40-ace5-e86061c677ed","name":"Chandler Bing","username": "chandler","email":"chandler@example.com","createdAt":1407357447018,"deactivatedAt":1409556487918,"externalDirectoryId":"2277399","accessRole":"super admin","photo":"https://gravatar.com/avatar/gp781413e3bb44143bddf43589b03038?s=26&d=404","linkedContactIds":["38f608d5-9a60-4960-83c1-99d18f40c428"]}],"hasNext":false}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/users"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"8d49b010-cc6a-4f40-ace5-e86061c677ed","name":"Chandler Bing","username": "chandler","email":"chandler@example.com","createdAt":1407357447018,"deactivatedAt":1409556487918,"externalDirectoryId":"2277399","accessRole":"super admin","photo":"https://gravatar.com/avatar/gp781413e3bb44143bddf43589b03038?s=26&d=404","linkedContactIds":["38f608d5-9a60-4960-83c1-99d18f40c428"]},{"id":"00d29867-34e2-4100-99a0-2f9f10f9b93d","name":"Test User1","username":"testuser","email":"testuser@example.com","accessRole":"interviewer","photo":null,"createdAt":1711510651144,"deactivatedAt":1711510789761,"externalDirectoryId":null,"linkedContactIds":null,"jobTitle":null,"managerId":null}],"hasNext":false}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/users"),
			testclient.ExpectQuery("includeDeactivated", "true"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":{"id":"00d29867-34e2-4100-99a0-2f9f10f9b93d","name":"Test User1","username":"testuser","email":"testuser@example.com","accessRole":"interviewer","photo":null,"createdAt":1711510651144,"deactivatedAt":null,"externalDirectoryId":null,"linkedContactIds":null,"jobTitle":null,"managerId":null}}`,
			testclient.ExpectMethod(http.MethodPost),
			testclient.ExpectPath("/v1/users/00d29867-34e2-4100-99a0-2f9f10f9b93d/reactivate"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":{"id":"00d29867-34e2-4100-99a0-2f9f10f9b93d","name":"Test User1","username":"testuser","email":"testuser@example.com","accessRole":"interviewer","photo":null,"createdAt":1711510651144,"deactivatedAt":null,"externalDirectoryId":"testuser","linkedContactIds":null,"jobTitle":null,"managerId":null}}`,
			testclient.ExpectMethod(http.MethodPut),
			testclient.ExpectPath("/v1/users/00d29867-34e2-4100-99a0-2f9f10f9b93d"),
		),
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"id":"00d29867-34e2-4100-99a0-2f9f10f9b93d","name":"Test User1","username":"testuser","email":"testuser@example.com","accessRole":"interviewer","photo":null,"createdAt":1711510651144,"deactivatedAt":1711510789761,"externalDirectoryId":"testuser","linkedContactIds":null,"jobTitle":null,"managerId":null}],"hasNext":false}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/users"),
			testclient.ExpectQuery("accessRole", "interviewer"),
			testclient.ExpectQuery("email", "testuser@example.com"),
			testclient.ExpectQuery("external_directory_id", "testuser"),
		),
	)

	httpClient := http.Client{
		Transport: s,
	}

	ta := assert.New(t)

	c := NewClient(WithHTTPClient(&httpClient))
	ctx := context.Background()
	var leverError *model.LeverError

	// List users.
	listReq := NewListUsersRequest()
	listResp, err := c.ListUsers(ctx, listReq)

	if ta.NoError(err) {
		if ta.NotEmpty(listResp.Users) {
			ta.Len(listResp.Users, 1)
			ta.Equal(listResp.Users[0].ID, "8d49b010-cc6a-4f40-ace5-e86061c677ed")
		}
	}

	// Create a user with an existing email.
	createReq := NewCreateUserRequest("Test User1", "chandler@example.com")
	createResp, err := c.CreateUser(ctx, createReq)

	ta.Nil(createResp)
	if ta.Error(err) {
		if ta.ErrorAs(err, &leverError) {
			ta.Equal(leverError.Code, "BadRequestError")
			if ta.NotNil(leverError.HTTPResponse) {
				ta.Equal(leverError.HTTPResponse.StatusCode, http.StatusBadRequest)
			}
		}
	}

	// Create a user with an invalid email domain.
	createReq = NewCreateUserRequest("Test User1", "testuser@google.com")
	createResp, err = c.CreateUser(ctx, createReq)

	ta.Nil(createResp)
	if ta.Error(err) {
		if ta.ErrorAs(err, &leverError) {
			ta.Equal(leverError.Code, "BadRequestError")
			if ta.NotNil(leverError.HTTPResponse) {
				ta.Equal(leverError.HTTPResponse.StatusCode, http.StatusBadRequest)
			}
		}
	}

	// Create a new, valid user.
	createReq = NewCreateUserRequest("Test User1", "testuser@example.com")
	createReq.AccessRole = "interviewer"
	createResp, err = c.CreateUser(ctx, createReq)
	var userID string

	if ta.NoError(err) {
		if ta.NotNil(createResp) {
			userID = createResp.User.ID
			ta.Equal(createResp.User.Name, "Test User1")
			ta.Equal(createResp.User.Email, "testuser@example.com")
			ta.Equal(createResp.User.AccessRole, "interviewer")
			ta.NotNil(createResp.User.CreatedAt)
			ta.Nil(createResp.User.DeactivatedAt)
		}
	} else {
		t.Fatal("Failed to create user; cannot proceed with remaining tests")
	}

	// Get the user we just created.
	getReq := NewGetUserRequest(userID)
	getResp, err := c.GetUser(ctx, getReq)

	if ta.NoError(err) {
		if ta.NotNil(getResp.User) {
			ta.Equal(getResp.User.ID, userID)
			ta.Equal(getResp.User.Name, "Test User1")
			ta.Equal(getResp.User.Email, "testuser@example.com")
		}
	}

	// Deactivate the user.
	deactivateReq := NewDeactivateUserRequest(userID)
	deactivateResp, err := c.DeactivateUser(ctx, deactivateReq)

	if ta.NoError(err) {
		if ta.NotNil(deactivateResp) {
			ta.Equal(deactivateResp.User.ID, userID)
			ta.Equal(deactivateResp.User.Name, "Test User1")
			ta.Equal(deactivateResp.User.Email, "testuser@example.com")
			ta.NotNil(deactivateResp.User.DeactivatedAt)
		}
	} else {
		t.Fatal("Failed to deactivate user; cannot proceed with remaining tests")
	}

	// Attempt to get the deactivated user; this should fail.
	getReq = NewGetUserRequest(userID)
	getResp, err = c.GetUser(ctx, getReq)

	if ta.Nil(getResp) {
		if ta.Error(err) {
			if ta.ErrorAs(err, &leverError) {
				ta.Equal(leverError.Code, "ResourceNotFound")
				if ta.NotNil(leverError.HTTPResponse) {
					ta.Equal(leverError.HTTPResponse.StatusCode, http.StatusNotFound)
				}
			}
		}
	}

	// List all users, without including the deactivated users.
	listReq = NewListUsersRequest()
	listResp, err = c.ListUsers(ctx, listReq)

	if ta.NoError(err) {
		if ta.NotNil(listResp) {
			if ta.Len(listResp.Users, 1) {
				ta.Equal(listResp.Users[0].ID, "8d49b010-cc6a-4f40-ace5-e86061c677ed")
			}
		}
	}

	// List all users, including the deactivated users.
	listReq = NewListUsersRequest()
	listReq.IncludeDeactivated = true
	listResp, err = c.ListUsers(ctx, listReq)

	if ta.NoError(err) {
		if ta.NotNil(listResp) {
			ta.Len(listResp.Users, 2)
		}
	}

	// Reactivate the user.
	reactivateReq := NewReactivateUserRequest(userID)
	reactivateResp, err := c.ReactivateUser(ctx, reactivateReq)

	if ta.NoError(err) {
		if ta.NotNil(reactivateResp) {
			ta.Equal(reactivateResp.User.ID, userID)
			ta.Equal(reactivateResp.User.Name, "Test User1")
			ta.Equal(reactivateResp.User.Email, "testuser@example.com")
		}
	}

	// Update the user.
	updateReq := NewUpdateUserRequestFromUser(reactivateResp.User)
	updateReq.ExternalDirectoryID = "testuser"
	updateResp, err := c.UpdateUser(ctx, updateReq)

	if ta.NoError(err) && ta.NotNil(updateResp) {
		ta.Equal(updateResp.User.ID, userID)
		ta.Equal(updateResp.User.Name, "Test User1")
		ta.Equal(updateResp.User.Email, "testuser@example.com")
		ta.Equal(updateResp.User.ExternalDirectoryID, "testuser")
	} else {
		t.Fatal("Failed to update user; cannot proceed with remaining tests")
	}

	// Find this user.
	listReq = NewListUsersRequest()
	listReq.AccessRole = append(listReq.AccessRole, "interviewer")
	listReq.Email = append(listReq.Email, "testuser@example.com")
	listReq.ExternalDirectoryID = append(listReq.ExternalDirectoryID, "testuser")

	listResp, err = c.ListUsers(ctx, listReq)
	if ta.NoError(err) {
		if ta.NotNil(listResp) {
			ta.Len(listResp.Users, 1)
			ta.Equal(listResp.Users[0].ID, userID)
		}
	}
}
