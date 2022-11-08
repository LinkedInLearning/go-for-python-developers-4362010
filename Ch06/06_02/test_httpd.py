from http import HTTPStatus

import httpd


def test_health():
    with httpd.app.test_client() as client:
        resp = client.get('/health')
    assert resp.status_code == HTTPStatus.OK
