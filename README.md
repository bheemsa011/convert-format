# convert-format

This application converts a JSON object to a new format as requested in the problem statement.

## Documentation

This documentation outlines how to use the `convert-object` endpoint of the Go server to convert a JSON object into a desired format and send it to a webhook.

## How to use : 

Clone this repo
run main.go file in terminal

### Endpoint Details

- **Base URL:** http://localhost:8080/
- **Endpoint:** `/convert-object`
- **Method:** POST
- **Content-Type:** application/json
- **Parameters:**
  - `webhook_url`: The URL of the webhook where the converted object will be sent.

### Example Request

```http
POST http://localhost:8080/convert-object?webhook_url=https://webhook.site/d18d2a54-cc8a-40d8-b892-77ab1a853940
Content-Type: application/json
