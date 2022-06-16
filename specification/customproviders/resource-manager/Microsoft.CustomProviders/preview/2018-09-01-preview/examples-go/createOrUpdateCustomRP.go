package armcustomproviders_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customproviders/resource-manager/Microsoft.CustomProviders/preview/2018-09-01-preview/examples/createOrUpdateCustomRP.json
func ExampleCustomResourceProviderClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcustomproviders.NewCustomResourceProviderClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testRG",
		"newrp",
		armcustomproviders.CustomRPManifest{
			Location: to.Ptr("eastus"),
			Properties: &armcustomproviders.CustomRPManifestProperties{
				Actions: []*armcustomproviders.CustomRPActionRouteDefinition{
					{
						Name:        to.Ptr("TestAction"),
						Endpoint:    to.Ptr("https://mytestendpoint/"),
						RoutingType: to.Ptr(armcustomproviders.ActionRoutingProxy),
					}},
				ResourceTypes: []*armcustomproviders.CustomRPResourceTypeRouteDefinition{
					{
						Name:        to.Ptr("TestResource"),
						Endpoint:    to.Ptr("https://mytestendpoint2/"),
						RoutingType: to.Ptr(armcustomproviders.ResourceTypeRoutingProxyCache),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
