package config

type Config struct {
	App struct {
		ServiceName string
		Mode        string
		Port        string
	}

	Rds struct {
		MasterHostname  string
		ReplicaHostname string
		Username        string
		Password        string
		DBName          string
		Port            string
	}

	Aws struct {
		DefaultRegion   string
		AccessKeyID     string
		SecretAccessKey string

		Sqs struct {
			URL                string
			KakaoAlimtalkQueue string
		}
	}

	Grpc struct {
		LivekitApiURL   string
		MarketingApiURL string
	}

	ThirdPartyLib struct {
		Slack struct {
			HookURL string
			Channel string
		}

		Sentry struct {
			DSN string
		}

		Kafka struct {
			Brokers string
		}

		Infobank struct {
			SenderKey string
		}
	}
}
