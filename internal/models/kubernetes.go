package models

type KubernetesInfo struct {
	PodName   string `json:"podName"`
	Namespace string `json:"namespace"`
	NodeName  string `json:"nodeName"`
	PodIP     string `json:"podIP"`
	HostIP    string `json:"hostIP"`
}
