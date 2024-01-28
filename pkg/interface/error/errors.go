package error

import (
	"net/http"

	"github.com/stackus/errors"
	"interview/pkg/core/errs"
)

var (
	ErrRedirect            = errs.RedirectNotFoundError(nil, http.StatusFound)
	ErrInternalServerError = errs.ThrowInternalServerError(errors.ErrInternalServerError)
)
