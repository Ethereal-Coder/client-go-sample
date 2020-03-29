package flexible_client

import (
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {
	clientset, err := getClient("/home/sunyh/.kube/config")
	if err != nil {
		panic("get clientset failed")
	}
	listNodes(clientset)
}

func getClient(pathToCfg string) (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	if pathToCfg == "" {
		fmt.Println("Using in cluster config")
		config, err = rest.InClusterConfig()
	} else {
		fmt.Println("Using out of cluster config")
		config, err = clientcmd.BuildConfigFromFlags("", pathToCfg)
	}
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func listNodes(clientset *kubernetes.Clientset) {
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Failed to poll the nodes: %v", err)
	}

	for i, node := range nodes.Items {
		fmt.Println(i)
		jsonBytes, err := json.Marshal(node)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(jsonBytes))

		var storage int64
		for _, image := range node.Status.Images {
			storage += storage + image.SizeBytes
		}
		fmt.Println(storage)
	}
}
