package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-06-02-preview/AgentPoolsListBootstrapData.json
func ExampleAgentPoolsClient_ListBootstrapData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgentPoolsClient().ListBootstrapData(ctx, "rg1", "clustername1", "flexnode1", armcontainerservice.ListBootstrapDataRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservice.AgentPoolsClientListBootstrapDataResponse{
	// 	PoolBootstrapData: armcontainerservice.PoolBootstrapData{
	// 		Azure: &armcontainerservice.BootstrapAzureConfig{
	// 			ResourceManagerEndpoint: to.Ptr("https://management.azure.com/"),
	// 			TargetCluster: &armcontainerservice.BootstrapTargetCluster{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	// 			},
	// 			TargetAgentPoolName: to.Ptr("flexnode1"),
	// 			BootstrapToken: &armcontainerservice.BootstrapTokenInfo{
	// 				Token: to.Ptr("<redacted-bootstrap-token>"),
	// 			},
	// 		},
	// 		Components: &armcontainerservice.BootstrapComponentVersions{
	// 			Kubernetes: to.Ptr("1.32.7"),
	// 			Containerd: to.Ptr("2.0.4"),
	// 			Runc: to.Ptr("1.2.1"),
	// 		},
	// 		Networking: &armcontainerservice.BootstrapNetworkingConfig{
	// 			DNSServiceIP: to.Ptr("10.0.0.10"),
	// 			CniVersion: to.Ptr("1.4.1"),
	// 		},
	// 		Node: &armcontainerservice.BootstrapNodeConfig{
	// 			MaxPods: to.Ptr[int32](110),
	// 			Labels: map[string]*string{
	// 				"node-role": to.Ptr("edge"),
	// 			},
	// 			Taints: []*string{
	// 				to.Ptr("aks.azure.com/flex-node=true:NoSchedule"),
	// 			},
	// 			Kubelet: &armcontainerservice.BootstrapKubeletConfig{
	// 				ClusterFQDN: to.Ptr("cluster-xxxx.hcp.eastus2.azmk8s.io"),
	// 				CaCertData: to.Ptr("<redacted-base64-pem-ca-cert>"),
	// 			},
	// 		},
	// 	},
	// }
}
