package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBGarnetClusterGet.json
func ExampleGarnetClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGarnetClustersClient().Get(ctx, "garnet-prod-rg", "garnet-prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.GarnetClustersClientGetResponse{
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
	// 			EndPoints: []*armcosmos.GarnetClusterResourcePropertiesEndPointsItem{
	// 				{
	// 					IPAddress: to.Ptr("10.0.0.13"),
	// 					Port: to.Ptr[int32](9999),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("10.0.0.14"),
	// 					Port: to.Ptr[int32](9999),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("10.0.0.15"),
	// 					Port: to.Ptr[int32](9999),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("10.0.0.19"),
	// 					Port: to.Ptr[int32](9999),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("10.0.0.20"),
	// 					Port: to.Ptr[int32](9999),
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("10.0.0.21"),
	// 					Port: to.Ptr[int32](9999),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
