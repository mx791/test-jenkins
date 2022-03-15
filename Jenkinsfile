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
                        reportName: 'Weather Report',
                        reportTitles: 'The Report'] 
                    )
                    sh "docker image rm --force etl-${env.BUILD_ID}"
                }
            }
        }
        stage("Envoie du rapport") {
            steps {
                emailext (
                    subject: "JENKINS : résultat du build ${env.JOB_NAME} - ${env.BUILD_NUMBER}",
                    body: '${FILE,path="datas/data.html"}',
                    // body: "Bon bah ça a fonctionné",
                    to: "charles.laurioz@gmail.com",
                    mimeType: 'text/html'
                )
            }
        }
    }
}