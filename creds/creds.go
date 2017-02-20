package creds

import (
	"os"
)

func GetDbHost() string {
	return getEnvDefault("DB_HOST", "")
}

func GetDbUserName() string {
	return getEnvDefault("DB_USERNAME", "")
}

func GetDbPassword() string {
	return getEnvDefault("DB_PASSWORD", "")
}

func GetPort() string {
	return getEnvDefault("PORT", "")
}

func GetPreconfiguredPlans() string {
	return getEnvDefault("PRECONFIGURED_PLANS", `[
        {
          "service_id": "b9e4332e-b42b-4680-bda5-ea1506797474",
          "name": "standard",
          "display_name": "Standard",
          "description": "Standard storage class",
          "features": {"storage_class": "STANDARD"}
        },
        {
          "service_id": "b9e4332e-b42b-4680-bda5-ea1506797474",
          "name": "nearline",
          "display_name": "Nearline",
          "description": "Nearline storage class",
          "features": {"storage_class": "NEARLINE"}
        },
        {
          "service_id": "b9e4332e-b42b-4680-bda5-ea1506797474",
          "name": "reduced_availability",
          "display_name": "Durable Reduced Availability",
          "description": "Durable Reduced Availability storage class",
          "features": {"storage_class": "DURABLE_REDUCED_AVAILABILITY"}
        },
        {
          "service_id": "628629e3-79f5-4255-b981-d14c6c7856be",
          "name": "default",
          "display_name": "Default",
          "description": "PubSub Default plan",
          "features": ""
        },
        { "service_id": "f80c0a3e-bd4d-4809-a900-b4e33a6450f1",
          "name": "default",
          "display_name": "Default",
          "description": "BigQuery default plan",
          "features": ""
        },
        {
          "service_id": "5ad2dce0-51f7-4ede-8b46-293d6df1e8d4",
          "name": "default",
          "display_name": "Default",
          "description": "Machine Learning api default plan",
          "features": ""
        }
      ]`)
}

func GetRootCreds() string {
	return getEnvDefault("ROOT_SERVICE_ACCOUNT_JSON", ``)
}

func GetSecurityUserName() string {
	return getEnvDefault("SECURITY_USER_NAME", "")
}

func GetSecurityUserPassword() string {
	return getEnvDefault("SECURITY_USER_PASSWORD", "")
}

func GetServices() string {
	return getEnvDefault("SERVICES", `[
        {
          "id": "b9e4332e-b42b-4680-bda5-ea1506797474",
          "description": "A Powerful, Simple and Cost Effective Object Storage Service",
          "name": "google-storage",
          "bindable": true,
          "plan_updateable": false,
          "metadata": {
            "displayName": "Google Cloud Storage",
            "longDescription": "A Powerful, Simple and Cost Effective Object Storage Service",
            "documentationUrl": "https://cloud.google.com/storage/docs/overview",
            "supportUrl": "https://cloud.google.com/support/",
            "imageUrl": "https://cloud.google.com/_static/images/cloud/products/logos/svg/storage.svg"
          },
          "tags": ["gcp", "storage"]
        },
        {
          "id": "628629e3-79f5-4255-b981-d14c6c7856be",
          "description": "A global service for real-time and reliable messaging and streaming data",
          "name": "google-pubsub",
          "bindable": true,
          "plan_updateable": false,
          "metadata": {
            "displayName": "Google PubSub",
            "longDescription": "A global service for real-time and reliable messaging and streaming data",
            "documentationUrl": "https://cloud.google.com/pubsub/docs/",
            "supportUrl": "https://cloud.google.com/support/",
            "imageUrl": "https://cloud.google.com/_static/images/cloud/products/logos/svg/pubsub.svg"
          },
          "tags": ["gcp", "pubsub"]
        },
        {
          "id": "f80c0a3e-bd4d-4809-a900-b4e33a6450f1",
          "description": "A fast, economical and fully managed data warehouse for large-scale data analytics",
          "name": "google-bigquery",
          "bindable": true,
          "plan_updateable": false,
          "metadata": {
            "displayName": "Google BigQuery",
            "longDescription": "A fast, economical and fully managed data warehouse for large-scale data analytics",
            "documentationUrl": "https://cloud.google.com/bigquery/docs/",
            "supportUrl": "https://cloud.google.com/support/",
            "imageUrl": "https://cloud.google.com/_static/images/cloud/products/logos/svg/bigquery.svg"
          },
          "tags": ["gcp", "bigquery"]
        },
        {
          "id": "4bc59b9a-8520-409f-85da-1c7552315863",
          "description": "Google Cloud SQL is a fully-managed MySQL database service",
          "name": "google-cloudsql",
          "bindable": true,
          "plan_updateable": false,
          "metadata": {
            "displayName": "Google CloudSQL",
            "longDescription": "Google Cloud SQL is a fully-managed MySQL database service",
            "documentationUrl": "https://cloud.google.com/sql/docs/",
            "supportUrl": "https://cloud.google.com/support/",
            "imageUrl": "https://cloud.google.com/_static/images/cloud/products/logos/svg/sql.svg"
          },
          "tags": ["gcp", "cloudsql"]
        },
        {
          "id": "5ad2dce0-51f7-4ede-8b46-293d6df1e8d4",
          "description": "Machine Learning Apis including Vision, Translate, Speech, and Natural Language",
          "name": "google-ml-apis",
          "bindable": true,
          "plan_updateable": false,
          "metadata": {
            "displayName": "Google Machine Learning APIs",
            "longDescription": "Machine Learning Apis including Vision, Translate, Speech, and Natural Language",
            "documentationUrl": "https://cloud.google.com/ml/",
            "supportUrl": "https://cloud.google.com/support/",
            "imageUrl": "https://cloud.google.com/_static/images/cloud/products/logos/svg/machine-learning.svg"
          },
          "tags": ["gcp", "ml"]
        }
      ]`)
}

func getEnvDefault(key string, def string) string {
	if e, ok := os.LookupEnv(key); ok {
		return e
	}
	return def
}
