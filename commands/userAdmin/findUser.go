package userAdmin

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func FindUserByEmail(email string) (user datadog.User, err error) {
	ctx := datadog.NewDefaultContext(context.Background())

	opt := datadog.ListUsersOptionalParameters{Filter: &email}

	configuration := datadog.NewConfiguration()

	client := datadog.NewAPIClient(configuration)
	r, _, err := client.UsersApi.ListUsers(ctx, opt)
	if err != nil {
		return
	}

	user = r.GetData()[0]

	return
}
