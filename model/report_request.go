package model

type ReportRequest struct {
	FromDate      string `json:"fromDate"`
	ToDate        string `json:"toDate"`
	DateRangeType string `json:"dateRangeType"`
	Active        string `json:"active"`
	Flag          string `json:"flag"`
	ReportRequestFilters
}

type ReportRequestFilters struct {
	AppName             string `json:"appName"`
	Dsp                 string `json:"dsp"`
	MobileAdUnit        string `json:"mobileAdUnit"`
	Country             string `json:"country"`
	AdType              string `json:"adType"`
	SupplyType          string `json:"supplyType"`
	Crid                string `json:"crid"`
	DetectedLang        string `json:"detectedLang"`
	Reviewer            string `json:"reviewer"`
	Title               string `json:"title"`
	ImageUrl            string `json:"imageUrl"`
	ThumbnailUrl        string `json:"thumbnailUrl"`
	VideoUrl            string `json:"videoUrl"`
	AppLinkUrl          string `json:"appLinkUrl"`
	DetectedCategoryUrl string `json:"detectedCategoryUrl"`
	RejectComment       string `json:"rejectComment"`
	Team                string `json:"team"`
	Pic                 string `json:"pic"`
}
