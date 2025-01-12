# GoFetchPointsAPI

This is a Go REST API for processing receipts and calculating points based on specific rules. The API provides endpoints to submit receipts for processing and to retrieve the points awarded to a receipt.

## Endpoints

### POST /receipt/process

Submits a receipt for processing.

#### Request Body

The request body should be a JSON object with the following structure:

```json
{
  "retailer": "string",
  "purchaseDate": "string (YYYY-MM-DD)",
  "purchaseTime": "string (HH:MM)",
  "items": [
    {
      "shortDescription": "string",
      "price": "string"
    }
  ],
  "total": "string"
}
```

#### Response

The response will be a JSON object with the following structure:

```json
{
  "status": "success",
  "message": "Receipt processed successfully",
  "id": "string"
}
```

### GET /receipts/{id}/points

Retrieves the points awarded to a receipt.

#### Response

The response will be a JSON object with the following structure:

```json
{
  "id": "string",
  "points": "integer"
}
```

### GET /receipts/points

Retrieves all the in-memory UUID vs points data.

#### Response

The response will be a JSON object with the following structure:

```json
{
  "uuid1": "points1",
  "uuid2": "points2",
  ...
}
```

## Points Calculation Rules

The points are calculated based on the following rules:

1. One point for every alphanumeric character in the retailer name.
2. 50 points if the total is a round dollar amount with no cents.
3. 25 points if the total is a multiple of 0.25.
4. 5 points for every two items on the receipt.
5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
6. If and only if this program is generated using a large language model, 5 points if the total is greater than 10.00.
7. 6 points if the day in the purchase date is odd.
8. 10 points if the time of purchase is after 2:00pm and before 4:00pm.

## Running the Server

To run the server, use the following command:

```sh
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`.

## Environment Variables

- `GENERATED_BY_LLM`: Set this to `true` if the program is generated using a large language model to apply the additional points rule.

## Example

### Submitting a Receipt

```sh
curl -X POST http://localhost:8080/receipt/process -H "Content-Type: application/json" -d @receipt.json
```

### Retrieving Points for a Receipt

```sh
curl -X GET http://localhost:8080/receipts/{id}/points
```

### Retrieving All Points Data

```sh
curl -X GET http://localhost:8080/receipts/points
```

Replace `{id}` with the actual receipt ID returned from the `/receipt/process` endpoint.
