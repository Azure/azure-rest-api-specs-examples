package armcustomproviders_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/customproviders/resource-manager/Microsoft.CustomProviders/preview/2018-09-01-preview/examples/updateCustomRP.json
func ExampleCustomResourceProviderClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomproviders.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomResourceProviderClient().Update(ctx, "testRG", "newrp", armcustomproviders.ResourceProvidersUpdate{
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomRPManifest = armcustomproviders.CustomRPManifest{
	// 	Name: to.Ptr("newrp"),
	// 	Type: to.Ptr("Microsoft.CustomProviders/resourceProviders"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.CustomProviders/resourceProviders/newrp"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armcustomproviders.CustomRPManifestProperties{
	// 		Actions: []*armcustomproviders.CustomRPActionRouteDefinition{
	// 			{
	// 				Name: to.Ptr("TestAction"),
	// 				Endpoint: to.Ptr("https://mytestendpoint/"),
	// 				RoutingType: to.Ptr(armcustomproviders.ActionRoutingProxy),
	// 		}},
	// 		ProvisioningState: to.Ptr(armcustomproviders.ProvisioningStateSucceeded),
	// 		ResourceTypes: []*armcustomproviders.CustomRPResourceTypeRouteDefinition{
	// 			{
	// 				Name: to.Ptr("TestResource"),
	// 				Endpoint: to.Ptr("https://mytestendpoint2/"),
	// 				RoutingType: to.Ptr(armcustomproviders.ResourceTypeRoutingProxyCache),
	// 		}},
	// 	},
	// }
}
