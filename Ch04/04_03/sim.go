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
	l.Lng = lng
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

func main() {
	car, err := NewCar("g0ph3r", 32.5253837, 34.9427434)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	car.Move(32.0641339, 34.8742343)
	fmt.Printf("%#v\n", car)
}
