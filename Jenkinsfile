pipeline {
    agent any
    stages {
        stage('List files') {
            steps {
                sh '''
                cd user
                ls -la
                '''
            }
        }
        stage('Login to DockerHub') {
            environment {
                DOCKERHUB_CREDENTIALS = credentials('maximehrt-dockerhub')
            }
            steps {
                sh "echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin"
            }
        }
        stage('Setup Docker Buildx') {
            steps {
                // Créer et utiliser un nouvel environnement Buildx
                sh "docker buildx create --name mybuilder --use"
                // Initialiser l'environnement Buildx
                sh "docker buildx inspect mybuilder --bootstrap"
            }
        }
        stage('Build and Push Docker Image') {
            steps {
                script {
                    def dockerImagePath = 'maximehrt/devops-project-app'
                    def dockerImageTag = 'latest'
                    // Construire et pousser l'image avec Buildx
                    sh "docker buildx build --platform linux/amd64,linux/arm64/v8 -t ${dockerImagePath}:${dockerImageTag} -f ./user/Dockerfile . --push"
                }
            }
        }
        stage('Logout from DockerHub') {
            steps {
                sh "docker logout"
            }
        }
    }
}
