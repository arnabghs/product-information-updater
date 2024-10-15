package updatePriceModel

// RequestBody represents the structure of the incoming POST request
type RequestBody struct {
	ID      string `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
	// Add other fields as necessary
}

type ProductEvent struct {
	ID        string `json:"id" bson:"id"`
	Message   string `json:"message" bson:"message"`
	ProductID string `json:"productID" bson:"productID"`
}
