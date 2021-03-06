package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Endpoint-PUT-External-WithCustomHeaders.json
func ExampleEndpointsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewEndpointsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"azuresdkfornetautoresttrafficmanager1421",
		"azsmnet6386",
		armtrafficmanager.EndpointTypeExternalEndpoints,
		"azsmnet7187",
		armtrafficmanager.Endpoint{
			Name: to.Ptr("azsmnet7187"),
			Type: to.Ptr("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
			Properties: &armtrafficmanager.EndpointProperties{
				CustomHeaders: []*armtrafficmanager.EndpointPropertiesCustomHeadersItem{
					{
						Name:  to.Ptr("header-1"),
						Value: to.Ptr("value-1"),
					},
					{
						Name:  to.Ptr("header-2"),
						Value: to.Ptr("value-2"),
					}},
				EndpointLocation: to.Ptr("North Europe"),
				EndpointStatus:   to.Ptr(armtrafficmanager.EndpointStatusEnabled),
				Target:           to.Ptr("foobar.contoso.com"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
