class Location:
    def __init__(self, lat, lng):
        if lat < -90 or lat > 90:
            raise ValueError(f'invalid lat: {lat}')
        if lng < -180 or lng > 180:
            raise ValueError(f'invalid lng: {lng}')

        self.lat = lat
        self.lng = lng

    def __repr__(self):
        cls = self.__class__.__name__
        return f'{cls}({self.lat!r}, {self.lng!r})'


if __name__ == '__main__':
    loc = Location(32.5253837, 34.9427434)
    print(loc)
