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

        HELM_RELEASE    = "devops-dashboard"
        K8S_NAMESPACE   = "devops-dashboard"
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
                    echo "========================================="
                    echo "Repository Information"
                    echo "========================================="

                    pwd
                    ls -lah

                    echo
                    git log --oneline -1
                '''
            }
        }

        stage('Verify Environment') {
            steps {
                sh '''
                    echo "========================================="
                    echo "Verify Environment"
                    echo "========================================="

                    docker --version
                    kubectl version --client
                    helm version
                    trivy --version
                    git --version
                '''
            }
        }

        stage('Verify Kubernetes Access') {
            steps {
                sh '''
                    echo "========================================="
                    echo "Kubernetes Cluster"
                    echo "========================================="

                    kubectl get nodes
                '''
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '''
                    echo "========================================="
                    echo "Build Docker Image"
                    echo "========================================="

                    docker build -t ${LOCAL_IMAGE} .
                '''
            }
        }

        stage('Scan Docker Image (Trivy)') {
            steps {
                sh '''
                    echo "========================================="
                    echo "Trivy Scan"
                    echo "========================================="

                    trivy image \
                      --scanners vuln \
                      --severity HIGH,CRITICAL \
                      --exit-code 0 \
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
                    echo "========================================="
                    echo "Push Docker Images"
                    echo "========================================="

                    docker tag ${LOCAL_IMAGE} ${REMOTE_IMAGE}
                    docker tag ${LOCAL_IMAGE} ${LATEST_IMAGE}

                    docker push ${REMOTE_IMAGE}
                    docker push ${LATEST_IMAGE}

                    docker logout
                '''
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                sh '''
                    echo "========================================="
                    echo "Deploying with Helm"
                    echo "========================================="

                    helm upgrade --install ${HELM_RELEASE} ./helm \
                      --namespace ${K8S_NAMESPACE} \
                      --create-namespace \
                      --set image.repository=${IMAGE_NAME} \
                      --set image.tag=${BUILD_NUMBER}
                '''
            }
        }

        stage('Wait For Rollout') {
            steps {
                sh '''
                    echo "========================================="
                    echo "Waiting for Rollout"
                    echo "========================================="

                    kubectl rollout status deployment/${HELM_RELEASE} \
                      -n ${K8S_NAMESPACE} \
                      --timeout=300s
                '''
            }
        }

        stage('Verify Deployment') {
            steps {
                sh '''
                    kubectl get deployment -n ${K8S_NAMESPACE}
                    kubectl get pods -o wide -n ${K8S_NAMESPACE}
                '''
            }
        }

        stage('Verify Service') {
            steps {
                sh '''
                    kubectl get svc -n ${K8S_NAMESPACE}
                '''
            }
        }

        stage('Verify Ingress') {
            steps {
                sh '''
                    kubectl get ingress -n ${K8S_NAMESPACE}
                '''
            }
        }

        stage('Health Check') {
            steps {
                sh '''
                    kubectl get all -n ${K8S_NAMESPACE}
                '''
            }
        }
    }

    post {

        success {
            echo "========================================="
            echo "CI/CD Pipeline completed successfully."
            echo "========================================="
        }

        failure {
            echo "========================================="
            echo "Pipeline failed."
            echo "========================================="
        }

        always {
            cleanWs()
        }
    }
}
