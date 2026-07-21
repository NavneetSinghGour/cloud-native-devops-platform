# Kubernetes Cluster Architecture

## Overview

The application is deployed on a Kubernetes cluster created using **kubeadm**.

The cluster consists of one control plane node and two worker nodes.

Traffic enters through the NGINX Ingress Controller and is routed to the DevOps Dashboard. The application integrates with the complete observability stack for monitoring, logging, and distributed tracing.

---

# Kubernetes Architecture

```mermaid
flowchart TB

%% =====================================
%% USER
%% =====================================

USER["🌍 End User"]

USER --> INGRESS

%% =====================================
%% CLUSTER
%% =====================================

subgraph CLUSTER["☸️ Kubernetes Cluster"]

INGRESS["NGINX Ingress Controller"]

SERVICE["Dashboard Service"]

DEPLOY["Dashboard Deployment"]

subgraph WORKERS["Worker Nodes"]

APP1["Dashboard Pod"]

APP2["Dashboard Pod"]

end

SERVICE --> DEPLOY

DEPLOY --> APP1
DEPLOY --> APP2

%% =====================================
%% OBSERVABILITY
%% =====================================

subgraph OBS["📊 Observability"]

PROM["Prometheus"]

GRAF["Grafana"]

FB["Fluent Bit DaemonSet"]

ES["Elasticsearch StatefulSet"]

KIB["Kibana"]

OTEL["OpenTelemetry Collector"]

JAEGER["Jaeger"]

end

APP1 --> PROM
APP2 --> PROM

APP1 --> FB
APP2 --> FB

FB --> ES
ES --> KIB

APP1 --> OTEL
APP2 --> OTEL

OTEL --> JAEGER

PROM --> GRAF

end

INGRESS --> SERVICE
```

---

# Cluster Components

## Control Plane

Responsible for:

- Kubernetes API Server
- Scheduler
- Controller Manager
- etcd
- Cluster management

---

## Worker Nodes

Host all production workloads.

The worker nodes run:

- DevOps Dashboard
- NGINX Ingress Controller
- Prometheus
- Grafana
- Fluent Bit
- Elasticsearch
- Kibana
- OpenTelemetry Collector
- Jaeger

---

## Request Flow

1. User accesses the application.
2. Request reaches the NGINX Ingress Controller.
3. Ingress routes traffic to the Dashboard Service.
4. Service forwards traffic to the Dashboard Pods.
5. Dashboard responds to the client.
6. Metrics, logs, and traces are generated simultaneously.

---

# Platform Features

- High Availability using multiple replicas
- Service-based networking
- Ingress-based routing
- Rolling updates
- Cluster-wide log collection
- Centralized monitoring
- Distributed tracing
