pipeline {
    agent any

    options {
        ansiColor('xterm')
        timestamps()
        buildDiscarder(logRotator(numToKeepStr: '10'))
        disableConcurrentBuilds()
    }

    environment {
        APP_NAME        = "devops-dashboard"

        DOCKER_USERNAME = "navneet2004"

        IMAGE_NAME      = "${DOCKER_USERNAME}/devops-dashboard"

        IMAGE_TAG       = "${BUILD_NUMBER}"

        LOCAL_IMAGE     = "${APP_NAME}:${IMAGE_TAG}"

        REMOTE_IMAGE    = "${IMAGE_NAME}:${IMAGE_TAG}"

        LATEST_IMAGE    = "${IMAGE_NAME}:latest"
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

                    echo "===== Latest Commit ====="

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

                    echo "===== Git ====="

                    git --version
                '''
            }
        }

        stage('Verify Kubernetes Access') {
            steps {
                sh '''
                    echo "===== Kubernetes Cluster ====="

                    kubectl get nodes
                '''
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '''
                    echo "===== Building Docker Image ====="

                    docker build \
                        -t ${LOCAL_IMAGE} \
                        .
                '''
            }
        }

        stage('Scan Docker Image (Trivy)') {
            steps {
                sh '''
                    echo "===== Scanning Docker Image ====="

                    trivy image \
                        --scanners vuln \
                        --severity HIGH,CRITICAL \
                        --exit-code 1 \
                        ${LOCAL_IMAGE}
                '''
            }
        }

        stage('Authenticate Docker Hub') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {

                    sh '''
                        echo "===== Docker Login ====="

                        echo "$DOCKER_PASS" | docker login \
                            -u "$DOCKER_USER" \
                            --password-stdin
                    '''
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                sh '''
                    echo "===== Tagging Docker Images ====="

                    docker tag ${LOCAL_IMAGE} ${REMOTE_IMAGE}

                    docker tag ${LOCAL_IMAGE} ${LATEST_IMAGE}

                    echo

                    echo "===== Pushing Build Image ====="

                    docker push ${REMOTE_IMAGE}

                    echo

                    echo "===== Pushing Latest Image ====="

                    docker push ${LATEST_IMAGE}

                    docker logout
                '''
            }
        }
    }

    post {

        always {
            cleanWs()
        }

        success {
            echo "======================================="
            echo "CI Pipeline completed successfully."
            echo "======================================="
        }

        failure {
            echo "======================================="
            echo "CI Pipeline failed."
            echo "======================================="
        }
    }
}
