package trackingmore

// Response is the message envelope for the TrackingMore API response
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta is used to communicate extra information about the response to the developer.
type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
