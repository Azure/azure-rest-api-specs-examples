package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/preview/2022-04-01-preview/examples/Profile-PUT-WithNestedEndpoints.json
func ExampleProfilesClient_CreateOrUpdate_profilePutWithNestedEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewProfilesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myresourcegroup", "parentprofile", armtrafficmanager.Profile{
		Location: to.Ptr("global"),
		Properties: &armtrafficmanager.ProfileProperties{
			DNSConfig: &armtrafficmanager.DNSConfig{
				RelativeName: to.Ptr("parentprofile"),
				TTL:          to.Ptr[int64](35),
			},
			Endpoints: []*armtrafficmanager.Endpoint{
				{
					Name: to.Ptr("MyFirstNestedEndpoint"),
					Type: to.Ptr("Microsoft.Network/trafficManagerProfiles/nestedEndpoints"),
					Properties: &armtrafficmanager.EndpointProperties{
						EndpointStatus:        to.Ptr(armtrafficmanager.EndpointStatusEnabled),
						MinChildEndpoints:     to.Ptr[int64](2),
						MinChildEndpointsIPv4: to.Ptr[int64](1),
						MinChildEndpointsIPv6: to.Ptr[int64](2),
						Priority:              to.Ptr[int64](1),
						Target:                to.Ptr("firstnestedprofile.tmpreview.watmtest.azure-test.net"),
						Weight:                to.Ptr[int64](1),
					},
				},
				{
					Name: to.Ptr("MySecondNestedEndpoint"),
					Type: to.Ptr("Microsoft.Network/trafficManagerProfiles/nestedEndpoints"),
					Properties: &armtrafficmanager.EndpointProperties{
						EndpointStatus:        to.Ptr(armtrafficmanager.EndpointStatusEnabled),
						MinChildEndpoints:     to.Ptr[int64](2),
						MinChildEndpointsIPv4: to.Ptr[int64](2),
						MinChildEndpointsIPv6: to.Ptr[int64](1),
						Priority:              to.Ptr[int64](2),
						Target:                to.Ptr("secondnestedprofile.tmpreview.watmtest.azure-test.net"),
						Weight:                to.Ptr[int64](1),
					},
				}},
			MonitorConfig: &armtrafficmanager.MonitorConfig{
				Path:                      to.Ptr("/testpath.aspx"),
				IntervalInSeconds:         to.Ptr[int64](10),
				Port:                      to.Ptr[int64](80),
				TimeoutInSeconds:          to.Ptr[int64](5),
				ToleratedNumberOfFailures: to.Ptr[int64](2),
				Protocol:                  to.Ptr(armtrafficmanager.MonitorProtocolHTTP),
			},
			ProfileStatus:        to.Ptr(armtrafficmanager.ProfileStatusEnabled),
			TrafficRoutingMethod: to.Ptr(armtrafficmanager.TrafficRoutingMethodPriority),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
