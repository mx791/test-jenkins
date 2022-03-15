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
                }
            }
        }
        stage("Envoie du rapport") {
            emailext (
                subject: "STARTED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
                body: """Bon bah ça a fonctionné""",
                recipientProviders: [[$class: 'DevelopersRecipientProvider']],
                to: "charles.laurioz@gmail.com"
            )
        }
    }
}