package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/abd5d0016f12f6862cae88ef70f1333e84e20c07/specification/trafficmanager/resource-manager/Microsoft.Network/preview/2022-04-01-preview/examples/Endpoint-PUT-External-WithGeoMapping.json
func ExampleEndpointsClient_CreateOrUpdate_endpointPutExternalWithGeoMapping() {
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
			GeoMapping: []*string{
				to.Ptr("GEO-AS"),
				to.Ptr("GEO-AF")},
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
	// 		GeoMapping: []*string{
	// 			to.Ptr("GEO-AS"),
	// 			to.Ptr("GEO-AF")},
	// 			Priority: to.Ptr[int64](1),
	// 			Target: to.Ptr("foobar.contoso.com"),
	// 			Weight: to.Ptr[int64](1),
	// 		},
	// 	}
}
