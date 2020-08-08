package fetching_test

import (
	"errors"
	"testing"

	beerscli "github.com/patriciabonaldy/punkapi/domain/entity"
	service "github.com/patriciabonaldy/punkapi/domain/ports"
	mock "github.com/patriciabonaldy/punkapi/domain/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestFetchBeers(t *testing.T) {
	repo := new(mock.MockRepository)
	repo.On("GetBeers").Return(buildMockBeers(), nil)
	fetch := service.NewFetchy(repo)

	b, err := fetch.FetchBeers()
	if err != nil {
		assert.Error(t, err)
	}
	assert.NotNil(t, b)
}

func buildMockBeers() []beerscli.Beer {
	return []beerscli.Beer{
		beerscli.NewBeer(
			127,
			"Mad Jack Mixer",
			"Molson",
			"Canada",
			"Domestic Specialty",
			"23.95",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Canada",
			"Grolsch Export B.V.",
			"49.50",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			124,
			"Mad Jack Mixer",
			"Molson",
			"Canada",
			"Domestic Specialty",
			"23.95",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			8520139,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Canada",
			"Grolsch Export B.V.",
			"49.50",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			155,
			"Mad Jack Mixer",
			"Molson",
			"Canada",
			"Domestic Specialty",
			"23.95",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			8520160,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Canada",
			"Grolsch Export B.V.",
			"49.50",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			127,
			"Mad Jack Mixer",
			"Molson",
			"Canada",
			"Domestic Specialty",
			"23.95",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Canada",
			"Grolsch Export B.V.",
			"49.50",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			127,
			"Mad Jack Mixer",
			"Molson",
			"Canada",
			"Domestic Specialty",
			"23.95",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Canada",
			"Grolsch Export B.V.",
			"49.50",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			127,
			"Mad Jack Mixer",
			"Molson",
			"Canada",
			"Domestic Specialty",
			"23.95",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Canada",
			"Grolsch Export B.V.",
			"49.50",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			127,
			"Mad Jack Mixer",
			"Molson",
			"Canada",
			"Domestic Specialty",
			"23.95",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Canada",
			"Grolsch Export B.V.",
			"49.50",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			127,
			"Mad Jack Mixer",
			"Molson",
			"Canada",
			"Domestic Specialty",
			"23.95",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Canada",
			"Grolsch Export B.V.",
			"49.50",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			127,
			"Mad Jack Mixer",
			"Molson",
			"Canada",
			"Domestic Specialty",
			"23.95",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Canada",
			"Grolsch Export B.V.",
			"49.50",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			127,
			"Mad Jack Mixer",
			"Molson",
			"Canada",
			"Domestic Specialty",
			"23.95",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
		beerscli.NewBeer(
			8520130,
			"Grolsch 0.0",
			"Non-Alcoholic Beer",
			"Canada",
			"Grolsch Export B.V.",
			"49.50",
			"https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/58651371/",
		),
	}

}

func TestFetchByID(t *testing.T) {
	repo := new(mock.MockRepository)
	repo.On("GetBeers").Return(buildMockBeers(), nil)

	//se defina una tabla de set de datos
	tests := map[string]struct {
		repo  repo.BeerRepo
		input int
		want  int
		err   error
	}{
		"valid beer":            {repo: repo, input: 127, want: 127, err: nil},
		"not found beer":        {repo: repo, input: 99999, err: errors.New("error")},
		"error with repository": {repo: repo, err: errors.New("error")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			fetch := service.NewFetchy(tc.repo)
			b, err := fetch.FetchByID(tc.input)

			if tc.err != nil {
				assert.Error(t, err)
			}

			if tc.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, tc.want, b.ProductID)

		})
	}

}
