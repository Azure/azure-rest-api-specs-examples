package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PrivateLinkResources_Get.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbiprivatelinks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "resourceGroup", "azureResourceName", "tenant", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armpowerbiprivatelinks.PrivateLinkResource{
	// 	Name: to.Ptr("tenant"),
	// 	Type: to.Ptr("Microsoft.PowerBI/{resourceType}/privateLinkResources"),
	// 	ID: to.Ptr("subscriptions/a0020869-4d28-422a-89f4-c2413130d73c/resourceGroups/resourceGroup/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateLinkResources/tenant"),
	// 	Properties: &armpowerbiprivatelinks.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("tenant"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("tenant"),
	// 			to.Ptr("capacity:3df897a4f10b49e9bddb0e9cf062adba")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.powerbi.com")},
	// 			},
	// 		}
}
