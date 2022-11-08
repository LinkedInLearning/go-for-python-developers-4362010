from urllib.request import URLError, urlopen


def check_url(url):
    try:
        with urlopen(url, timeout=5):
            return True
    except (URLError, TimeoutError):
        return False


if __name__ == '__main__':
    url = 'https://www.linkedin.com/learning/'
    print(check_url(url))
