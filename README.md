# mri

Little RESTful API that exposes NMAP functionalities.

## Request

```json
{
  "host": "8.8.8.8",
  "ports": [
    53
  ]
}
```

## Response

```json
{
  "request": {
    "host": "8.8.8.8",
    "ports": [
      53
    ]
  },
  "results": {
    "53": {
      "proto": "tcp",
      "service": "tcpwrapped",
      "state": "open",
      "version": ""
    }
  }
}
```