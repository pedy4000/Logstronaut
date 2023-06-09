# Deployment Procedure

This document provides a step-by-step walkthrough of the deployment process, including Argo, GitHub CI setup, and volume creation. 

## Prerequisites

Before you begin, you must create a Kubernetes cluster in DigitalOcean. This is because certain manifests utilize the `do-operator` (DigitalOcean operator). 

## Installation and Deployment

Once the cluster is set, follow the steps below:

1. **Install essential components:** Use the pre-prepared YAML files from DigitalOcean to install the Kubernetes Monitoring Stack, NGINX Ingress Controller, Cert-Manager, and ArgoCD.

2. **Setup CI/CD Environment:** Obtain your `DIGITALOCEAN_ACCESS_TOKEN` from the DigitalOcean API section. Subsequently, create a secret with the same name in GitHub for the CI pipeline. In ArgoCD, establish a route to the Helm charts in your repository. This will enable ArgoCD to synchronize each chart in the repository with the cluster. *Note: Ensure to replace your domain in the logstronaut YAML values in the ArgoCD parameter part.*

3. **Configure Cert-Manager:** Update the `lets-encrypt-issuer.yaml` file with your email. Also, complete the TLS portion in every manifest you're deploying.

4. **Set Up Ingress:** Specify your desired path and domain in the values section of each manifest you're deploying.

5. **ELK Stack:** To run this setup without any hiccups, you'll need at least 7 nodes. By default, the PersistentVolumeClaims for each Elasticsearch Pod will be 30GB. The setup will automatically run and read logs from the app console.

6. **Monitoring Stack:** Create your personal dashboard in Grafana and set alerts as needed. Metrics from the node exporter, cluster data, and app metrics are readily available for querying and monitoring.

## Logstronaut Configuration Map

In setting up Logstronaut, it's essential to create and populate the `logstronaut-cm` (Configuration Map). This includes a variety of key-value pairs derived from the `app.env.example` file. One of the critical keys is the `DB_ADDRESS`, which can be obtained in two ways:

1. **DigitalOcean Managed Database Resource**: Opt for a fully managed database service by DigitalOcean. This choice allows you to offload your database management tasks while benefiting from an automatic setup process. The setup process seamlessly integrates a connection URI that includes both the username and password.

2. **Manual Setup**: If you prefer a hands-on approach, you could manually set up the storage and database. However, remember that the full DigitalOcean volume does not support ReadWriteMany (RWM). Regardless of the approach chosen, the `DB_ADDRESS` needs to be set to your database's address, inclusive of the username and password within the connection URI.

These guidelines will ensure a smooth deployment process, fostering a reliable and robust environment for Logstronaut.