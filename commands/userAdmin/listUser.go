package userAdmin

import (
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/twistingmercury/godog/commands/utils"
)

func ListAllUsers() (users []DatadogUser, err error) {
	ctx := utils.NewContext()
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

	data := r.GetData()

	for _, u := range data {
		users = append(users, DatadogUser{
			Name:   *u.Attributes.Name,
			Email:  *u.Attributes.Email,
			ID:     *u.Id,
			Status: *u.Attributes.Status})
	}
	return
}
