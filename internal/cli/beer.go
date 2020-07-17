package cli

import (
	"encoding/csv"
	"io"
	"strconv"
)

// Volume representation the volume of beer into data struct
type Volume struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

// Beer representation of beer into data struct
type Beer struct {
	ProductID        int     `json:"id"`
	Name             string  `json:"name"`
	Tagline          string  `json:"tagline"`
	FirstBrewed      string  `json:"first_brewed"`
	Description      string  `json:"description"`
	ImageURL         string  `json:"image_url"`
	Abv              float32 `json:"abv"`
	Ibu              float32 `json:"ibu"`
	TargetFg         float32 `json:"target_fg"`
	TargetOg         float32 `json:"target_og"`
	Ebc              float32 `json:"ebc"`
	Srm              float32 `json:"srm"`
	Ph               float32 `json:"ph"`
	AttenuationLevel float32 `json:"attenuation_level"`
	VolumeBeer       *Volume `json:"volume"`
	BoilVolume       *Volume `json:"boil_volume"`
	Price            string  `json:"price"`
	BeerID           int     `json:"beer_id"`
	Category         string  `json:"category"`
	Brewer           string  `json:"brewer"`
	Country          string  `json:"country"`
}

//ToArray transforma struct a array
func ToArray(rm Beer) []string {
	arrayBeer := []string{strconv.Itoa(rm.ProductID),
		rm.Name,
		rm.Tagline,
		rm.FirstBrewed,
		rm.Description,
		rm.ImageURL,
		rm.Price,
		strconv.Itoa(rm.BeerID),
		rm.Category,
		rm.Brewer,
		rm.Country}
	return arrayBeer
}

//CSVheader func header
func CSVheader(w io.Writer) {
	cw := csv.NewWriter(w)
	cw.Write([]string{"ProductID", "Name", "Tagline", "FirstBrewed", "Description", "ImageURL", "Price", "BeerID", "Category", "Brewer", "Country"})
	cw.Flush()
}

//BeerRow func body
func (rm *Beer) BeerRow() []string {
	return []string{
		strconv.Itoa(rm.ProductID),
		rm.Name,
		rm.Tagline,
		rm.FirstBrewed,
		rm.Description,
		rm.ImageURL,
		rm.Price,
		strconv.Itoa(rm.BeerID),
		rm.Category,
		rm.Brewer,
		rm.Country,
		strconv.Itoa(rm.VolumeBeer.Value),
		rm.VolumeBeer.Unit}
}

//VolumerRow func body
func (rm *Volume) VolumerRow() []string {
	return []string{strconv.Itoa(rm.Value), rm.Unit}
}
