package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v11"
)

// Generated from example definition: 2026-04-15-preview/Caches_ResetSmbPassword.json
func ExampleCachesClient_BeginResetSmbPassword() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCachesClient().BeginResetSmbPassword(ctx, "myResourceGroup", "account1", "pool1", "cache1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.CachesClientResetSmbPasswordResponse{
	// 	Cache: armnetapp.Cache{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/caches/cache1"),
	// 		Name: to.Ptr("account1/pool1/cache1"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/caches"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.CacheProperties{
	// 			FilePath: to.Ptr("cache1"),
	// 			Size: to.Ptr[int64](107374182400),
	// 			ProvisioningState: to.Ptr(armnetapp.CacheProvisioningStateSucceeded),
	// 			CacheState: to.Ptr(armnetapp.CacheLifeCycleStateSucceeded),
	// 			CacheSubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/cacheSubnet"),
	// 			PeeringSubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/peeringSubnet"),
	// 			EncryptionKeySource: to.Ptr(armnetapp.EncryptionKeySourceMicrosoftNetApp),
	// 			OriginClusterInformation: &armnetapp.OriginClusterInformation{
	// 				PeerClusterName: to.Ptr("cluster1"),
	// 				PeerAddresses: []*string{
	// 					to.Ptr("10.0.0.5"),
	// 				},
	// 				PeerVserverName: to.Ptr("vserver1"),
	// 				PeerVolumeName: to.Ptr("volume1"),
	// 			},
	// 		},
	// 	},
	// }
}
