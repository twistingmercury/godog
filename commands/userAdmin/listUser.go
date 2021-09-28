package userAdmin

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

// ListAllUsers invokes the DD api to list out all active and pending users.
func ListAllUsers() (user []datadog.User, err error) {
	ctx := datadog.NewDefaultContext(context.Background())
	status := "Active,Pending"
	pageSize := int64(500)
	filter := "users"
	opt := datadog.ListUsersOptionalParameters{FilterStatus: &status, PageSize: &pageSize, Filter: &filter}

	configuration := datadog.NewConfiguration()

	client := datadog.NewAPIClient(configuration)
	r, _, err := client.UsersApi.ListUsers(ctx, opt)
	if err != nil {
		return
	}

	user = r.GetData()

	return
}
