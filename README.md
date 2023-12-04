# DSTI-DevOps-project-HEURTEVENT-Maxime-SI

## Summary
- [Installation](#installation)
- [Create a web application](#create-a-web-application)
- [Apply CI/CD pipeline](#apply-cicd-pipeline)
- [Configure and provision a virtual environment (Vagrant)](#configure-and-provision-a-virtual-environment-and-run-your-application-using-the-iac-approach)
- [Build Docker image and push it to Docker Hub](#build-docker-image-and-push-it-to-docker-hub)
- [Make container orchestration using Docker Composer](#make-container-orchestration-using-docker-composer)
- [Make docker orchestration using Kubernetes](#make-docker-orchestration-using-kubernetes)
- [Make a service mesh using Istio](#make-a-service-mesh-using-istio)
- [Implement Monitoring to your containerized application](#implement-monitoring-to-your-containerized-application)
- [Project Documentation](#project-documentation)

| Subject                                                         | Code | Max. grade | Status |
|:----------------------------------------------------------------|:----:|:----------:|--------|
| Enriched web application with automated tests                   | APP  |     +1     | ✅      |
| Continuous Integration and Continuous Delivery (and Deployment) | CICD |     +3     |        |
| Containerisation with Docker                                    |  D   |     +1     | ✅      |
| Orchestration with Docker Compose                               |  DC  |     +2     | ✅      |
| Orchestration with Kubernetes                                   | KUB  |     +3     |        |
| Service mesh using Istio                                        | IST  |     +2     |        |
| Infrastructure as code using Ansible                            | IAC  |     +3     | ✅      |
| Monitoring                                                      | MON  |     +2     |        |
| Accurate project documentation in README.md file                | DOC  |     +3     | ✅      |
| Each bonus task                                                 | BNS  |     +1     | ✅      |
| Each penalty                                                    | PNL  |     -1     |        |

| Bonus               |
|---------------------|
| Golang CRUD App     |
| Swagger             |
| Static Front        |
| Docker Hub pipeline |
| Docker Hub Overview |



## Installation
Start by cloning the repository:
```shell
git clone git@github.com:Maxime-Hrt/DSTI-DevOps-project-HEURTEVENT-Maxime-LECOQ-Maxence-SI.git
```
Then, move to the project directory:
```shell
cd DSTI-DevOps-project-HEURTEVENT-Maxime-LECOQ-Maxence-SI
```

## 1. Create a web application
### Prerequisites
The web application is a basic CRUD application that uses Redis as a database. It is written in Go and uses the
[Fiber](https://gofiber.io/) framework. To run the application, you will need to install:
- [Go](https://golang.org/doc/install)
- [Redis](https://redis.io/download)

### Run the application
To run the application, you will need to start Redis and then run the application. To do so, you can run the following commands:
```shell
# Start Redis
redis-server
```
```shell
# Move to the user directory
cd user
```
If you are running the application for the first time, you will need to install the dependencies:
```shell
# Install the dependencies
go get .
```
```shell
# Run the application
go build && ./devops-project
```
You can now access the application with the [Swagger UI](http://localhost:8080/swagger/index.html#/), the **static frontview** or simply verify it's good running with the [health check](http://localhost:8080/health).

### <p style="text-align: center;">Swagger UI</p>

<p align="center">
    <img alt="Swagger UI" src="Images/webapp/swagger_ui.png" width="850" />
</p>

### <p style="text-align: center;">Static Front</p>


<p align="center">
    <img alt="Static Front" src="Images/webapp/static_front.png" width="850"/>
</p>

### Run the tests
To run the tests, you will need to start Redis and then run the tests. To do so, you can run the following commands:
```shell
# Start Redis
redis-server
```
```shell
# Move to the user/test directory
cd user/test
```
```shell
# Run the tests
go test -v
```

## 2. Apply CI/CD pipeline
### Prerequisites
To apply the CI/CD pipeline, you will need to install:
- [Docker](https://docs.docker.com/get-docker/)
### Run the pipelines
Two pipelines are available using GitHub Actions. 
The first one described in `go.yml` is used to build the application and **run the tests** on each push.
The second one described in `dockerhub.yml` is used to **build the Docker image and push** it to Docker Hub on each release.

### <p style="text-align: center;"> Go Test Actions</p>
<p align="center">
    <img alt="Go Test Actions" src="Images/ci_cd/test_go.png" width="850"/>
</p>

### Deploy the application on Azure

## 3. Configure and provision a virtual environment (Vagrant)
### Prerequisites
To configure and provision a virtual environment, you will need to install:
- [VirtualBox](https://www.virtualbox.org/wiki/Downloads)
- [Vagrant](https://developer.hashicorp.com/vagrant/downloads)
- [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)

Install vagrant Accoding to the documentation for the **VirtualBox** provider.
### Configure with Vagrant
To build the virtual environment with Vagrant and Ansible, you can run the following command from the root directory:
```shell
# Run the Vagrant configuration
vagrant up
```

<p align="center">
    <img alt="vagrant up" src="Images/virtual_env/status.png" width="auto"/>
</p>

Verify that the virtual machine is running --> [Swagger](http://localhost:3000/swagger/index.html#/)

<p align="center">
    <img alt="vagrant up" src="Images/virtual_env/swagger_vagrant.png" width="850"/>
</p>

## 4. Build Docker image and push it to Docker Hub
### Prerequisites
To build the Docker image, you will need to install:
- [Docker](https://docs.docker.com/get-docker/)

Our Docker image is available on [Docker Hub](https://hub.docker.com/r/maximehrt/devops-project-image/tags)
you can pull it using the following command:
```shell
# Pull the Docker image
docker pull maximehrt/devops-project
```

### Build the Docker image
To build the Docker image, you can run the following command:
```shell
# Move to the user directory
cd user
```
```shell
# Build the Docker image
docker build -t devops-project-app .
```
You can now run the Docker image:
```shell
# Run the Docker image
docker run -p 8080:8080 devops-project-app
```
### Push the Docker image to Docker Hub
To push the Docker image to Docker Hub, you will need to login to Docker Hub:
```shell
# Login to Docker Hub
docker login
```
Tag the Docker image:
```shell
# Tag the Docker image
docker tag devops-project-app:latest maximehrt/devops-project-app:latest
```
Then, you can push the Docker image to Docker Hub:
```shell
# Push the Docker image to Docker Hub
docker push maximehrt/devops-project-app:latest
```
<p align="center">
    <img alt="docker_hub" src="Images/docker_img_hub/dockerhub.png" width="850"/>
</p>

## 5. Make container orchestration using Docker Compose
From the root directory, run the following command:
```shell
# Run the Docker Compose
docker-compose up
```
You can now access the application with the [Swagger UI](http://localhost:8080/swagger/index.html#/) or verify it's good running with the [health check](http://localhost:8080/health).

<p align="center">
    <img alt="docker desktop running" src="Images/docker_comp/desktop.png" width="850"/>
</p>

## 6. Make docker orchestration using Kubernetes

## 7. Make a service mesh using Istio

## 8. Implement Monitoring to your containerized application

## 9. Project Documentation


