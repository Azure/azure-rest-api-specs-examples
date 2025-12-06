package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/Caches_ListByCapacityPools.json
func ExampleCachesClient_NewListByCapacityPoolsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCachesClient().NewListByCapacityPoolsPager("myRG", "account1", "pool1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armnetapp.CachesClientListByCapacityPoolsResponse{
		// 	CacheList: armnetapp.CacheList{
		// 		Value: []*armnetapp.Cache{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/caches/cache1"),
		// 				Name: to.Ptr("account1/pool1/cache1"),
		// 				Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/caches"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetapp.CacheProperties{
		// 					Filepath: to.Ptr("cache-west-us2-01"),
		// 					Size: to.Ptr[int64](107374182400),
		// 					ProvisioningState: to.Ptr(armnetapp.CacheProvisioningStateSucceeded),
		// 					CacheState: to.Ptr(armnetapp.CacheLifeCycleStateSucceeded),
		// 					CacheSubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/cacheVnet/subnets/cacheSubnet1"),
		// 					PeeringSubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/icLifVnet/subnets/peeringSubnet1"),
		// 					EncryptionKeySource: to.Ptr(armnetapp.EncryptionKeySourceMicrosoftNetApp),
		// 					OriginClusterInformation: &armnetapp.OriginClusterInformation{
		// 						PeerClusterName: to.Ptr("cluster1"),
		// 						PeerAddresses: []*string{
		// 							to.Ptr("192.0.2.10"),
		// 							to.Ptr("192.0.2.11"),
		// 						},
		// 						PeerVserverName: to.Ptr("vserver1"),
		// 						PeerVolumeName: to.Ptr("originvol1"),
		// 					},
		// 					CifsChangeNotifications: to.Ptr(armnetapp.CifsChangeNotifyStateDisabled),
		// 					GlobalFileLocking: to.Ptr(armnetapp.GlobalFileLockingStateDisabled),
		// 					WriteBack: to.Ptr(armnetapp.EnableWriteBackStateDisabled),
		// 					Ldap: to.Ptr(armnetapp.LdapStateEnabled),
		// 					LdapServerType: to.Ptr(armnetapp.LdapServerTypeOpenLDAP),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
