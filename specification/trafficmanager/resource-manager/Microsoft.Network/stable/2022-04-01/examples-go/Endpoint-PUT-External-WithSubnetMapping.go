package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Endpoint-PUT-External-WithSubnetMapping.json
func ExampleEndpointsClient_CreateOrUpdate_endpointPutExternalWithSubnetMapping() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrafficmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().CreateOrUpdate(ctx, "azuresdkfornetautoresttrafficmanager2191", "azuresdkfornetautoresttrafficmanager8224", armtrafficmanager.EndpointTypeExternalEndpoints, "My%20external%20endpoint", armtrafficmanager.Endpoint{
		Name: to.Ptr("My external endpoint"),
		Type: to.Ptr("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
		Properties: &armtrafficmanager.EndpointProperties{
			EndpointStatus: to.Ptr(armtrafficmanager.EndpointStatusEnabled),
			Subnets: []*armtrafficmanager.EndpointPropertiesSubnetsItem{
				{
					First: to.Ptr("1.2.3.0"),
					Scope: to.Ptr[int32](24),
				},
				{
					First: to.Ptr("25.26.27.28"),
					Last:  to.Ptr("29.30.31.32"),
				}},
			Target: to.Ptr("foobar.contoso.com"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Endpoint = armtrafficmanager.Endpoint{
	// 	Name: to.Ptr("My external endpoint"),
	// 	Type: to.Ptr("Microsoft.Network/trafficManagerProfiles/externalEndpoints"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager2191/providers/Microsoft.Network/trafficManagerProfiles/azuresdkfornetautoresttrafficmanager8224/externalEndpoints/My external endpoint"),
	// 	Properties: &armtrafficmanager.EndpointProperties{
	// 		EndpointMonitorStatus: to.Ptr(armtrafficmanager.EndpointMonitorStatusCheckingEndpoint),
	// 		EndpointStatus: to.Ptr(armtrafficmanager.EndpointStatusEnabled),
	// 		Priority: to.Ptr[int64](1),
	// 		Subnets: []*armtrafficmanager.EndpointPropertiesSubnetsItem{
	// 			{
	// 				First: to.Ptr("1.2.3.0"),
	// 				Scope: to.Ptr[int32](24),
	// 			},
	// 			{
	// 				First: to.Ptr("25.26.27.28"),
	// 				Last: to.Ptr("29.30.31.32"),
	// 		}},
	// 		Target: to.Ptr("foobar.contoso.com"),
	// 		Weight: to.Ptr[int64](1),
	// 	},
	// }
}
