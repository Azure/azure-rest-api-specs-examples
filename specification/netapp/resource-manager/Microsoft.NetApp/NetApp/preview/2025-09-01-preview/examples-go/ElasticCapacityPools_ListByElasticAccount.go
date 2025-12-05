package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/ElasticCapacityPools_ListByElasticAccount.json
func ExampleElasticCapacityPoolsClient_NewListByElasticAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticCapacityPoolsClient().NewListByElasticAccountPager("myRG", "account1", nil)
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
		// page = armnetapp.ElasticCapacityPoolsClientListByElasticAccountResponse{
		// 	ElasticCapacityPoolListResult: armnetapp.ElasticCapacityPoolListResult{
		// 		Value: []*armnetapp.ElasticCapacityPool{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticCapacityPools/pool1"),
		// 				Name: to.Ptr("account1/pool1"),
		// 				Type: to.Ptr("Microsoft.NetApp/elasticAccounts/elasticCapacityPools"),
		// 				Location: to.Ptr("eastus"),
		// 				Zones: []*string{
		// 					to.Ptr("1"),
		// 					to.Ptr("2"),
		// 					to.Ptr("3"),
		// 				},
		// 				Properties: &armnetapp.ElasticCapacityPoolProperties{
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 					ServiceLevel: to.Ptr(armnetapp.ElasticServiceLevelZoneRedundant),
		// 					Size: to.Ptr[int64](4398046511104),
		// 					TotalThroughputMibps: to.Ptr[float64](281.474),
		// 					CurrentZone: to.Ptr("1"),
		// 					AvailabilityStatus: to.Ptr(armnetapp.ElasticResourceAvailabilityStatusOnline),
		// 					SubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
		// 					ActiveDirectoryConfigResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/activeDirectoryConfigs/activeDirectoryConfig1"),
		// 					Encryption: &armnetapp.ElasticEncryptionConfiguration{
		// 						ElasticPoolEncryptionKeySource: to.Ptr(armnetapp.ElasticPoolEncryptionKeySourceNetApp),
		// 						KeyVaultPrivateEndpointResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myKeyVault/privateEndpointConnections/myKeyVaultPec"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
