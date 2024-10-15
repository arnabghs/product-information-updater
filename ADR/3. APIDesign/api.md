## For updating product information

**Description:**

This endpoint allows updating product information (such as price, discount, etc.) for different locations.

Since it can either update or create one product info, we use POST instead of PUT.

**HTTP Method:** POST
**API Endpoint:** api/v1/products/{productID}

### Request Schema:
```
{
  "producId": "",
  "productName": "",
  "update": [
    {
      "attribute": "PRICE",          // event name
      "locations": [
        {
          "zoneID": "SWD01",         // Sweden or Germany, country code   
          "locationID": "STK99",     // locality, i.e Stockholm or Berlin 
          "value": "199.99",
          "currency": "EUR"
        }
      ]
    },
    {
      "attribute": "COUNTRY",
      "locations": [
        {
          "zoneID": "",
          "locationID": "",
          "value": "50"              // assuming stocks
        }
      ]
    },
    {
      "attribute": "DISCOUNT",
      "locations": [
        {
          "zoneID": "",
          "locationID": "",
          "value": "10",
          "discountCode": "BLACK_DRIDAY"
        }
      ]
    }
  ]
}
```
Note: sender_ID can be fetched from Auth token

### Response Schema
Successful Update (200 OK):
```
{
  "productId": "12345",
  "updatedAttributes" : ["PRICE", "DISCOUNT"]
  "message": "Product information update request has been successfully placed"
}
```

Validation Error (400 Bad Request):

If the request contains invalid data (e.g., missing required fields, invalid values, wrong locations), the API will respond with an error.
```
{
  "error": "Invalid request data",
  "message": "The field 'currency' is missing for location 'loc-002'."
}
```

Product Not Found (404 Not Found)

If the productId specified does not exist in the system
```
{
  "error": "Product not found",
  "message": "No product found with ID '12345'."
}
```

Internal Server Error (500 Internal Server Error)
```
{
  "error": "Server error",
  "message": "An error occurred while updating product information."
}
```

--------------------------------------------

## For fetching product updates

**Description:**

This endpoint fetches the latest updated information for a specific product across various locations.

**HTTP Method:** GET \
**API Endpoint:** api/v1/products/{productID}

**QueryParams:**
- ?pricingManagerId=”abc123”
- ?attribute=”PRICE”
- ?startdate=”10.10.10”&enddate=”12.10.10”

**Request URL Params:**\
productId : unique ID for each product (String, required)

### Response Schemas:
Successful Response (200 OK):\
Returns the most recent updates for the product across all locations under given region
```
{
  "producId": "",
  "productName": "",
  "updates": [
    {
      "attribute": "PRICE",
      "locations": [
        {
          "zoneID": "SWD01",
          "locationID": "STK99",
          "value": "199.99",
          "currency": "EUR"
        }
      ]
    },
    {
      "attribute": "DISCOUNT",
      "locations": [
        {
          "zoneID": "",
          "locationID": "",
          "value": "10",
          "discountCode": "BLACK_DRIDAY"
        }
      ]
    }
  ],
  "updatedBy": "pricing_manager_1",
  "updatedAt": "13:14:15Z"  // UTC
}
```
<br>

Not Found (404 Not Found):\
If the productId specified does not exist in the system
```
{
  "error": "Product not found",
  "message": "No product found with ID '12345'."
}
```
<br>

Validation Error (400 Bad Request):\
If the query params values are invalid, the API will respond with an error
```
{
  "error": "Invalid request data",
  "message": "invalid attribute"
}
```
<br>

Internal Server Error (500 Internal Server Error):\
If there is an unexpected error on the server
```
{
  "error": "Server error",
  "message": "An error occurred while fetching product information."
}
```
