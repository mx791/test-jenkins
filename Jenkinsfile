pipeline {
    agent any
    stages {
        stage("build") {
            steps {
                sh "docker build -t etl-${env.BUILD_ID} ."
            }
        }
        stage("build") {
            steps {
                sh "docker run etl-${env.BUILD_ID}"
            }
        }
    }
}