pipeline {
    agent any
    stages {
        stage("Build de l'image") {
            steps {
                sh "echo 'je build ton image'"
                sh "docker build --rm -t etl-${env.BUILD_ID} ."
            }
        }
        stage("Lancement du job") {
            steps {
                sh "echo 'je lance le process'"
                sh "mkdir -p datas"
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
                    emailext body: 'A Test EMail', recipientProviders: [[
                        $class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider'
                        ]], subject: 'Test'
                    }
                }
            }
        }
    }
}