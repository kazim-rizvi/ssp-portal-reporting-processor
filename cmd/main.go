package main

import (
	"log"
	"os"
	"ssp-portal-reporting-processor/config"
	"ssp-portal-reporting-processor/constants"
	"ssp-portal-reporting-processor/dao"
	"ssp-portal-reporting-processor/model"
	"ssp-portal-reporting-processor/service/aws/secrets_manager"
	"ssp-portal-reporting-processor/service/report"
	"ssp-portal-reporting-processor/utils"
	// "ssp-portal-reporting-processor/service/csv"
	// "ssp-portal-reporting-processor/email"
	// "ssp-portal-reporting-processor/s3"
)

func main() {

	// Fetch ENV Variables and Args
	args := os.Args
	currentProfile := os.Getenv(constants.PROFILE)
	projectVersion := os.Getenv(constants.PROJECT_VERSION)
	reportRequestJsonString := os.Getenv(constants.REPORT_REQUEST)

	// Log the data
	logData(projectVersion, currentProfile, args, reportRequestJsonString)

	// Load config for current profile
	config, err := config.LoadConfig(currentProfile)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		utils.ExitProgram(true)
	}

	// Deocode Secrets
	dbDetailRef, err := secrets_manager.RetrieveDbSecret(&config.DBConfig)
	if err != nil {
		log.Fatalf("Error fetching DB details from secrets manager: %v", err)
		utils.ExitProgram(true)
	}
	sesCredentialRef, err := secrets_manager.RetrieveEmailSecret(&config.EmailConfig)
	if err != nil {
		log.Fatalf("Error fetching SES credentials from secrets manager: %v", err)
		utils.ExitProgram(true)
	}
	s3CredentialRef := &config.S3Config

	//Unmarshall report request
	reportRequestRef, err := utils.UnmarshalJson[model.ReportRequest](reportRequestJsonString)

	// Log the data
	// TODO: Remove this
	utils.LogDetails(*dbDetailRef, "Db Details", false)
	utils.LogDetails(*sesCredentialRef, "SES Credentials", false)
	utils.LogDetails(*s3CredentialRef, "S3 Credentials", false)

	// Setup DB for MobileAd Primary and Tvad Pi
	mobileAdPrimaryConnectionRef, err := dao.NewDataFetcher(dbDetailRef.MobileAd.Primary)
	tvAdPiConnectionRef, err := dao.NewDataFetcher(dbDetailRef.TvAd.PIData)

	//Create Report Manager
	reportManager := report.NewReportManager(mobileAdPrimaryConnectionRef, tvAdPiConnectionRef, reportRequestRef)

	// fmt.Printf(cfg.DBConfig.Host)

	// Fetch data from the DB in a batched way
	// dataFetcher, err := db.NewDataFetcher(cfg.DBConfig)
	// if err != nil {
	// 	log.Fatalf("Error connecting to DB: %v", err)
	// }
	// data, err := dataFetcher.FetchDataBatched("select id, crid from ad_pool_app_info limit 10;")
	// dataFetcher.CloseConnection()
	// if err != nil {
	// 	log.Fatalf("Error fetching data: %v", err)
	// }

	// log.Printf("%+v\n", data)

	// // Write data to CSV
	// csvWriter := csv.NewCSVWriter("")
	// err = csvWriter.WriteCSV(data)
	// if err != nil {
	// 	log.Fatalf("Error writing CSV: %v", err)
	// }

	// // Upload CSV to S3
	// s3Uploader := s3.NewS3Uploader(cfg.S3Config)
	// presignedURL, err := s3Uploader.UploadFileAndGeneratePresignedURL(" ")

	// log.Println(presignedURL)

	// if err != nil {
	// 	log.Fatalf("Error uploading CSV to S3: %v", err)
	// }

	// // Send email notification
	// emailSender := email.NewEmailSender(cfg.EmailConfig)
	// err = emailSender.SendEmail()
	// if err != nil {
	// 	log.Fatalf("Error sending email: %v", err)
	// }
	utils.ExitProgram(false)
}

func logData(projectVersion string, currentProfile string, args []string, reportRequest string) {
	utils.LogDetails(projectVersion, "Project Version", false)
	utils.LogDetails(currentProfile, "Current Profile", false)
	utils.LogDetails(args, "Program Arguments", false)
	utils.LogDetails(reportRequest, "Report Request", false)
}
