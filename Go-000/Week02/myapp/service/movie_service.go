package service

import (
	"fmt"
	"myapp/repository"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repository.MovieRepository
}

func NewMovieServiceManger(repo repository.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string {
	fmt.Println("查询电影名称")
	return m.repo.GetMovieName()
}
