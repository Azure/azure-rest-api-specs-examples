package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/preview/2022-04-01-preview/examples/Endpoint-PUT-External-WithSubnetMapping.json
func ExampleEndpointsClient_CreateOrUpdate_endpointPutExternalWithSubnetMapping() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewEndpointsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "azuresdkfornetautoresttrafficmanager2191", "azuresdkfornetautoresttrafficmanager8224", armtrafficmanager.EndpointTypeExternalEndpoints, "My%20external%20endpoint", armtrafficmanager.Endpoint{
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
	// TODO: use response item
	_ = res
}
