from dataclasses import dataclass


@dataclass
class Location:
    lat: int
    lng: int

    def __post_init__(self):
        if self.lat < -90 or self.lat > 90:
            raise ValueError(f'invalid lat: {self.lat}')
        if self.lng < -180 or self.lng > 180:
            raise ValueError(f'invalid lng: {self.lng}')

    def move(self, lat, lng):
        self.lat = lat
        self.lng = lng


@dataclass
class Car(Location):
    id: int


if __name__ == '__main__':
    car = Car(id='g0ph3r', lat=32.5253837, lng=34.9427434)
    car.move(32.0641339, 34.8742343)
    print(car)
