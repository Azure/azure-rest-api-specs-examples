package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-GET-ByResourceGroup.json
func ExampleProfilesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrafficmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProfilesClient().NewListByResourceGroupPager("azuresdkfornetautoresttrafficmanager3640", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ProfileListResult = armtrafficmanager.ProfileListResult{
		// 	Value: []*armtrafficmanager.Profile{
		// 		{
		// 			Name: to.Ptr("azuresdkfornetautoresttrafficmanager1005"),
		// 			Type: to.Ptr("Microsoft.Network/trafficManagerProfiles"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager3640/providers/Microsoft.Network/trafficManagerProfiles/azuresdkfornetautoresttrafficmanager1005"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armtrafficmanager.ProfileProperties{
		// 				DNSConfig: &armtrafficmanager.DNSConfig{
		// 					Fqdn: to.Ptr("azuresdkfornetautoresttrafficmanager1005.tmpreview.watmtest.azure-test.net"),
		// 					RelativeName: to.Ptr("azuresdkfornetautoresttrafficmanager1005"),
		// 					TTL: to.Ptr[int64](35),
		// 				},
		// 				Endpoints: []*armtrafficmanager.Endpoint{
		// 					{
		// 						Name: to.Ptr("My external endpoint"),
		// 						Type: to.Ptr("Microsoft.Network/trafficManagerProfiles/externalEndpoints"),
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager3640/providers/Microsoft.Network/trafficManagerProfiles/azuresdkfornetautoresttrafficmanager1005/externalEndpoints/My external endpoint"),
		// 						Properties: &armtrafficmanager.EndpointProperties{
		// 							EndpointLocation: to.Ptr("North Europe"),
		// 							EndpointMonitorStatus: to.Ptr(armtrafficmanager.EndpointMonitorStatusCheckingEndpoint),
		// 							EndpointStatus: to.Ptr(armtrafficmanager.EndpointStatusEnabled),
		// 							Priority: to.Ptr[int64](1),
		// 							Target: to.Ptr("foobar.contoso.com"),
		// 							Weight: to.Ptr[int64](1),
		// 						},
		// 				}},
		// 				MonitorConfig: &armtrafficmanager.MonitorConfig{
		// 					Path: to.Ptr("/testpath.aspx"),
		// 					IntervalInSeconds: to.Ptr[int64](30),
		// 					Port: to.Ptr[int64](80),
		// 					ProfileMonitorStatus: to.Ptr(armtrafficmanager.ProfileMonitorStatusOnline),
		// 					TimeoutInSeconds: to.Ptr[int64](10),
		// 					ToleratedNumberOfFailures: to.Ptr[int64](3),
		// 					Protocol: to.Ptr(armtrafficmanager.MonitorProtocolHTTP),
		// 				},
		// 				ProfileStatus: to.Ptr(armtrafficmanager.ProfileStatusEnabled),
		// 				TrafficRoutingMethod: to.Ptr(armtrafficmanager.TrafficRoutingMethodPerformance),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("azuresdkfornetautoresttrafficmanager959"),
		// 			Type: to.Ptr("Microsoft.Network/trafficManagerProfiles"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager3640/providers/Microsoft.Network/trafficManagerProfiles/azuresdkfornetautoresttrafficmanager959"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armtrafficmanager.ProfileProperties{
		// 				DNSConfig: &armtrafficmanager.DNSConfig{
		// 					Fqdn: to.Ptr("azuresdkfornetautoresttrafficmanager959.tmpreview.watmtest.azure-test.net"),
		// 					RelativeName: to.Ptr("azuresdkfornetautoresttrafficmanager959"),
		// 					TTL: to.Ptr[int64](35),
		// 				},
		// 				Endpoints: []*armtrafficmanager.Endpoint{
		// 					{
		// 						Name: to.Ptr("My external endpoint"),
		// 						Type: to.Ptr("Microsoft.Network/trafficManagerProfiles/externalEndpoints"),
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager3640/providers/Microsoft.Network/trafficManagerProfiles/azuresdkfornetautoresttrafficmanager959/externalEndpoints/My external endpoint"),
		// 						Properties: &armtrafficmanager.EndpointProperties{
		// 							EndpointLocation: to.Ptr("North Europe"),
		// 							EndpointMonitorStatus: to.Ptr(armtrafficmanager.EndpointMonitorStatusOnline),
		// 							EndpointStatus: to.Ptr(armtrafficmanager.EndpointStatusEnabled),
		// 							Priority: to.Ptr[int64](1),
		// 							Target: to.Ptr("foobar.contoso.com"),
		// 							Weight: to.Ptr[int64](1),
		// 						},
		// 				}},
		// 				MonitorConfig: &armtrafficmanager.MonitorConfig{
		// 					Path: to.Ptr("/testpath.aspx"),
		// 					IntervalInSeconds: to.Ptr[int64](30),
		// 					Port: to.Ptr[int64](80),
		// 					ProfileMonitorStatus: to.Ptr(armtrafficmanager.ProfileMonitorStatusOnline),
		// 					TimeoutInSeconds: to.Ptr[int64](10),
		// 					ToleratedNumberOfFailures: to.Ptr[int64](3),
		// 					Protocol: to.Ptr(armtrafficmanager.MonitorProtocolHTTP),
		// 				},
		// 				ProfileStatus: to.Ptr(armtrafficmanager.ProfileStatusEnabled),
		// 				TrafficRoutingMethod: to.Ptr(armtrafficmanager.TrafficRoutingMethodPerformance),
		// 			},
		// 	}},
		// }
	}
}
