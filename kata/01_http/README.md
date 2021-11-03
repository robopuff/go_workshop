# K1 HTTP

Based on `basic/02_json` connection, let's add HTTP server that will
export the results to RESTful API.

The API should contain:
- _GET_ `/word/:word` Look for specific word or list of words, if separated by comma
- _GET_ `/healthz`    Check app status, should return host and server time

All responses should come with `Content-Type: application/json`
