package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBGarnetClusterCreate.json
func ExampleGarnetClustersClient_BeginCreateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGarnetClustersClient().BeginCreateUpdate(ctx, "garnet-prod-rg", "garnet-prod", armcosmos.GarnetClusterResource{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.GarnetClusterResourceProperties{
			SubnetID:          to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
			NodeCount:         to.Ptr[int32](4),
			NodeSKU:           to.Ptr("Standard_DS13_v2"),
			ReplicationFactor: to.Ptr[int32](2),
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
	// res = armcosmos.GarnetClustersClientCreateUpdateResponse{
	// 	GarnetClusterResource: &armcosmos.GarnetClusterResource{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/garnet-prod-rg/providers/Microsoft.DocumentDB/garnetClusters/garnet-prod"),
	// 		Name: to.Ptr("garnet-prod"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/garnetClusters"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armcosmos.GarnetClusterResourceProperties{
	// 			ProvisioningState: to.Ptr(armcosmos.GarnetCacheProvisioningStateSucceeded),
	// 			SubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
	// 		},
	// 	},
	// }
}
