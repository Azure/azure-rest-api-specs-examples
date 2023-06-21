package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Endpoint-PUT-External-WithCustomHeaders.json
func ExampleEndpointsClient_CreateOrUpdate_endpointPutExternalWithCustomHeaders() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrafficmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().CreateOrUpdate(ctx, "azuresdkfornetautoresttrafficmanager1421", "azsmnet6386", armtrafficmanager.EndpointTypeExternalEndpoints, "azsmnet7187", armtrafficmanager.Endpoint{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Endpoint = armtrafficmanager.Endpoint{
	// 	Name: to.Ptr("azsmnet7187"),
	// 	Type: to.Ptr("Microsoft.Network/trafficManagerProfiles/externalEndpoints"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager1421/providers/Microsoft.Network/trafficManagerProfiles/azsmnet6386/externalEndpoints/azsmnet7187"),
	// 	Properties: &armtrafficmanager.EndpointProperties{
	// 		CustomHeaders: []*armtrafficmanager.EndpointPropertiesCustomHeadersItem{
	// 			{
	// 				Name: to.Ptr("header-1"),
	// 				Value: to.Ptr("value-1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("header-2"),
	// 				Value: to.Ptr("value-2"),
	// 		}},
	// 		EndpointLocation: to.Ptr("North Europe"),
	// 		EndpointMonitorStatus: to.Ptr(armtrafficmanager.EndpointMonitorStatusCheckingEndpoint),
	// 		EndpointStatus: to.Ptr(armtrafficmanager.EndpointStatusEnabled),
	// 		Priority: to.Ptr[int64](1),
	// 		Target: to.Ptr("foobar.contoso.com"),
	// 		Weight: to.Ptr[int64](1),
	// 	},
	// }
}
