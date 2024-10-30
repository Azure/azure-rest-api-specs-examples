package armcontainerorchestratorruntime_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerorchestratorruntime/armcontainerorchestratorruntime"
)

// Generated from example definition: 2024-03-01/LoadBalancers_Get.json
func ExampleLoadBalancersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoadBalancersClient().Get(ctx, "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", "testlb", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerorchestratorruntime.LoadBalancersClientGetResponse{
	// 	LoadBalancer: &armcontainerorchestratorruntime.LoadBalancer{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesRuntime/loadBalancers/testlb"),
	// 		Name: to.Ptr("testlb"),
	// 		Type: to.Ptr("Microsoft.KubernetesRuntime/loadBalancers"),
	// 		Properties: &armcontainerorchestratorruntime.LoadBalancerProperties{
	// 			ProvisioningState: to.Ptr(armcontainerorchestratorruntime.ProvisioningStateSucceeded),
	// 			Addresses: []*string{
	// 				to.Ptr("192.168.50.1/24"),
	// 				to.Ptr("192.168.51.2-192.168.51.10"),
	// 			},
	// 			ServiceSelector: map[string]*string{
	// 				"app": to.Ptr("frontend"),
	// 			},
	// 			AdvertiseMode: to.Ptr(armcontainerorchestratorruntime.AdvertiseModeARP),
	// 			BgpPeers: []*string{
	// 				to.Ptr("bgpPeer1"),
	// 			},
	// 		},
	// 	},
	// }
}
