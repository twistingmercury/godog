package userAdmin

import (
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/twistingmercury/godog/commands/utils"
)

func FindUserByEmail(email string) (user *DatadogUser, err error) {
	ctx := utils.NewContext()
	opt := datadog.ListUsersOptionalParameters{Filter: &email}

	configuration := datadog.NewConfiguration()

	client := datadog.NewAPIClient(configuration)
	r, _, err := client.UsersApi.ListUsers(ctx, opt)
	if err != nil {
		return
	}

	data := r.GetData()

	if len(data) > 0 {
		user = &DatadogUser{
			Name:   *data[0].Attributes.Name,
			Email:  *data[0].Attributes.Email,
			ID:     *data[0].Id,
			Status: *data[0].Attributes.Status,
		}
	}

	return
}
