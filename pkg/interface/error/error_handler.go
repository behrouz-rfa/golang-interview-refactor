package error

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"interview/pkg/core/errs"
	"net/http"
)

func HandleError(c *gin.Context, err error) {
	e, ok := err.(*errs.HTTPError)
	if ok {
		switch e.Code {
		case http.StatusFound:
			if len(e.Message) > 0 {
				c.Redirect(e.Code, fmt.Sprintf("/?error=%s", e.Message))
				return
			}
			c.Redirect(e.Code, "/")
			return
		}
		c.AbortWithStatus(e.Code)
		return
	}

	c.Redirect(302, "/")

	return
}
