# Awwvision: Spring Boot edition

*Awwvision: Spring Boot edition* is a [Spring Boot](http://projects.spring.io/spring-boot/) demo application that uses the [Google Cloud Vision API](https://cloud.google.com/vision/) to classify (label) images from Reddit's [/r/aww](https://reddit.com/r/aww) subreddit, store the images and classifications in [Google Cloud Storage](https://cloud.google.com/storage/), and display the labeled results in a web app. It uses the GCP Service Broker to authenticate to the Vision and Storage APIs, and is based off of the [Python Awwvision sample app](https://github.com/GoogleCloudPlatform/cloud-vision/tree/master/python/awwvision).

This is not an official Google product.

Awwvision: Spring Boot edition has two endpoints:

1. A webapp that reads and displays the labels and associated images from GCS.
2. A scraper that downloads images from Reddit and classifies them using the Vision API.

## Prerequisites

1. Create a project in the [Google Cloud Platform Console](https://console.cloud.google.com).

1. [Enable billing](https://console.cloud.google.com/project/_/settings) for your project.

1. Enable the [Vision](https://console.cloud.google.com/apis/api/vision.googleapis.com) and [Storage](https://console.cloud.google.com/apis/api/storage_component) APIs. See the [Vision API Quickstart](https://cloud.google.com/vision/docs/quickstart) and [Storage API Quickstart](https://cloud.google.com/storage/docs/quickstart-console) for more information on using the two APIs.

## Build

The way we currently build the Docker image is kind of sad: we build the jar locally and then build the Docker image using that jar. First, to build the jar we run

```
mvn package -DskipTests
```

Then to create the Docker image we run

```
docker build . -t <tag:version>
```

## Deploy to GKE

Tag the previously built docker image to your GKE repo

```
docker tag <tag:version> us.gcr.io/<project_id>/<tag:version>
```

then push it with

```
gcloud docker -- push us.gcr.io/<project_id>/<tag:version>
```

Make sure you are on the correct cluster

```
kubectl config currenct-context
```

The easiest way to deploy is to use the awwvision.yaml found in the service-catalog repo because this app relies on secrets being exposed which are defined in that yaml file. To do that you would just run.

```
kubectl create -f awwvision.yaml
```

Alternatively you could run `kubectl run` manually and then edit the deployment to expose the secrets via `kubectl edit deployments/<deployment_name>`. You may have to restart the pods to get it to see the secrets.

Regardless of your method of deployment you need to expose it with

```
kubectl expose deployment awwvision --type=LoadBalancer --target-port=8080 --name=awwvision-lb
```

## Run the application on Cloud Foundry

1. Log in to your Cloud Foundry using the `cf login` command.

1. From the main project directory, build an executable jar and push it to Cloud Foundry. This step will initially fail due to lack of credentials.
    ```
    mvn package -DskipTests && cf push -p target/awwvision-spring-0.0.1-SNAPSHOT.jar awwvision
    ```

1. Create a Storage Bucket:
    ```
	cf create-service google-storage standard awwvision-storage -c '{"name": "awwvision-bucket"}'
    ```

    Make sure the name of the bucket matches the name specified as `gcp-storage-bucket` in [application.properties](./src/main/resources/application.properties). You can also have the service broker generate the bucket name for you by omitting the `-c` and everything after it, and modifying the code to pull the bucket name from the `VCAP_SERVICES` environment variable (see [VisionConfig.java](./src/main/java/com/google/cloud/servicebroker/awwvision/VisionConfig.java) for how this is done to parse the credentials.)

1. Bind the bucket to your app and give the service account storage object admin permissions:
    ```
    cf bind-service awwvision awwvision-storage -c '{"role":"storage.objectAdmin"}'
    ```

1. Restage the app so the new environment variables take effect:
    ```
    cf restage awwvision
    ```

### Visit the application and start the crawler

Once your application is running, visit awwvision.\[your-cf-instance-url\]/reddit to start crawling. The page will display "Scrape completed." once it is done. From there, visit awwvision.\[your-cf-instance-url\] to view your images!
