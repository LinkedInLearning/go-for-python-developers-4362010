from flask import Flask

app = Flask(__name__)


@app.route('/health')
def health():
    return 'OK\n'  # FIXME


if __name__ == '__main__':
    app.run(port=8080)
