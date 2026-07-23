# 🚀 Cloud-Native DevOps Platform on AWS

<p align="center">

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker)
![Kubernetes](https://img.shields.io/badge/Kubernetes-kubeadm-326CE5?style=for-the-badge&logo=kubernetes)
![Jenkins](https://img.shields.io/badge/Jenkins-CI%2FCD-D24939?style=for-the-badge&logo=jenkins)
![Terraform](https://img.shields.io/badge/Terraform-IaC-844FBA?style=for-the-badge&logo=terraform)
![Ansible](https://img.shields.io/badge/Ansible-Automation-EE0000?style=for-the-badge&logo=ansible)
![Prometheus](https://img.shields.io/badge/Prometheus-Monitoring-E6522C?style=for-the-badge&logo=prometheus)
![Grafana](https://img.shields.io/badge/Grafana-Dashboard-F46800?style=for-the-badge&logo=grafana)
![ElasticSearch](https://img.shields.io/badge/Elastic-Logging-005571?style=for-the-badge&logo=elasticsearch)
![Jaeger](https://img.shields.io/badge/Jaeger-Tracing-66CFE3?style=for-the-badge)

</p>

---

## 📖 Project Overview

This project demonstrates a complete **Cloud-Native DevOps platform** built on **AWS** using Infrastructure as Code (Terraform), Configuration Management (Ansible), CI/CD Automation (Jenkins), Container Orchestration (Kubernetes), and a full Observability stack.

The platform automatically provisions AWS infrastructure, configures Kubernetes nodes, deploys a Go-based web application through Helm, and provides end-to-end monitoring, centralized logging, and distributed tracing.

The entire deployment follows DevOps best practices and represents a production-style workflow suitable for modern cloud-native applications.

---

# 🏗 Architecture

<p align="center">
<img src="docs/images/architecture-diagram.png" width="100%">
</p>

---

# ✨ Key Features

- 🚀 Infrastructure provisioning using Terraform
- ⚙️ Automated server configuration with Ansible
- ☸️ Multi-node Kubernetes Cluster (kubeadm)
- 🐳 Containerized Go Application
- 🔄 Jenkins CI/CD Pipeline
- 📦 Docker Hub Image Registry
- 🚢 Helm-based Kubernetes Deployment
- 🌐 NGINX Ingress Controller
- 📈 Prometheus Monitoring
- 📊 Grafana Dashboards
- 📝 Centralized Logging with Fluent Bit + Elasticsearch + Kibana
- 🔍 Distributed Tracing using OpenTelemetry + Jaeger
- ☁️ Hosted on AWS EC2
