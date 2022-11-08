class MagicError(Exception):
    pass


def file_type(file_name):
    with open(file_name, 'rb') as fp:
        data = fp.read(64)

    if data[:4] == b'\x00\x00\x01\x00':
        return 'ico'
    if data[:6] == b'\x47\x49\x46\x38\x37\x61':
        return 'gif'
    if data[:4] == b'\xFF\xD8\xFF\xF0':
        return 'jpg'
    if data[:8] == b'\x89\x50\x4E\x47\x0D\x0A\x1A\x0A':
        return 'png'

    raise MagicError(f'{file_name!r} - unknown file type')


if __name__ == '__main__':
    file_name = 'Ch03/challenge/head.png'
    try:
        typ = file_type(file_name)
        print(f'{file_name} -> {typ}')
    except (FileNotFoundError, MagicError) as err:
        print(f'error: {err}')
