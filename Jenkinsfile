pipeline {
    agent any

    options {
        ansiColor('xterm')
        timestamps()
        buildDiscarder(logRotator(numToKeepStr: '10'))
        disableConcurrentBuilds()
    }

    environment {
        APP_NAME = 'devops-dashboard'
    }

    stages {

        stage('Checkout Source') {
            steps {
                checkout scm
            }
        }

        stage('Repository Information') {
            steps {
                sh '''
                    echo "===== Repository ====="
                    pwd
                    ls -lah
                    echo
                    echo "===== Git Commit ====="
                    git log --oneline -1
                '''
            }
        }

        stage('Verify Environment') {
            steps {
                sh '''
                    echo "===== Docker ====="
                    docker --version

                    echo
                    echo "===== Kubectl ====="
                    kubectl version --client

                    echo
                    echo "===== Helm ====="
                    helm version

                    echo
                    echo "===== Trivy ====="
                    trivy --version

                    echo
                    echo "===== Go ====="
                    go version

                    echo
                    echo "===== Git ====="
                    git --version
                '''
            }
        }

        stage('Verify Kubernetes Access') {
            steps {
                sh '''
                    kubectl get nodes
                '''
            }
        }
    }

    post {

        always {
            cleanWs()
        }

        success {
            echo 'Pipeline completed successfully.'
        }

        failure {
            echo 'Pipeline failed.'
        }
    }
}
