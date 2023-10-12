package model

type CreativeReviewModel struct {
	CreativeReviewModelLight
	CreativeReviewModelHeavy
}

type CreativeReviewModelLight struct {
	Id   int
	Crid string
}

type CreativeReviewModelHeavy struct {
	IdBig   int
	CridBig string
}
