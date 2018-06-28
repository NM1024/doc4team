package main

import (
	"doc4team/web/controllers"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/view"

	"doc4team/tools"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// Load the template files.
	views := iris.HTML("./views", ".html")
	app.RegisterView(views)
	templateFunc(views)
	app.StaticWeb("/", "./public")
	// Register our controllers.
	// mvc.New(app.Party("/movies")).Handle(new(controllers.MovieController))
	// But,
	// You can also split the code you write to configure an mvc.Application
	// using the `mvc.Configure` method, as shown below.
	mvc.Configure(app.Party("/doc"), doc)
	mvc.New(app).Handle(new(controllers.HomeController))

	// http://localhost:8080/movies
	// http://localhost:8080/movies/1
	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:"+os.Getenv("PORT")),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

// note the mvc.Application, it's not iris.Application.
func doc(app *mvc.Application) {
	// Add the basic authentication(admin:password) middleware
	// for the /movies based requests.
	// app.Router.Use(middleware.BasicAuth)

	// Create our movie repository with some (memory) data from the datasource.
	// repo := repositories.NewMovieRepository(datasource.Movies)
	// Create our movie service, we will bind it to the movie app's dependencies.
	// app.Register(services.GetApiDocInstance())
	// app.Register(services.GetToMarkDownInstance())

	// serve our movies controller.
	// Note that you can serve more than one controller
	// you can also create child mvc apps using the `movies.Party(relativePath)` or `movies.Clone(app.Party(...))`
	// if you want.
	app.Handle(new(controllers.DocController))
}

func templateFunc(html *view.HTMLEngine) {
	html.AddFunc("HttpMethodList", tools.HttpMethodList)
	html.AddFunc("Obj2String", tools.Obj2String)
}
