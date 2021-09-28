package userAdmin

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

// FindUser searches for a one or more users that are matched against the
// supplied filter.
func FindUser(filter string) (ok bool, user DatadogUser, err error) {
	ctx := datadog.NewDefaultContext(context.Background())
	opt := datadog.ListUsersOptionalParameters{Filter: &filter}
	configuration := datadog.NewConfiguration()
	client := datadog.NewAPIClient(configuration)
	r, _, err := client.UsersApi.ListUsers(ctx, opt)
	if err != nil {
		return
	}
	if *r.Data == nil || len(*r.Data) == 0 {
		return
	}

	ok = true
	user = NewUser(r.GetData()[0])
	return
}
