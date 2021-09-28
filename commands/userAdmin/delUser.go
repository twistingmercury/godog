package userAdmin

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func DisableUser(email string) (user datadog.User, err error) {
	user, err = FindUserByEmail(email)
	if err != nil {
		return
	}
	ctx := datadog.NewDefaultContext(context.Background())
	attribs := *datadog.NewUserUpdateAttributes()
	attribs.SetDisabled(true)
	data := *datadog.NewUserUpdateData(attribs, *user.Id, datadog.UsersType("users"))
	body := *datadog.NewUserUpdateRequest(data)

	cfg := datadog.NewConfiguration()
	client := datadog.NewAPIClient(cfg)

	r, _, err := client.UsersApi.UpdateUser(ctx, *user.Id, body)
	if err != nil {
		return
	}
	user = *r.Data
	return
}
