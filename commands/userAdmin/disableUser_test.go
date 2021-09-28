package userAdmin

import (
	"reflect"
	"testing"
)

func TestDisableUser(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name     string
		args     args
		wantOk   bool
		wantUser DatadogUser
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotUser, err := DisableUser(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("DisableUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("DisableUser() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("DisableUser() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
