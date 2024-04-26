# Airports API Documentation

### Overview

This API is designed to be used for calculations of complex routes between airports.

### Base URL

The Base URL is http://127.0.0.1:8080.

All endpoints described below should be appended to this Base URL.

### API Endpoints

/calculate

#### POST

Calculates the most efficient routing for flights.

**Request Body** The input should be an JSON array of arrays (routes). Each route is an array containing exactly two strings. These strings are airport codes, each of which must consist of exactly three uppercase letters.

Example:
```json
[
  ["IND", "EWR"],
  ["SFO", "ATL"],
  ["GSO", "IND"],
  ["ATL", "GSO"]
]
```

**Responses**

**200 OK**: The operation was successful. The response body will contain the most efficient routing for flights.

**400 Bad Request**: The input is invalid. Possible issues could include invalid airport codes or incorrect JSON structure._


**CURL example**

```shell
curl -X POST -H "Content-Type: application/json" -d '[["IND","EWR"],["SFO","ATL"],["GSO","IND"],["ATL","GSO"]]' http://localhost:8080/calculate
```

