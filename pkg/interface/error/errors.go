package error

import (
	"github.com/stackus/errors"
	"interview/pkg/core/errs"

	"net/http"
)

var ErrRedirect = errs.RedirectNotFoundError(nil, http.StatusFound)
var ErrInternalServerError = errs.ThrowInternalServerError(errors.ErrInternalServerError)
