package handler

import (
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"interview/pkg/core/common"
	"interview/pkg/core/dto"
	"interview/pkg/core/port"
	errHandler "interview/pkg/interface/error"
	"interview/static"
)

type TaxController struct {
	cartService port.CartService
}

func NewTaxController(cartService port.CartService) *TaxController {
	return &TaxController{cartService: cartService}
}

func (t *TaxController) ShowAddItemForm(c *gin.Context) {
	//_, err := c.Request.Cookie("ice_session_id")
	session := sessions.Default(c)
	cookie := session.Get(common.IceSessionIdKey)
	if cookie == nil {
		session.Set(common.IceSessionIdKey, uuid.NewString())
		err := session.Save()
		if err != nil {
			errHandler.HandleError(c, errHandler.ErrRedirect)
			return
		}
		cookie = session.Get(common.IceSessionIdKey)
	}

	data := map[string]interface{}{
		"Error": c.Query("error"),
		//"cartItems": cartItems,
	}

	if cookie != nil {
		data["CartItems"] = t.cartService.GetCartData(cookie.(string))
	}

	html, err := renderTemplate(data)
	if err != nil {
		errHandler.HandleError(c, errHandler.ErrInternalServerError)
		return
	}

	c.Header("Content-Type", "text/html")
	c.String(200, html)
}

func (t *TaxController) AddItem(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get(common.IceSessionIdKey)
	if value == nil {
		errHandler.HandleError(c, errHandler.ErrRedirect)
		return
	}

	form, err := getCartItemForm(c)
	if err != nil {
		errHandler.HandleError(c, errHandler.ErrRedirect.Msg(err.Error()))
		return
	}

	err = t.cartService.AddItemToCart(form, value.(string))
	if err != nil {
		errHandler.HandleError(c, err)
		return
	}
	c.Redirect(302, "/")
}

func (t *TaxController) DeleteCartItem(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get(common.IceSessionIdKey)
	if value == nil {
		errHandler.HandleError(c, errHandler.ErrRedirect)
		return
	}

	cartItemIDString := c.Query("cart_item_id")
	if cartItemIDString == "" {
		errHandler.HandleError(c, errHandler.ErrRedirect.Msg("invalid cart item id"))
		return
	}

	err := t.cartService.DeleteCartItem(cartItemIDString, value.(string))
	if err != nil {
		errHandler.HandleError(c, err)
		return
	}

	c.Redirect(302, "/")
}

func getCartItemForm(c *gin.Context) (*dto.CartItemForm, error) {
	if c.Request.Body == nil {
		return nil, fmt.Errorf("body cannot be nil")
	}

	form := &dto.CartItemForm{}

	if err := binding.FormPost.Bind(c.Request, form); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return form, nil
}

func renderTemplate(pageData interface{}) (string, error) {
	// Read and parse the HTML template file
	// template.ParseFS(static.WebUI)

	tmpl := template.Must(template.New("").ParseFS(static.WebUI, "*.html"))
	if tmpl == nil {
		return "", fmt.Errorf("Error parsing template ")
	}

	// Create a strings.Builder to store the rendered template
	var renderedTemplate strings.Builder

	err := tmpl.ExecuteTemplate(&renderedTemplate, "add_item_form.html", pageData)
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v ", err)
	}

	// Convert the rendered template to a string
	resultString := renderedTemplate.String()

	return resultString, nil
}
