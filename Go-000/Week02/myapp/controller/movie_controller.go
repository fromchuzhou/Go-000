package controller

import (
	"myapp/repository"
	"myapp/service"

	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {
}

func (controller *MovieController) Get() mvc.View {
	movieReposiroty := repository.NewMovieManager()
	movieService := service.NewMovieServiceManger(movieReposiroty)
	MoiveResult := movieService.ShowMovieName()

	return mvc.View{
		Name: "movie/index.html",
		Data: MoiveResult,
	}
}
