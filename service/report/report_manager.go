package report

import (
	"ssp-portal-reporting-processor/dao"
	"ssp-portal-reporting-processor/model"
)

type ReportManager struct {
	MobileAdPrimaryConnection *dao.DataFetcher
	TvAdPiConnection          *dao.DataFetcher
	ReportRequest             *model.ReportRequest
}

func NewReportManager(mobileAdPrimaryConnection *dao.DataFetcher,
	tvAdPiConnection *dao.DataFetcher, reportRequest *model.ReportRequest) *ReportManager {
	return &ReportManager{
		MobileAdPrimaryConnection: mobileAdPrimaryConnection,
		TvAdPiConnection:          tvAdPiConnection,
		ReportRequest:             reportRequest,
	}
}

