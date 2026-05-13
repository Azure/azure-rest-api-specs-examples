package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBGarnetClusterListByResourceGroup.json
func ExampleGarnetClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGarnetClustersClient().NewListByResourceGroupPager("garnet-prod-rg", nil)
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
		// page = armcosmos.GarnetClustersClientListByResourceGroupResponse{
		// 	ListGarnetClusters: armcosmos.ListGarnetClusters{
		// 		Value: []*armcosmos.GarnetClusterResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/garnet-prod-rg/providers/Microsoft.DocumentDB/garnetClusters/gc1"),
		// 				Name: to.Ptr("garnet-prod"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/garnetClusters"),
		// 				Location: to.Ptr("West US"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armcosmos.GarnetClusterResourceProperties{
		// 					ProvisioningState: to.Ptr(armcosmos.GarnetCacheProvisioningStateSucceeded),
		// 					SubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
