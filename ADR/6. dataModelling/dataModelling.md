## Product_Update_Information_Schema

**Database Type:** NoSQL

**Common Queries:**

- get document by productId
- get document by updatedBy
- get document by updateAttribute
- get document by updatedAt

**Indexing:**

Since search by **productID,** **updatedAttribute, updatedBy** & **updatedAt**
are most common, we should create index on that

```
{
  "producId": "",
  "productName": "",
  "updateAttribute": "PRICE",
  "locations": [
    {
      "zoneID": "SWD01",
      "locationID": "STK99",
      "value": "199.99",
      "currency": "EUR"
    }
  ],
  "updatedBy": "abc",
  "updatedAt": "13:14:15Z"
}
```