#!/bin/bash

set -eux
# gcloud login authしておく
gcloud config set project $PROJECT_ID
gcloud services enable container.googleapis.com containerregistry.googleapis.com \
    servicenetworking.googleapis.com cloudresourcemanager.googleapis.com \
    sqladmin.googleapis.com iamcredentials.googleapis.com cloudbuild.googleapis.com

gcloud iam service-accounts create terraform --display-name="terraform-account" --project=$PROJECT_ID

gcloud iam service-accounts keys create $KEY_SAVE_PATH --iam-account terraform@$PROJECT_ID.iam.gserviceaccount.com
gcloud projects add-iam-policy-binding $PROJECT_ID --member serviceAccount:terraform@$PROJECT_ID.iam.gserviceaccount.com --role roles/compute.admin
gcloud projects add-iam-policy-binding $PROJECT_ID --member serviceAccount:terraform@$PROJECT_ID.iam.gserviceaccount.com --role roles/storage.admin
gcloud projects add-iam-policy-binding $PROJECT_ID --member serviceAccount:terraform@$PROJECT_ID.iam.gserviceaccount.com --role roles/container.clusterAdmin
gcloud projects add-iam-policy-binding $PROJECT_ID --member serviceAccount:terraform@$PROJECT_ID.iam.gserviceaccount.com --role roles/iam.serviceAccountAdmin
gcloud projects add-iam-policy-binding $PROJECT_ID --member serviceAccount:terraform@$PROJECT_ID.iam.gserviceaccount.com --role roles/iam.serviceAccountUser
gcloud projects add-iam-policy-binding $PROJECT_ID --member serviceAccount:terraform@$PROJECT_ID.iam.gserviceaccount.com --role roles/servicenetworking.networksAdmin
gcloud projects add-iam-policy-binding $PROJECT_ID --member serviceAccount:terraform@$PROJECT_ID.iam.gserviceaccount.com --role roles/cloudsql.admin
gcloud projects add-iam-policy-binding $PROJECT_ID --member serviceAccount:terraform@$PROJECT_ID.iam.gserviceaccount.com --role roles/dns.admin
gcloud projects add-iam-policy-binding $PROJECT_ID --member serviceAccount:terraform@$PROJECT_ID.iam.gserviceaccount.com --role roles/iam.securityAdmin
