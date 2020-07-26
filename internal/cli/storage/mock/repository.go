package mock

import (
	beerscli "github.com/patriciabonaldy/punkapi/internal"
	"github.com/stretchr/testify/mock"
)

//MockRepository struct Mock of Repository
type MockRepository struct {
	// add a Mock object instance
	mock.Mock
}

//GetBeers funct() ([]beerscli.Beer, error)
func (m *MockRepository) GetBeers() ([]beerscli.Beer, error) {
	args := m.Called()
	return args.Get(0).([]beerscli.Beer), args.Error(1)
}
