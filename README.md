convert-format
This application converts a JSON object to a new format as requested in the problem statement.

Documentation
This documentation outlines how to use the convert-object endpoint of the Go server to convert a JSON object into a desired format and send it to a webhook.

Endpoint Details
Base URL: http://localhost:8080/
Endpoint: /convert-object
Method: POST
Content-Type: application/json
Parameters:
webhook_url: The URL of the webhook where the converted object will be sent.
Example Request
json
Copy code
POST http://localhost:8080/convert-object
Content-Type: application/json

{
  "ev": "top_cta_clicked",
  "et": "clicked",
  "id": "cl_app_id_001",
  "uid": "cl_app_id_001-uid-001",
  "mid": "cl_app_id_001-uid-001",
  "t": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
  "p": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
  "l": "en-US",
  "sc": "1920 x 1080",
  "atrk1": "button_text",
  "atrv1": "Free trial",
  "atrt1": "string",
  ...
}
Response
Status Code: 200 OK
Content-Type: text/plain
Body: "Request received and sent to worker successfully"
Example Postman Usage
Open Postman and create a new POST request.
Set the request URL to http://localhost:8080/convert-object.
Set the request body to a JSON object containing the data to be converted.
Add a parameter webhook_url with the value of the webhook URL where you want to send the converted object.
Click on "Send" to make the request.
