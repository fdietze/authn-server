package services

import (
	"strconv"

	"github.com/keratin/authn-server/lib/compat"
	"github.com/keratin/authn-server/ops"
	"github.com/pquerna/otp/totp"

	"github.com/keratin/authn-server/app"
	"github.com/keratin/authn-server/app/data"
	"github.com/keratin/authn-server/app/tokens/resets"
	"github.com/pkg/errors"
)

func PasswordResetter(store data.AccountStore, r ops.ErrorReporter, cfg *app.Config, token string, password string, totpCode string) (int, error) {
	claims, err := resets.Parse(token, cfg)
	if err != nil {
		return 0, FieldErrors{{"token", ErrInvalidOrExpired}}
	}

	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, errors.Wrap(err, "Atoi")
	}

	account, err := store.Find(id)
	if err != nil {
		return 0, errors.Wrap(err, "Find")
	}
	if account == nil {
		return 0, FieldErrors{{"account", ErrNotFound}}
	} else if account.Locked {
		return 0, FieldErrors{{"account", ErrLocked}}
	} else if account.Archived() {
		return 0, FieldErrors{{"account", ErrLocked}}
	}

	if claims.LockExpired(account.PasswordChangedAt) {
		return 0, FieldErrors{{"token", ErrInvalidOrExpired}}
	}

	//Check OTP MFA
	if account.TOTPEnabled() {
		secret, err := compat.Decrypt([]byte(account.TOTPSecret.String), cfg.DBEncryptionKey)
		if err != nil {
			return 0, errors.Wrap(err, "TOTPDecrypt")
		}
		if !totp.Validate(totpCode, secret) {
			return 0, FieldErrors{{"otp", ErrInvalidOrExpired}}
		}
	}

	return account.ID, PasswordSetter(store, r, cfg, id, password)
}
