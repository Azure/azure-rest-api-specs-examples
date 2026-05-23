package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager/v2"
)

// Generated from example definition: 2024-04-01-preview/Profile-PUT-WithEndpointsAndRecordType.json
func ExampleProfilesClient_CreateOrUpdate_profilePutWithEndpointsAndRecordType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrafficmanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().CreateOrUpdate(ctx, "azuresdkfornetautoresttrafficmanager2583", "azuresdkfornetautoresttrafficmanager6192", armtrafficmanager.Profile{
		Properties: &armtrafficmanager.ProfileProperties{
			ProfileStatus:        to.Ptr(armtrafficmanager.ProfileStatusEnabled),
			TrafficRoutingMethod: to.Ptr(armtrafficmanager.TrafficRoutingMethodPerformance),
			DNSConfig: &armtrafficmanager.DNSConfig{
				RelativeName: to.Ptr("azuresdkfornetautoresttrafficmanager6192"),
				TTL:          to.Ptr[int64](35),
			},
			MonitorConfig: &armtrafficmanager.MonitorConfig{
				Protocol:                  to.Ptr(armtrafficmanager.MonitorProtocolHTTP),
				Port:                      to.Ptr[int64](80),
				Path:                      to.Ptr("/testpath.aspx"),
				IntervalInSeconds:         to.Ptr[int64](10),
				TimeoutInSeconds:          to.Ptr[int64](5),
				ToleratedNumberOfFailures: to.Ptr[int64](2),
			},
			Endpoints: []*armtrafficmanager.Endpoint{
				{
					Name: to.Ptr("My external endpoint"),
					Type: to.Ptr("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
					Properties: &armtrafficmanager.EndpointProperties{
						Target:           to.Ptr("foobar.contoso.com"),
						EndpointStatus:   to.Ptr(armtrafficmanager.EndpointStatusEnabled),
						EndpointLocation: to.Ptr("North Europe"),
					},
				},
			},
			RecordType: to.Ptr(armtrafficmanager.RecordTypeCNAME),
		},
		Location: to.Ptr("global"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armtrafficmanager.ProfilesClientCreateOrUpdateResponse{
	// 	Profile: armtrafficmanager.Profile{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/azuresdkfornetautoresttrafficmanager2583/providers/Microsoft.Network/trafficManagerProfiles/azuresdkfornetautoresttrafficmanager6192"),
	// 		Name: to.Ptr("azuresdkfornetautoresttrafficmanager6192"),
	// 		Type: to.Ptr("Microsoft.Network/trafficManagerProfiles"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armtrafficmanager.ProfileProperties{
	// 			ProfileStatus: to.Ptr(armtrafficmanager.ProfileStatusEnabled),
	// 			TrafficRoutingMethod: to.Ptr(armtrafficmanager.TrafficRoutingMethodPerformance),
	// 			DNSConfig: &armtrafficmanager.DNSConfig{
	// 				RelativeName: to.Ptr("azuresdkfornetautoresttrafficmanager6192"),
	// 				Fqdn: to.Ptr("azuresdkfornetautoresttrafficmanager6192.tmpreview.watmtest.azure-test.net"),
	// 				TTL: to.Ptr[int64](35),
	// 			},
	// 			MonitorConfig: &armtrafficmanager.MonitorConfig{
	// 				ProfileMonitorStatus: to.Ptr(armtrafficmanager.ProfileMonitorStatusCheckingEndpoints),
	// 				Protocol: to.Ptr(armtrafficmanager.MonitorProtocolHTTP),
	// 				Port: to.Ptr[int64](80),
	// 				Path: to.Ptr("/testpath.aspx"),
	// 				IntervalInSeconds: to.Ptr[int64](10),
	// 				ToleratedNumberOfFailures: to.Ptr[int64](2),
	// 				TimeoutInSeconds: to.Ptr[int64](5),
	// 			},
	// 			Endpoints: []*armtrafficmanager.Endpoint{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/azuresdkfornetautoresttrafficmanager2583/providers/Microsoft.Network/trafficManagerProfiles/azuresdkfornetautoresttrafficmanager6192/externalEndpoints/My external endpoint"),
	// 					Name: to.Ptr("My external endpoint"),
	// 					Type: to.Ptr("Microsoft.Network/trafficManagerProfiles/externalEndpoints"),
	// 					Properties: &armtrafficmanager.EndpointProperties{
	// 						EndpointStatus: to.Ptr(armtrafficmanager.EndpointStatusEnabled),
	// 						EndpointMonitorStatus: to.Ptr(armtrafficmanager.EndpointMonitorStatusCheckingEndpoint),
	// 						Target: to.Ptr("foobar.contoso.com"),
	// 						Weight: to.Ptr[int64](1),
	// 						Priority: to.Ptr[int64](1),
	// 						EndpointLocation: to.Ptr("North Europe"),
	// 					},
	// 				},
	// 			},
	// 			RecordType: to.Ptr(armtrafficmanager.RecordTypeCNAME),
	// 		},
	// 	},
	// }
}
