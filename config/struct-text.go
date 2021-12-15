package config

const (
	IGoreConfigStructText = `
	type IGoreConfig struct {
		Username    string json:"username"
		Password    string json:"password"
		Host        string json:"host"
		Port        string json:"port"
		Subdatabase string json:"subdatabase"
	
		S3_REGION string json:"S3_REGION"
		S3_BUCKET string json:"S3_BUCKET"
	}
	`
)
