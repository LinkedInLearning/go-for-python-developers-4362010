from dataclasses import dataclass


@dataclass
class Location:
    lat: float
    lng: float

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
    id: str


def move_all(items, lat, lng):
    for item in items:
        item.move(lat, lng)


if __name__ == '__main__':
    items = [
        Location(32.3477669, 34.9160405),
        Car(id='pyth0n', lat=32.5253837, lng=34.9427434),
    ]
    move_all(items, 32.0641339, 34.8742343)
    for item in items:
        print(item)
