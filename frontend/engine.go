package frontend

import (
	"github.com/gofiber/template/html"
)

func NewEngine() *html.Engine {
	engine := html.New("frontend/pages", ".html")
	engine.Reload(true)

	return engine
}
