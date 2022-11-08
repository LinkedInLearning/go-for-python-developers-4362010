package main

import "fmt"

type Location struct {
	Lat float64
	Lng float64
}

func NewLocation(lat, lng float64) (Location, error) {
	if lat < -90 || lat > 90 {
		return Location{}, fmt.Errorf("invalid lat: %#v", lat)
	}
	if lng < -180 || lng > 180 {
		return Location{}, fmt.Errorf("invalid lng: %#v", lng)
	}

	loc := Location{lat, lng}
	return loc, nil
}

func (l *Location) Move(lat, lng float64) {
	l.Lat = lat
	l.Lat = lng
}

type Car struct {
	ID string
	Location
}

func NewCar(id string, lat, lng float64) (Car, error) {
	loc, err := NewLocation(lat, lng)
	if err != nil {
		return Car{}, err
	}

	car := Car{
		ID:       id,
		Location: loc,
	}
	return car, nil
}

type Mover interface {
	Move(float64, float64)
}

func moveAll(items []Mover, lat, lng float64) {
	for _, item := range items {
		item.Move(lat, lng)
	}
}

func main() {

	items := []Mover{
		&Location{32.3477669, 34.9160405},
		&Car{
			ID: "g0ph3r",
			Location: Location{
				Lat: 32.5253837,
				Lng: 34.9427434,
			},
		},
	}
	moveAll(items, 32.0641339, 34.8742343)
	for _, item := range items {
		fmt.Printf("%#v\n", item)

	}
}
