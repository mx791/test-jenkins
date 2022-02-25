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
                sh "docker run -v ./datas://usr/src/app/datas etl-${env.BUILD_ID}"
            }
            post {
                always {
                    publishHTML (target : [allowMissing: false,
                        alwaysLinkToLastBuild: true,
                        keepAll: true,
                        reportDir: 'datas',
                        reportFiles: 'datas.html',
                        reportName: 'My Reports',
                        reportTitles: 'The Report'] 
                    )
                }

            }
        }
    }
}