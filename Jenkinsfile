pipeline {
    agent any
    stages {
        stage("build") {
            steps {
                sh "echo 'je build ton image'"
                sh "docker build -t etl-${env.BUILD_ID} ."
            }
        }
        stage("run") {
            steps {
                sh "echo 'je lance le process'"
                sh "docker run etl-${env.BUILD_ID}"
            }
        }
    }
}