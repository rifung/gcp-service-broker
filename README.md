# Cloud Foundry Service Broker for Google Cloud Platform

Depends on [lager](https://github.com/pivotal-golang/lager) and [gorilla/mux](https://github.com/gorilla/mux).

Requires Go 1.6 and the associated buildpack.

## Examples

See the [examples](https://github.com/GoogleCloudPlatform/gcp-service-broker/tree/master/examples/) folder.

## Prerequisites

### Set up a GCP Project

1. go to [Google Cloud Console](https://console.cloud.google.com) and sign up, walking through the setup wizard
1. next to the Google Cloud Platform logo in the upper left-hand corner, click the dropdown and select "Create Project"
1. give your project a name and click "Create"
1. when the project is created (a notification will show in the upper right), refresh the page.

### Enable APIS

1. Navigate to **API Manager > Library**.
1. Enable the <a href="https://console.cloud.google.com/apis/api/cloudresourcemanager.googleapis.com/overview">Google Cloud Resource Manager API</a>
1. Enable the <a href="https://console.cloud.google.com/apis/api/iam.googleapis.com/overview">Google Identity and Access Management (IAM) API</a>
1. If you want to enable Cloud SQL as a service, enable the <a href="https://console.cloud.google.com/apis/api/sqladmin/overview">Cloud SQL API</a>
1. If you want to enable Cloud Storage as a service, enable the <a href="https://console.cloud.google.com/apis/api/storage_component/overview">Cloud Storage API</a>

Rest of these are not necessary for awwvision demo

1. If you want to enable BigQuery as a service, enable the <a href="https://console.cloud.google.com/apis/api/bigquery/overview">BigQuery API</a>
1. If you want to enable Pub/Sub as a service, enable the <a href="https://console.cloud.google.com/apis/api/pubsub/overview">Cloud Pub/Sub API</a>
1. If you want to enable Bigtable as a service, enable the <a href="https://console.cloud.google.com/apis/api/bigtableadmin/overview">Bigtable Admin API</a>

### Create a root service account

1. From the GCP console, navigate to **IAM & Admin > Service accounts** and click **Create Service Account**.
1. Enter a **Service account name**.
1. Select the checkbox to **Furnish a new Private Key**, and then click **Create**.
1. Save the automatically downloaded key file to a secure location.
1. Navigate to **IAM & Admin > IAM** and locate your service account.
1. From the dropdown on the right, choose **Project > Owner** and click **Save**.

### Set up a backing database

1. create new MySQL instance (for example, using CloudSQL)
1. run `CREATE DATABASE servicebroker;`
1. run `CREATE USER '<username>'@'%' IDENTIFIED BY '<password>';`
1. run `GRANT ALL PRIVILEGES ON servicebroker.* TO '<username>'@'%' WITH GRANT OPTION;`
1. (optional) create SSL certs for the database and save them somewhere secure

If using CloudSQL you may have to set up permissions to allow incoming connections

### Set required constants

The PORT defaults to 8080. If you want to change it then update the Dockerfile (assuming you will build with Docker).

Fill in the constants in creds/creds.go. Note that you can use multiline string constants with backticks '\`' in Java.

* `ROOT_SERVICE_ACCOUNT_JSON` (the string version of the credentials file created for the Owner level Service Account)
* `SECURITY_USER_NAME` (a username to sign all service broker requests with - the same one used in cf create-service-broker)
* `SECURITY_USER_PASSWORD` (a password to sign all service broker requests with - the same one used in cf create-service-broker)
* `DB_HOST` (the host for the database to back the service broker)
* `DB_USERNAME` (the database username for the service broker to use)
* `DB_PASSWORD` (the database password for the service broker to use)

For example, if your DB username is servicebroker, then you would change

```
func GetDbUserName() string {
	return getEnvDefault("DB_USERNAME", "")
}
```

to

```
func GetDbUserName() string {
	return getEnvDefault("DB_USERNAME", "servicebroker")
}
```

Note that changing PORT doesn't do anything if you are building from Docker since the Dockerfile will set the environment variable which is used whenever present.

### optional env constants

* `DB_PORT` (defaults to 3306)
* `CA_CERT`
* `CLIENT_CERT`
* `CLIENT_KEY`
* `CLOUDSQL_CUSTOM_PLANS` (A map of plan names to string maps with fields `guid`, `name`, `description`, `tier`,
`pricing_plan`, `max_disk_size`, `display_name`, and `service` (Cloud SQL's service id)) - if unset, the service
will be disabled. e.g.,

```json
{
    "test_plan": {
        "name": "test_plan",
        "description": "testplan",
        "tier": "D8",
        "pricing_plan": "PER_USE",
        "max_disk_size": "15",
        "display_name": "FOOBAR",
        "service": "4bc59b9a-8520-409f-85da-1c7552315863"
    }
}
```


## Building

From the root directory run

`docker build . -t <tag:version>`

The version is optional but makes it easier to redeploy if you make a mistake and need to update

You can also build locally with `go build`

## Deploying

This assumes you are deploying to GKE, have already made your cluster, and are authorized to it.

Tag your previously made image to the Docker repo of your project via

```
docker tag <tag:version> us.gcr.io/<project_id>/<tag:version>
```

push the image to your Docker repo

```
gcloud docker -- push us.gcr.io/<project_id>/<tag:version>
```

Make sure your kubectl is configured to use the correct cluster

```
kubctl config current-context
```

If not then set it to the right one

```
gcloud container clusters get-credentials <cluster_name> --zone=<cluster_zone>
```

Now deploy it

```
kubectl run gcp-sb --image=us.gcr.io/<project_id>/<tag:version> --port=8080
```

Obviously if you changed the port, modify the command to use the port you exposed in the Dockerfile. You should be able to see the deployment and pod

```
kubectl get deployments
kubectl get pods
```

Now we need to expose the service

```
kubectl expose deployment/gcp-sb --type=LoadBalancer --target-port=8080 --name=gcp-sb-lb
```

Again, change the port to the one you chose to expose if it's something other than 8080.

You should be able to see the service running

```
kubectl get services
```

Make note of that external IP since you will need it when configuring the service catalog.

The GCP Service Broker should now be running! You can make sure it's working with

```
curl <external_ip>
```

which should return `Not Authorized`. You can see the catalog with

```
curl -u <user>:<password> <external_ip>/v2/catalog
```

where the user and password are the SECURITY_USER_NAME and SECURITY_USER_PASSWORD from creds.go

# This is not an official Google product.
