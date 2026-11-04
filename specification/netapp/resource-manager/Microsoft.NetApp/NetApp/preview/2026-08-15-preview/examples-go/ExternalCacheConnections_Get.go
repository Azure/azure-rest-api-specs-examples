package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v11"
)

// Generated from example definition: 2026-08-15-preview/ExternalCacheConnections_Get.json
func ExampleExternalCacheConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExternalCacheConnectionsClient().Get(ctx, "contosoNetAppRg", "contosoCacheConnection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.ExternalCacheConnectionsClientGetResponse{
	// 	ExternalCacheConnection: armnetapp.ExternalCacheConnection{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalCacheConnections/contosoCacheConnection"),
	// 		Name: to.Ptr("contosoCacheConnection"),
	// 		Type: to.Ptr("Microsoft.NetApp/externalCacheConnections"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.ExternalCacheConnectionProperties{
	// 			ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
	// 			OriginVolumeResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/netAppAccounts/contosoAccount/capacityPools/contosoPool/volumes/contosoVolume"),
	// 			ExternalCacheName: to.Ptr("contosoCache"),
	// 			ExternalSVMResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contosoExternalCluster/externalSvms/contosoExternalSvm"),
	// 			ClusterPeerSubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.Network/virtualNetworks/contosoVnet/subnets/contosoSubnet"),
	// 			ClusterPeerResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contosoExternalCluster/clusterPeers/peerA"),
	// 			SVMPeerResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contosoExternalCluster/clusterPeers/peerA/svmPeers/contosoSvmPeerA"),
	// 			SetupInformation: &armnetapp.SetupInformation{
	// 				SetupStatus: to.Ptr(armnetapp.ExternalConnectionSetupStateNoActionRequired),
	// 				StatusInformation: to.Ptr("Setup is complete."),
	// 				Command: to.Ptr(""),
	// 			},
	// 		},
	// 	},
	// }
}
