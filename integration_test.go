package fsadapter

import (
	"github.com/casbin/casbin/v2"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCasbinFsAdapter_Integration_OS(t *testing.T) {
	tests := []struct {
		name           string
		sub            string
		obj            string
		act            string
		expectedResult bool
	}{
		{
			name:           "TestCasbinFsAdapter_Integration_Anonymous_IndexShouldBeTrue",
			sub:            "",
			obj:            "/",
			act:            "GET",
			expectedResult: true,
		},
		{
			name:           "TestCasbinFsAdapter_Integration_Anonymous_UsersShouldBeFalse",
			sub:            "",
			obj:            "/user",
			act:            "GET",
			expectedResult: false,
		},
		{
			name:           "TestCasbinFsAdapter_Integration_Anonymous_UserShouldBeFalse",
			sub:            "",
			obj:            "/user/786",
			act:            "GET",
			expectedResult: false,
		},
		{
			name:           "TestCasbinFsAdapter_Integration_Anonymous_CreateUserShouldBeFalse",
			sub:            "",
			obj:            "/user",
			act:            "POST",
			expectedResult: false,
		},
		{
			name:           "TestCasbinFsAdapter_Integration_Anonymous_UpdateUserShouldBeFalse",
			sub:            "",
			obj:            "/user/786",
			act:            "PATCH",
			expectedResult: false,
		},
		{
			name:           "TestCasbinFsAdapter_Integration_Anonymous_SignInShouldBeTrue",
			sub:            "",
			obj:            "/sign-in",
			act:            "POST",
			expectedResult: true,
		},
		{
			name:           "TestCasbinFsAdapter_Integration_Admin_IndexShouldBeTrue",
			sub:            "admin",
			obj:            "/",
			act:            "GET",
			expectedResult: true,
		},
		{
			name:           "TestCasbinFsAdapter_Integration_Admin_UsersShouldBeTrue",
			sub:            "admin",
			obj:            "/user",
			act:            "GET",
			expectedResult: true,
		},
		{
			name:           "TestCasbinFsAdapter_Integration_Admin_UserShouldBeTrue",
			sub:            "admin",
			obj:            "/user/786",
			act:            "GET",
			expectedResult: true,
		},
		{
			name:           "TestCasbinFsAdapter_Integration_Admin_CreateUserShouldBeTrue",
			sub:            "admin",
			obj:            "/user",
			act:            "POST",
			expectedResult: true,
		},
		{
			name:           "TestCasbinFsAdapter_Integration_Admin_UpdateUserShouldBeTrue",
			sub:            "admin",
			obj:            "/user/786",
			act:            "PATCH",
			expectedResult: true,
		},
		{
			name:           "TestCasbinFsAdapter_Integration_Admin_SignInShouldBeTrue",
			sub:            "admin",
			obj:            "/sign-in",
			act:            "POST",
			expectedResult: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fs := os.DirFS("./examples/config/")
			policies := NewAdapter(fs, "policy.csv")
			model, err := NewModel(fs, "model.conf")
			if err != nil {
				assert.Fail(t, err.Error())
			}

			enforcer, err := casbin.NewEnforcer(model, policies)
			if err != nil {
				assert.Fail(t, err.Error())
			}

			err = enforcer.LoadPolicy()
			if err != nil {
				assert.Fail(t, err.Error())
			}

			actualResult, err := enforcer.Enforce(test.sub, test.obj, test.act)
			if err != nil {
				assert.Fail(t, err.Error())
			}

			assert.Equal(t, test.expectedResult, actualResult)
		})
	}
}
