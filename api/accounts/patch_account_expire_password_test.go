package accounts_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/keratin/authn-server/api/accounts"
	"github.com/keratin/authn-server/api/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPatchAccountExpirePassword(t *testing.T) {
	app := test.App()
	server := test.Server(app, accounts.Routes(app))
	defer server.Close()

	client := test.NewClient(server).Authenticated(app.Config)

	t.Run("unknown account", func(t *testing.T) {
		res, err := client.Patch("/accounts/999999/expire_password", url.Values{})
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("active account", func(t *testing.T) {
		account, err := app.AccountStore.Create("active@test.com", []byte("bar"))
		require.NoError(t, err)

		res, err := client.Patch(fmt.Sprintf("/accounts/%v/expire_password", account.Id), url.Values{})
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		account, err = app.AccountStore.Find(account.Id)
		require.NoError(t, err)
		assert.True(t, account.RequireNewPassword)
	})
}
