package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay/v2"
)

// Generated from example definition: 2026-07-01-preview/Clusters/ListAvailableClustersGet.json
func ExampleClustersClient_ListAvailableClusterRegion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().ListAvailableClusterRegion(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrelay.ClustersClientListAvailableClusterRegionResponse{
	// 	AvailableRelayClustersList: armrelay.AvailableRelayClustersList{
	// 		Value: []*armrelay.AvailableRelayClusterRegion{
	// 			{
	// 				Location: to.Ptr("southcentralus"),
	// 			},
	// 			{
	// 				Location: to.Ptr("westus2"),
	// 			},
	// 		},
	// 	},
	// }
}
