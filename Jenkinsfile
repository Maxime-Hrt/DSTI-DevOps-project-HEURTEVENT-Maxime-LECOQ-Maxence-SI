pipeline {
    agent {
        docker {
            image 'docker:19.03'
            args '-v /var/run/docker.sock:/var/run/docker.sock'
        }
    }
    environment {
        DOCKERHUB_CREDENTIALS = credentials('maximehrt-dockerhub')
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
        stage('Build Docker Image') {
            steps {
                script {
                    def dockerImageName = 'maximehrt/devops-project-app'
                    def dockerImageTag = 'latest'
                    sh "docker build -t ${dockerImageName}:${dockerImageTag} -f ./user/Dockerfile ."
                }
            }
        }
        stage('Push to DockerHub') {
            environment {
                DOCKERHUB_CREDENTIALS = credentials('maximehrt-dockerhub')
            }
            steps {
                script {
                    def dockerImageName = 'maximehrt/devops-project-app'
                    def dockerImageTag = 'latest'
                    sh "echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin"
                    sh "docker tag devops-project-app ${dockerImageName}:${dockerImageTag}"
                    sh "docker push ${dockerImageName}:${dockerImageTag}"
                    sh "docker logout"
                }
            }
        }
    }
}
