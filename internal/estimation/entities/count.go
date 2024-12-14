package entities

type CountRequest struct {
	Segment string
}

type CountResponse struct {
	Count int64
}
