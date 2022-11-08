def split_ext(path):
    i = path.rfind('.')
    if i == -1:
        return path, ''
    return path[:i], path[i:]


if __name__ == '__main__':
    root, ext = split_ext('app.py')
    print(f'{root=}, {ext=}')
