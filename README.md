# Logstronaut

Welcome to the Logstronaut project, a task assignment for **InPress** company. Logstronaut, built with Golang, takes the simplicity of message-saving to the next level, storing them in a PostgreSQL database with robust Kubernetes infrastructure and Helm charts for a streamlined setup.

The application is designed to provide comprehensive monitoring (via Prometheus, Grafana, and Alert Manager), and extensive logging (through Elasticsearch, Logstash, and Kibana). A built-in cert-manager and argoCD config further enhance its functionality. What's more, it makes message transmission as straightforward as sending a POST request to [inpress.sadeghian.info](http://inpress.sadeghian.info) with a JSON body containing a "message" key.

## Contents

1. [Getting Started](docs/GettingStarted.md): A step-by-step guide to build and run the application.
2. [APIs Usage](docs/APIsUsage.md): Detailed Postman documentation for API usage.
3. [System Architecture](docs/SystemArchitecture.md): A comprehensive view into the structure of Logstronaut.
4. [Deployment Procedure](docs/DeploymentProcedure.md): Step-by-step walkthrough of the deployment process, including Argo and GitHub CI setup and volume creation.

Feel free to explore and contribute to our project!

## License
This project is under the [MIT License](LICENSE).

