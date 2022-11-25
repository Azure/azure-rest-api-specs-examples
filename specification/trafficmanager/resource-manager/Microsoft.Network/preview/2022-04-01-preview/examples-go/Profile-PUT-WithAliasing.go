package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/preview/2022-04-01-preview/examples/Profile-PUT-WithAliasing.json
func ExampleProfilesClient_CreateOrUpdate_profilePutWithAliasing() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewProfilesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "azuresdkfornetautoresttrafficmanager2583", "azuresdkfornetautoresttrafficmanager6192", armtrafficmanager.Profile{
		Location: to.Ptr("global"),
		Properties: &armtrafficmanager.ProfileProperties{
			AllowedEndpointRecordTypes: []*armtrafficmanager.AllowedEndpointRecordType{
				to.Ptr(armtrafficmanager.AllowedEndpointRecordTypeDomainName)},
			DNSConfig: &armtrafficmanager.DNSConfig{
				RelativeName: to.Ptr("azuresdkfornetautoresttrafficmanager6192"),
				TTL:          to.Ptr[int64](35),
			},
			Endpoints: []*armtrafficmanager.Endpoint{
				{
					Name: to.Ptr("My external endpoint"),
					Type: to.Ptr("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
					Properties: &armtrafficmanager.EndpointProperties{
						EndpointLocation: to.Ptr("North Europe"),
						EndpointStatus:   to.Ptr(armtrafficmanager.EndpointStatusEnabled),
						Target:           to.Ptr("foobar.contoso.com"),
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
			TrafficRoutingMethod: to.Ptr(armtrafficmanager.TrafficRoutingMethodPerformance),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
