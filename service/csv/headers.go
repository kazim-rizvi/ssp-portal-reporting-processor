package csv

type CreativeReviewModel struct {
	CreativeReviewModelOnDemand
	CreativeReviewModelScheduled
}

type CreativeReviewModelOnDemand struct {
	Id   int
	Crid string
}

type CreativeReviewModelScheduled struct {
	IdBig   int
	CridBig string
}
