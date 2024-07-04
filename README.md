# Proxy-server

This is a proxy server built using Go. It forwards requests to a specified server and returns the response to the client. 

## Public URL

```bash
https://proxy-server-rl0v.onrender.com
```

## Getting started

1. Clone the repository:

    ```bash
    git clone https://github.com/luciorim/proxy-server.git
    cd proxy-server
    ```

2. Start the server:

    ```bash
    make up
    ```

## Endpoints

### `POST /proxy`

**Description:** Sends a proxy request and stores information about the request and response.

**Request Data:** JSON object with the following fields:
- `Method` (string) — HTTP request method (e.g., `"GET"` or `"POST"`).
- `URL` (string) — URL for the request.
- `Headers` (map[string]string) — Request headers.

**Response Data:** JSON object with the following fields:
- `ID` (string) — Unique identifier for the request.
- `URL` (string) — URL of the request.
- `Status` (int) — HTTP status code of the response.
- `Headers` (map[string][]string) — Response headers.
- `Length` (int) — Length of the response body in bytes.

**Example Request:**
```json
{
  "method": "GET",
  "url": "http://www.google.com",
  "headers": {
    "Authorization": "Basic kfdafgd7f6djfhksd_"
  }
}
```
**Example Response:**
```json
{
  "id": "1",
  "url": "http://www.google.com",
  "status": 200,
  "headers": {
      "Cache-Control": [
          "private, max-age=0"
      ],
      "Content-Security-Policy-Report-Only": [
          "object-src 'none';base-uri 'self';script-src 'nonce-LDRFyBmTRsMPB1L95xs_3A' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp"
      ],
      "Content-Type": [
          "text/html; charset=ISO-8859-1"
      ],
      "Date": [
          "Thu, 04 Jul 2024 18:35:54 GMT"
      ],
      "Expires": [
          "-1"
      ],
      "P3p": [
          "CP=\"This is not a P3P policy! See g.co/p3phelp for more info.\""
      ],
      "Server": [
          "gws"
      ],
      "Set-Cookie": [
          "AEC=AVYB7cp5ac_6hK1N509Jzv6IkTp4O1xQmoB0SAW-dF9hIPVRuNN6gguC7A; expires=Tue, 31-Dec-2024 18:35:54 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax",
          "NID=515=ri6fbJSmQoJm6_BOhJ84Z--mgn3sUu4TSoDPTUVBiQD5w232dA7kNlL4zZM_VAGmA1zlsLxd9ORBInMj1THLEvZrah3r7MY2MDSvG0h6hYYpXhxLl8LU73f_YenNrjvR2kyIWBoaEXp7bk89w85_Ks-zb6f-8n2tuUaUcjHGM4c; expires=Fri, 03-Jan-2025 18:35:54 GMT; path=/; domain=.google.com; HttpOnly"
      ],
      "X-Frame-Options": [
          "SAMEORIGIN"
      ],
      "X-Xss-Protection": [
          "0"
      ]
  },
  "length": 21187
}
```

### `GET /proxy`
**Description:** Returns history of all requests.

**Example Response:**
```json
{
    "1_req": {
        "method": "GET",
        "url": "http://www.google.com",
        "headers": {
            "Authorization": "Basic kfdafgd7f6djfhksd_"
        }
    },
    "1_res": {
        "id": "1",
        "url": "http://www.google.com",
        "status": 200,
        "headers": {
            "Cache-Control": [
                "private, max-age=0"
            ],
            "Content-Security-Policy-Report-Only": [
                "object-src 'none';base-uri 'self';script-src 'nonce-UDt8Q40WM5OCjZ8RRmgXYw' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp"
            ],
            "Content-Type": [
                "text/html; charset=ISO-8859-1"
            ],
            "Date": [
                "Thu, 04 Jul 2024 19:18:45 GMT"
            ],
            "Expires": [
                "-1"
            ],
            "P3p": [
                "CP=\"This is not a P3P policy! See g.co/p3phelp for more info.\""
            ],
            "Server": [
                "gws"
            ],
            "Set-Cookie": [
                "AEC=AVYB7coUJkX3DR8O3i6GEgseGPiEaYL6jGT2S899VF-kWeCy1tH82Rss5Vk; expires=Tue, 31-Dec-2024 19:18:45 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax",
                "NID=515=OvET70nUz3IHA9gne7tEUDDjQIB53azYlPQl9VSoQMeTrH5ptghq62v4E7H-_kbl_Lm87aFA9MPtJaakVvioAqrQ3KY2mQ945Mp09hPfCgOF3y4JN7rKn-ddhOcfkuINVqHB4aR3rEfJ9C5VCQfehEqqPCP7QvbqgrL4zKa9oBs; expires=Fri, 03-Jan-2025 19:18:45 GMT; path=/; domain=.google.com; HttpOnly"
            ],
            "X-Frame-Options": [
                "SAMEORIGIN"
            ],
            "X-Xss-Protection": [
                "0"
            ]
        },
        "length": 21566
    }
}

```
