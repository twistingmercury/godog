package userAdmin

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

// DisableUser invokes the DD api to disable an active or pending user account.
func DisableUser(email string) (ok bool, user DatadogUser, err error) {
	ok, user, err = FindUser(email)
	if err != nil {
		return
	}
	ctx := datadog.NewDefaultContext(context.Background())
	attribs := *datadog.NewUserUpdateAttributes()
	attribs.SetDisabled(true)
	data := *datadog.NewUserUpdateData(attribs, user.Id, datadog.UsersType("users"))
	body := *datadog.NewUserUpdateRequest(data)

	cfg := datadog.NewConfiguration()
	client := datadog.NewAPIClient(cfg)

	r, _, err := client.UsersApi.UpdateUser(ctx, user.Id, body)
	if err != nil {
		return
	}
	user = NewUser(*r.Data)
	return
}
