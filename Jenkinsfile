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
        stage('Build Docker Image') {
            steps {
                script {
                    def dockerImageName = 'devops-project-app'
                    sh "docker build -t ${dockerImageName} -f ./user/Dockerfile ."
                }
            }
        }
        stage('Push to DockerHub') {
            environment {
                DOCKERHUB_CREDENTIALS = credentials('maximehrt-dockerhub')
            }
            steps {
                script {
                    def dockerImagePath = 'maximehrt/devops-project-app'
                    def dockerImageTag = 'latest'
                    sh "echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin"
                    sh "docker tag devops-project-app ${dockerImagePath}:${dockerImageTag}"
                    sh "docker push ${dockerImagePath}:${dockerImageTag}"
                    sh "docker logout"
                }
            }
        }
    }
}
