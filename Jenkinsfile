pipeline {
    agent {
        docker {
            image 'docker:19.03'
            args '-v /var/run/docker.sock:/var/run/docker.sock'
        }
    }

    environment {
        DOCKER_USERNAME = credentials('docker-username')
        DOCKERHUB_TOKEN = credentials('dockerhub-token')
        IMAGE_NAME = "maximehrt/devops-project-app:latest"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('List files') {
            steps {
                sh '''
                cd user
                ls -la
                '''
            }
        }
        stage('Login to Docker Hub') {
            steps {
                script {
                    // Les étapes suivantes utiliseront cette authentification
                    docker.withRegistry('https://registry.hub.docker.com', 'dockerhub-credentials') {
                    }
                }
            }
        }
        stage('Build and Push Docker Image') {
            steps {
                script {
                    docker.build("${IMAGE_NAME}", "-f ./user/Dockerfile .").push()
                }
            }
        }
    }
}
