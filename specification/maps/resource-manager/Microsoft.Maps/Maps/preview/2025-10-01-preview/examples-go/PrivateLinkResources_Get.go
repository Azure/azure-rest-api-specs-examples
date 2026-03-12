package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps/v2"
)

// Generated from example definition: 2025-10-01-preview/PrivateLinkResources_Get.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("21a9967a-e8a9-4656-a70b-96ff1c4d05a0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "myResourceGroup", "myMapsAccount", "mapsAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmaps.PrivateLinkResourcesClientGetResponse{
	// 	PrivateLinkResource: &armmaps.PrivateLinkResource{
	// 		Name: to.Ptr("mapsAccount"),
	// 		Type: to.Ptr("Microsoft.Maps/accounts/privateLinkResources"),
	// 		ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount/privateLinkResources/mapsAccount"),
	// 		Properties: &armmaps.PrivateLinkResourceProperties{
	// 			GroupID: to.Ptr("mapsAccount"),
	// 			RequiredMembers: []*string{
	// 				to.Ptr("myMapsAccount"),
	// 			},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.eastus.maps.azure.com"),
	// 			},
	// 		},
	// 	},
	// }
}
