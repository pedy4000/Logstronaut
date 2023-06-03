# Getting Started

This guide will help you build the Logstronaut application and run it using Docker.

## Prerequisites
- GoLang
- Docker

## Build
Follow these steps to build the Logstronaut service:

1. Open a terminal window.
2. Navigate to the project root directory.
3. Run `make build`

This will create a binary executable for the Logstronaut service.

## Build Docker Image
After building the application, you can create a Docker image as follows:

1. Ensure Docker is running on your machine.
2. From the project root directory, run `docker build -t logstronaut .`

This will build a Docker image with the tag "logstronaut".

## Run Docker Container
To run a Docker container from the image you created:

1. From the terminal, run `docker run -p 80:8080 logstronaut`

This will start a Docker container, mapping the container's port 8000 to port 8000 on your host machine.
