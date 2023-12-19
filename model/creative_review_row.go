package model

import "fmt"

type CreativeReviewModel struct {
	CreativeReviewModelOnDemand
	CreativeReviewModelScheduled
}

type CreativeReviewModelOnDemand struct {
	Id            int
	Crid          string
	SubmittedDate string
	CreatedDate   string
}

type CreativeReviewModelScheduled struct {
	IdBig   int
	CridBig string
}

func (c CreativeReviewModelOnDemand) toCsvString() string {
	return fmt.Sprintf("%d,%s,%s,%s\n", c.Id, c.Crid, c.SubmittedDate, c.CreatedDate)
}
