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
        stage('Build and Push Docker Image') {
            steps {
                script {
                    def dockerImagePath = 'maximehrt/devops-project-app'
                    def dockerImageTag = 'latest'
                    // Utilisation de docker buildx pour construire et pousser l'image
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
