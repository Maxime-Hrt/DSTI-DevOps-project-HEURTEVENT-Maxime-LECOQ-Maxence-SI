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
