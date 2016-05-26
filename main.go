package main

import (
	"github.com/hypertornado/prago"
	"github.com/hypertornado/prago/extensions"
)

func main() {
	app := prago.NewApp("svatba", "1")
	app.AddMiddleware(extensions.BuildMiddleware{[][2]string{{"public", ""}}})
	app.AddMiddleware(prago.MiddlewareRun{start})
	prago.Must(app.Init())
}

func start(app *prago.App) {
	prago.Must(app.LoadTemplatePath("templates/*"))

	app.MainController().Get("/fotky", func(request prago.Request) {
		prago.Render(request, 200, "photos")
	})
}
