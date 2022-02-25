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
                sh "mkdir datas || true"
                sh "docker run etl-${env.BUILD_ID} > datas/data.html"
            }
            post {
                always {
                    publishHTML (target : [allowMissing: false,
                        alwaysLinkToLastBuild: true,
                        keepAll: true,
                        reportDir: 'datas',
                        reportFiles: 'data.html',
                        reportName: 'My Reports',
                        reportTitles: 'The Report'] 
                    )
                }
            }
        }
    }
}