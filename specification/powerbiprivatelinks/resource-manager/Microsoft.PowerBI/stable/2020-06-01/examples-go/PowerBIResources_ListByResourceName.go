package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PowerBIResources_ListByResourceName.json
func ExamplePowerBIResourcesClient_ListByResourceName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbiprivatelinks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPowerBIResourcesClient().ListByResourceName(ctx, "resourceGroup", "azureResourceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TenantResourceArray = []*armpowerbiprivatelinks.TenantResource{
	// 	{
	// 		Name: to.Ptr("conection 1"),
	// 		Location: to.Ptr("west us"),
	// 		Properties: &armpowerbiprivatelinks.TenantProperties{
	// 			PrivateEndpointConnections: []*armpowerbiprivatelinks.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("myPrivateEndpointConnection"),
	// 					Type: to.Ptr("string"),
	// 					ID: to.Ptr("/subscriptions/a0020869-4d28-422a-89f4-c2413130d73c/resourceGroups/resourceGroup/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateEndpointConnections/myPrivateEndpointConnection.58ffb8de-89ad-41eb-9f8f-de0a7db9d721"),
	// 					Properties: &armpowerbiprivatelinks.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armpowerbiprivatelinks.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/a0020869-4d28-422a-89f4-c2413130d73c/resourceGroups/resourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armpowerbiprivatelinks.ConnectionState{
	// 							Description: to.Ptr("My private endpoint connection"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armpowerbiprivatelinks.PersistedConnectionStatusPending),
	// 						},
	// 					},
	// 			}},
	// 			TenantID: to.Ptr("ac2bc297-8a3e-46f3-972d-87c2b4ae6e2f"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// }}
}
