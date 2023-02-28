package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	Instance *Config
	once     sync.Once
)

func Initialize() *Config {
	if Instance == nil {
		once.Do(func() {
			if os.Getenv("MODE") == "" {
				err := godotenv.Load(getRootPath() + "/.env")
				if err != nil {
					log.Fatal("Error loading .env file")
				}
			}

			Instance = &Config{}
			Instance.App.ServiceName = os.Getenv("SERVICE_NAME")
			Instance.App.Mode = os.Getenv("MODE")
			Instance.App.Port = os.Getenv("PORT")

			Instance.Rds.MasterHostname = os.Getenv("RDS_MASTER_HOSTNAME")
			Instance.Rds.ReplicaHostname = os.Getenv("RDS_REPLICA_HOSTNAME")
			Instance.Rds.Username = os.Getenv("RDS_USERNAME")
			Instance.Rds.Password = os.Getenv("RDS_PASSWORD")
			Instance.Rds.Port = os.Getenv("RDS_PORT")
			Instance.Rds.DBName = os.Getenv("RDS_DB_NAME")

			Instance.Aws.DefaultRegion = os.Getenv("AWS_DEFAULT_REGION")
			Instance.Aws.AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
			Instance.Aws.SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
			Instance.Aws.Sqs.URL = os.Getenv("SQS_URL")
			Instance.Aws.Sqs.KakaoAlimtalkQueue = os.Getenv("SQS_KAKAO_ALIMTALK_QUEUE")

			Instance.Grpc.LivekitApiURL = os.Getenv("LIVEKIT_API_SERVICE")
			Instance.Grpc.MarketingApiURL = os.Getenv("MARKETING_API_SERVICE")

			Instance.ThirdPartyLib.Slack.HookURL = os.Getenv("SLACK_ALRAM_URL")
			Instance.ThirdPartyLib.Slack.Channel = os.Getenv("SLACK_CHANNEL")
			Instance.ThirdPartyLib.Sentry.DSN = os.Getenv("SENTRY_DSN")
			Instance.ThirdPartyLib.Kafka.Brokers = os.Getenv("KAFKA_BROKERS")
			Instance.ThirdPartyLib.Infobank.SenderKey = os.Getenv("INFOBANK_SENDER_KEY")
		})
	}

	return Instance
}

func getRootPath() string {
	var _, b, _, _ = runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../..")
}
