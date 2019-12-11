package frontend

import "github.com/erdwolfapp/Erdwolf/internal/app"

type FrontendHandlers struct {
	app *app.Application
}

func New(app *app.Application) FrontendHandlers {
	return FrontendHandlers {
		app: app,
	}
}