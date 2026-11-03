package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v11"
)

// Generated from example definition: 2026-08-15-preview/ClusterPeers_Get.json
func ExampleClusterPeersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClusterPeersClient().Get(ctx, "contosoNetAppRg", "contoso-external-cluster", "peerA", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.ClusterPeersClientGetResponse{
	// 	ClusterPeer: armnetapp.ClusterPeer{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contoso-external-cluster/clusterPeers/peerA"),
	// 		Name: to.Ptr("peerA"),
	// 		Type: to.Ptr("Microsoft.NetApp/externalClusters/clusterPeers"),
	// 		Properties: &armnetapp.ClusterPeerProperties{
	// 			AnfClusterName: to.Ptr("contoso-anf-cluster"),
	// 			OntapClusterName: to.Ptr("contoso-ontap-cluster"),
	// 			OntapVersion: to.Ptr("9.15.1"),
	// 			Status: to.Ptr(armnetapp.ClusterPeerStatusActive),
	// 			HybridReplicationResourceIDs: []*string{
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/hybridReplications/contosoHybridReplication"),
	// 			},
	// 			Command: to.Ptr("cluster peer create -peer-addrs 10.0.2.100,10.0.2.101"),
	// 			NetworkInterfaces: []*armnetapp.ExternalResourceNetworkInterfaces{
	// 				{
	// 					NetworkInterfaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.Network/networkInterfaces/contosoNic1"),
	// 					PeerIPAddress: to.Ptr("10.0.2.100"),
	// 				},
	// 				{
	// 					NetworkInterfaceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.Network/networkInterfaces/contosoNic2"),
	// 					PeerIPAddress: to.Ptr("10.0.2.101"),
	// 				},
	// 			},
	// 			ClusterPeerSubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.Network/virtualNetworks/contosoVnet/subnets/contosoSubnet"),
	// 		},
	// 	},
	// }
}
