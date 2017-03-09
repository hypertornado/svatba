package main

import (
	"github.com/hypertornado/prago"
	"github.com/hypertornado/prago/extensions"
	administration "github.com/hypertornado/prago/extensions/admin"
	"math/rand"
	//"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const (
	version = "1.1.6"
)

var admin *administration.Admin

func main() {
	app := prago.NewApp("svatba", version)

	//app.AddMiddleware(&extensions.Sessions{})
	//app.AddMiddleware(&extensions.Gorm{})
	//admin = administration.NewAdmin("/admin", "Svatba Bude!")
	//app.AddMiddleware(admin)

	app.AddMiddleware(extensions.BuildMiddleware{[][2]string{{"public", ""}, {"templates", ""}}})
	app.AddMiddleware(prago.MiddlewareServer{start})
	prago.Must(app.Init())
}

func start(app *prago.App) {
	//prago.Must(app.LoadTemplatePath("templates/*"))
}

/*
func start(app *prago.App) {
	prago.Must(app.LoadTemplatePath("templates/*"))

	app.MainController().Get("/fotky", func(request prago.Request) {
		var files []*administration.File
		prago.Must(admin.Query().OrderDesc("id").Limit(20).Get(&files))
		list := []map[string]interface{}{}
		for _, v := range files {
			list = append(list, toLayout(*v))
		}
		request.SetData("files", list)

		request.SetData("yield", "photos")
		prago.Render(request, 200, "_layout")
	})

	app.MainController().Post("/nahrat", func(request prago.Request) {

		fileUploadPath := app.Config().GetString("fileUploadPath")
		if !strings.HasSuffix(fileUploadPath, "/") {
			fileUploadPath += "/"
		}

		multipartFiles := request.Request().MultipartForm.File["file"]
		if len(multipartFiles) != 1 {
			panic("must have 1 file selected")
		}

		file, err := administration.UploadFile(multipartFiles[0], fileUploadPath)
		if err != nil {
			panic(err)
		}
		file.Description = request.Params().Get("description")

		if request.Params().Get("print") == "on" {
			file.UserId = -2
		}

		prago.Must(admin.Create(file))
		prago.Redirect(request, "/nahrano/"+file.UID)
	})

	app.MainController().Get("/nahrano/:uid", func(request prago.Request) {
		var file administration.File
		prago.Must(admin.Query().WhereIs("uid", request.Params().Get("uid")).Get(&file))

		request.SetData("img", toLayout(file))

		request.SetData("yield", "loaded")
		prago.Render(request, 200, "_layout")
	})

	app.MainController().Get("/fotky/vse", func(request prago.Request) {
		var files []*administration.File
		prago.Must(admin.Query().OrderDesc("id").Get(&files))

		request.SetData("files", toListLayout(files))
		request.SetData("yield", "list")
		prago.Render(request, 200, "_layout")
	})

	app.MainController().Get("/fotky/slideshow", func(request prago.Request) {
		request.SetData("yield", "slideshow")
		prago.Render(request, 200, "_layout")
	})

	var minutesUntilShuffle = 2

	app.MainController().Get("/api/slideshow", func(request prago.Request) {
		if request.Params().Get("type") == "latest" {
			var file administration.File
			prago.Must(admin.Query().OrderDesc("id").Limit(1).Get(&file))
			if file.CreatedAt.After(time.Now().Add(-(time.Duration(minutesUntilShuffle) * time.Minute))) {
				administration.WriteApi(request, toLayout(file), 200)
				return
			}
		}

		var files []*administration.File
		prago.Must(admin.Query().OrderDesc("id").Get(&files))

		rn := rand.Intn(len(files))
		file := files[rn]
		lf := toLayout(*file)

		administration.WriteApi(request, lf, 200)

	})

	app.MainController().Get("/api/print", func(request prago.Request) {
		var files []*administration.File
		prago.Must(admin.Query().OrderDesc("id").WhereIs("userid", -2).Get(&files))
		administration.WriteApi(request, toListLayout(files), 200)
	})

}

func toListLayout(files []*administration.File) []map[string]interface{} {
	list := []map[string]interface{}{}
	for _, v := range files {
		list = append(list, toLayout(*v))
	}
	return list
}

func toLayout(file administration.File) map[string]interface{} {
	ret := make(map[string]interface{})
	ret["description"] = file.Description
	ret["url"] = file.GetLarge()
	ret["print"] = false
	if file.UserId == -2 {
		ret["print"] = true
	}
	ret["uid"] = file.UID
	ret["original"] = file.GetOriginal()
	return ret
}
*/
