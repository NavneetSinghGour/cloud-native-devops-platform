# 🚀 Cloud-Native DevOps Platform on AWS

![AWS](https://img.shields.io/badge/AWS-EC2%20%7C%20VPC%20%7C%20IAM-FF9900?style=for-the-badge&logo=amazonaws&logoColor=white)
![Terraform](https://img.shields.io/badge/Terraform-IaC-7B42BC?style=for-the-badge&logo=terraform&logoColor=white)
![Ansible](https://img.shields.io/badge/Ansible-Automation-EE0000?style=for-the-badge&logo=ansible&logoColor=white)
![Jenkins](https://img.shields.io/badge/Jenkins-CI/CD-D24939?style=for-the-badge&logo=jenkins&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Containerization-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Kubernetes](https://img.shields.io/badge/Kubernetes-kubeadm-326CE5?style=for-the-badge&logo=kubernetes&logoColor=white)
![Helm](https://img.shields.io/badge/Helm-Packaging-0F1689?style=for-the-badge&logo=helm&logoColor=white)
![Go](https://img.shields.io/badge/Go-Backend-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Prometheus](https://img.shields.io/badge/Prometheus-Monitoring-E6522C?style=for-the-badge&logo=prometheus&logoColor=white)
![Grafana](https://img.shields.io/badge/Grafana-Dashboard-F46800?style=for-the-badge&logo=grafana&logoColor=white)
![Fluent Bit](https://img.shields.io/badge/Fluent%20Bit-Logging-49BDA5?style=for-the-badge)
![Elasticsearch](https://img.shields.io/badge/Elasticsearch-Search-005571?style=for-the-badge&logo=elasticsearch&logoColor=white)
![Kibana](https://img.shields.io/badge/Kibana-Visualization-005571?style=for-the-badge&logo=kibana&logoColor=white)
![OpenTelemetry](https://img.shields.io/badge/OpenTelemetry-Tracing-425CC7?style=for-the-badge)
![Jaeger](https://img.shields.io/badge/Jaeger-Distributed%20Tracing-66CFE3?style=for-the-badge)

---

> Enterprise-grade DevOps platform demonstrating Infrastructure as Code, Configuration Management, CI/CD Automation, Kubernetes Orchestration, and End-to-End Observability.

![Architecture](docs/images/architecture-overview.png)

---

# 📖 Project Overview

This project demonstrates the implementation of a production-style cloud-native DevOps platform on AWS. It showcases the complete software delivery lifecycle, starting from infrastructure provisioning and server configuration to continuous integration, continuous deployment, Kubernetes orchestration, and full-stack observability.

The infrastructure is provisioned using Terraform, configured with Ansible, and hosts a Kubernetes cluster built using kubeadm. A dedicated Jenkins server automates the CI/CD pipeline by building the Go application, creating Docker images, scanning them with Trivy, publishing them to Docker Hub, and deploying them to Kubernetes using Helm.

The platform also includes a complete observability stack with Prometheus, Grafana, Fluent Bit, Elasticsearch, Kibana, OpenTelemetry Collector, and Jaeger to provide monitoring, centralized logging, and distributed tracing.

---

# ✨ Key Features

## ☁️ Infrastructure Automation

- AWS Infrastructure Provisioning using Terraform
- Automated Configuration Management using Ansible
- Infrastructure as Code (IaC)
- Reproducible Infrastructure Deployment

---

## ⚙️ Continuous Integration & Delivery

- GitHub Webhook Integration
- Automated Jenkins Pipeline
- Go Application Build
- Docker Image Creation
- Trivy Security Scan
- Docker Hub Image Publishing
- Helm-based Kubernetes Deployment

---

## ☸️ Kubernetes Platform

- kubeadm-based Kubernetes Cluster
- NGINX Ingress Controller
- Rolling Updates
- Self-Healing Deployments
- Service Discovery
- Replica-based High Availability

---

## 📊 Observability

### Monitoring

- Prometheus Metrics Collection
- Grafana Dashboards

### Logging

- Fluent Bit
- Elasticsearch
- Kibana

### Distributed Tracing

- OpenTelemetry Collector
- Jaeger

---

# 🏗️ Architecture

## High-Level Architecture

![Architecture](docs/images/architecture-overview.png)

Additional Architecture Documentation:

- [Executive Architecture](docs/diagrams/01-executive-architecture.md)
- [AWS Infrastructure](docs/diagrams/02-aws-infrastructure.md)
- [CI/CD Pipeline](docs/diagrams/03-cicd-pipeline.md)
- [Kubernetes Architecture](docs/diagrams/04-kubernetes-architecture.md)

---

# 🛠️ Technology Stack

| Category | Technologies |
|-----------|--------------|
| Cloud Platform | AWS EC2, VPC, IAM |
| Infrastructure as Code | Terraform |
| Configuration Management | Ansible |
| Source Control | Git, GitHub |
| CI/CD | Jenkins, GitHub Webhooks |
| Programming Language | Go |
| Containerization | Docker |
| Container Registry | Docker Hub |
| Container Orchestration | Kubernetes (kubeadm) |
| Package Management | Helm |
| Ingress Controller | NGINX Ingress |
| Monitoring | Prometheus, Grafana |
| Logging | Fluent Bit, Elasticsearch, Kibana |
| Distributed Tracing | OpenTelemetry Collector, Jaeger |
| Security | Trivy |

---

# 📁 Repository Structure

```text
.
├── ansible/
├── docker/
├── docs/
│   ├── diagrams/
│   ├── images/
│   └── assets/
├── helm/
├── internal/
├── kubernetes/
├── observability/
├── scripts/
├── static/
├── templates/
├── terraform/
├── Jenkinsfile
├── go.mod
├── go.sum
└── README.md
```

---

# 🚀 Project Workflow

```text
Developer
     │
     ▼
GitHub Repository
     │
GitHub Webhook
     │
     ▼
Jenkins Pipeline
     │
 ├── Build Go Application
 ├── Build Docker Image
 ├── Trivy Scan
 ├── Push Image to Docker Hub
 └── Helm Deployment
     │
     ▼
Kubernetes Cluster
     │
     ▼
DevOps Dashboard
     │
     ├── Metrics → Prometheus → Grafana
     ├── Logs → Fluent Bit → Elasticsearch → Kibana
     └── Traces → OpenTelemetry → Jaeger
```

---

# 📸 Screenshots

> Screenshots will be added after deployment.

- AWS Infrastructure
- Jenkins Pipeline
- Kubernetes Cluster
- DevOps Dashboard
- Prometheus
- Grafana
- Kibana
- Jaeger

---

# 🚀 Getting Started

Detailed deployment documentation is available inside the project documentation.

- Terraform Infrastructure
- Ansible Configuration
- Jenkins CI/CD
- Kubernetes Deployment
- Monitoring
- Logging
- Distributed Tracing

---

# 🔮 Future Enhancements

- GitHub Actions Support
- ArgoCD GitOps Deployment
- Horizontal Pod Autoscaler
- Cluster Autoscaler
- Loki Integration
- Multi-Environment Deployments
- Multi-Region Disaster Recovery

---

# 👨‍💻 Author

## Navneet Singh Gour

**DevOps Engineer | Cloud | Kubernetes | AWS | Terraform | Ansible | Jenkins**

- GitHub: https://github.com/NavneetSinghGour
- LinkedIn: https://www.linkedin.com/in/navneetsinghgour/
- Email: navneetsinghgour480@gmail.com

If you found this project useful, consider giving it a ⭐ on GitHub.
