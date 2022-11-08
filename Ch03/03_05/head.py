def file_head(file_name, size):
    with open(file_name, 'rb') as fp:
        data = fp.read(size)

    if len(data) < size:
        raise ValueError(f'{file_name!r} too small')

    return data


if __name__ == '__main__':
    data = file_head('Ch03/03_05/head.png', 8)
    print(data)
