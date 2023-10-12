package main

import (
	"log"
	"os"

	"ssp-portal-reporting-processor/config"
	// "ssp-portal-reporting-processor/csv"
	// "ssp-portal-reporting-processor/email"
	// "ssp-portal-reporting-processor/s3"
)

func main() {

	args := os.Args
	currentProfile := os.Getenv("PROFILE")
	projectVersion := os.Getenv("PROJECT_VERSION")
	logDetails(currentProfile, "Current Profile", false)
	logDetails(args, "Program Arguments", false)
	logDetails(projectVersion, "Project Version", false)
	// Print version also
	// Load environment configuration
	env := "dev" // Change this based on the environment
	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		exitProgram(true)
	}

	log.Printf("%+v\n", cfg)
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
	exitProgram(false)
}

func logDetails(val any, title string, isPanic bool) {
	if isPanic {
		log.Panicf("%v ::\n%v\n", title, val)
	} else {
		log.Printf("%v ::\n%v\n", title, val)
	}
}

func exitProgram(failed bool) {
	if failed {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
