package repository

type MovieRepository interface {
	GetMovieName(Id string) (string, error)
}

type MovieManager struct {
}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName(Id string) (string, error) {
	//I don't know
	return
}
