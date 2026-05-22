package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/PrivateLinkResources_Get.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "exampleresourcegroup", "examplecluster", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.PrivateLinkResourcesClientGetResponse{
	// 	PrivateLinkResource: &armhorizondb.PrivateLinkResource{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/privateLinkResources"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/privateLinkResources/default"),
	// 		Properties: &armhorizondb.PrivateLinkResourceProperties{
	// 			GroupID: to.Ptr("default"),
	// 			RequiredMembers: []*string{
	// 				to.Ptr("examplecluster.clusterid"),
	// 				to.Ptr("examplecluster.clusterid.ro"),
	// 			},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.horizondb.azure.com"),
	// 			},
	// 		},
	// 	},
	// }
}
