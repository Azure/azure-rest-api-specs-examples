package armcontainerorchestratorruntime_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerorchestratorruntime/armcontainerorchestratorruntime"
)

// Generated from example definition: 2024-03-01/BgpPeers_CreateOrUpdate.json
func ExampleBgpPeersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBgpPeersClient().BeginCreateOrUpdate(ctx, "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", "testpeer", armcontainerorchestratorruntime.BgpPeer{
		Properties: &armcontainerorchestratorruntime.BgpPeerProperties{
			MyAsn:       to.Ptr[int32](64500),
			PeerAsn:     to.Ptr[int32](64501),
			PeerAddress: to.Ptr("10.0.0.1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerorchestratorruntime.BgpPeersClientCreateOrUpdateResponse{
	// 	BgpPeer: &armcontainerorchestratorruntime.BgpPeer{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesRuntime/bgpPeers/testpeer"),
	// 		Name: to.Ptr("testpeer"),
	// 		Type: to.Ptr("Microsoft.KubernetesRuntime/BgpPeers"),
	// 		Properties: &armcontainerorchestratorruntime.BgpPeerProperties{
	// 			ProvisioningState: to.Ptr(armcontainerorchestratorruntime.ProvisioningStateSucceeded),
	// 			MyAsn: to.Ptr[int32](64500),
	// 			PeerAsn: to.Ptr[int32](64501),
	// 			PeerAddress: to.Ptr("10.0.0.1"),
	// 		},
	// 	},
	// }
}
