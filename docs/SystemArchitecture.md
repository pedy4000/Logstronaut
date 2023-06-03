# System Architecture

This document provides an overview of the system architecture for the Logstronaut project.

## DigitalOcean Cluster

The Logstronaut system is designed to run on a DigitalOcean cluster, ensuring scalable and reliable operation. The cluster consists of:

- **Three worker nodes**: Each armed with a single CPU and 2GB of RAM.
- **One high-capacity worker node**: Equipped with 4 CPUs and 8GB of RAM.

The entire cluster is strategically located in Amsterdam to optimize data handling and latency management.

![lenz workspace screenshot](./src/lenz_workspace.png?raw=true "Workspace Screenshot")

## CI/CD Pipeline

In order to maintain code quality and ensure quick and reliable deployments, we've established a CI/CD pipeline that uses **GitHub Actions** and **ArgoCD**. The flow of operations is as follows:

1. **[GitHub Actions](../.github/workflows/docker-image.yml)**: Handles the building and testing of the Docker image.
2. **DigitalOcean Registry**: A secure space where our Docker image is stored post-build.
3. **ArgoCD**: Monitors the Helm charts in our repository and synchronizes them with the Kubernetes clusters.

## ArgoCD Applications

Our ArgoCD environment hosts several applications, each serving different roles in the system:

![argocd applications screenshot](./src/argocd_applications.png?raw=true "argocd Screenshot")

### Logstronaut

This is the core application of our project, which also encompasses the main service of the app and the database resources. The Logstronaut service is streamlined via a custom Helm chart. It incorporates a **Deployment resource** that manages the lifecycle of the main application pods.

![argocd logstronaut screenshot](./src/argocd_logstronaut.png?raw=true "argocd logstronaut Screenshot")

### Logger Stack: Elasticsearch, Logstash, Kibana (ELK)

This application includes multiple stateful sets and pods to manage the system's logs. It leverages Elasticsearch for search, Logstash for centralized logging, and Kibana for visualization.

![argocd elk screenshot](./src/argocd_elk.png?raw=true "argocd elk Screenshot")

### Cert-manager

Cert-manager is responsible for the issuance of certificates. It employs Let's Encrypt to create a certificate issuer, thereby ensuring a secure, automated, and open certificate management process.

### External-Ingress

The External-Ingress application manages ingress resources for ArgoCD and Grafana. It allows both of these services to be accessed externally by routing from the specified domain.

## PostgreSQL Database

Our PostgreSQL database is set up to deliver peak performance. We've allocated 1 CPU, 1 GB of RAM, and a 10GB disk from DigitalOcean. In line with the principles of **Infrastructure as Code (IaC)**, our YAML files automate the creation of:

- **ConfigMap**: Houses the connection information.
- **Secret resources**: Stores the database administrator credentials.

This automated setup enables a seamless and secure database configuration process.
Captured below is the architecture of the Managed Database, deployed using the DigitalOcean Kubernetes Operator (do-operator).

![database architecture screenshot](./src/managed_database_architecture.jpeg?raw=true "database architecture Screenshot")

## Logstronaut Service Deployment

In particular, we have a Deployment resource that oversees the update process for the main application pods. Its operation includes:

- Initiating new pods with every update.
- Conducting a readiness check before marking the deployment as complete.
- Implementing a rollback to the previous stable version if a deployment error is encountered.

This ensures continuous service availability.

![logstronaut deployment screenshot](./src/logstronaut_deployment.png?raw=true "logstronaut deployment Screenshot")

## Helm Templates

We've utilized Helm templates for creating all the YAML files. This approach facilitates flexibility and reusability, aligning perfectly with our vision of a scalable and robust system architecture.
