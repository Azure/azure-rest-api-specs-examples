package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-GET-BySubscription.json
func ExampleProfilesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrafficmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProfilesClient().NewListBySubscriptionPager(nil)
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
		// 			Name: to.Ptr("azsmnet5183"),
		// 			Type: to.Ptr("Microsoft.Network/trafficManagerProfiles"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azsmnet1719/providers/Microsoft.Network/trafficManagerProfiles/azsmnet5183"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armtrafficmanager.ProfileProperties{
		// 				DNSConfig: &armtrafficmanager.DNSConfig{
		// 					Fqdn: to.Ptr("azsmnet4696.tmpreview.watmtest.azure-test.net"),
		// 					RelativeName: to.Ptr("azsmnet4696"),
		// 					TTL: to.Ptr[int64](35),
		// 				},
		// 				Endpoints: []*armtrafficmanager.Endpoint{
		// 				},
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
		// 			Name: to.Ptr("azuresdkfornetautoresttrafficmanager3440"),
		// 			Type: to.Ptr("Microsoft.Network/trafficManagerProfiles"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager5168/providers/Microsoft.Network/trafficManagerProfiles/azuresdkfornetautoresttrafficmanager3440"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armtrafficmanager.ProfileProperties{
		// 				DNSConfig: &armtrafficmanager.DNSConfig{
		// 					Fqdn: to.Ptr("azuresdkfornetautoresttrafficmanager3440.tmpreview.watmtest.azure-test.net"),
		// 					RelativeName: to.Ptr("azuresdkfornetautoresttrafficmanager3440"),
		// 					TTL: to.Ptr[int64](35),
		// 				},
		// 				Endpoints: []*armtrafficmanager.Endpoint{
		// 					{
		// 						Name: to.Ptr("My external endpoint"),
		// 						Type: to.Ptr("Microsoft.Network/trafficManagerProfiles/externalEndpoints"),
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager5168/providers/Microsoft.Network/trafficManagerProfiles/azuresdkfornetautoresttrafficmanager3440/externalEndpoints/My external endpoint"),
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
		// 		},
		// 		{
		// 			Name: to.Ptr("onesdk8819"),
		// 			Type: to.Ptr("Microsoft.Network/trafficManagerProfiles"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/onesdk9785/providers/Microsoft.Network/trafficManagerProfiles/onesdk8819"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armtrafficmanager.ProfileProperties{
		// 				DNSConfig: &armtrafficmanager.DNSConfig{
		// 					Fqdn: to.Ptr("onesdk7242.tmpreview.watmtest.azure-test.net"),
		// 					RelativeName: to.Ptr("onesdk7242"),
		// 					TTL: to.Ptr[int64](51),
		// 				},
		// 				Endpoints: []*armtrafficmanager.Endpoint{
		// 					{
		// 						Name: to.Ptr("MyNestedEndpoint"),
		// 						Type: to.Ptr("Microsoft.Network/trafficManagerProfiles/nestedEndpoints"),
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/onesdk9785/providers/Microsoft.Network/trafficManagerProfiles/onesdk8819/nestedEndpoints/MyNestedEndpoint"),
		// 						Properties: &armtrafficmanager.EndpointProperties{
		// 							EndpointLocation: to.Ptr("West Europe"),
		// 							EndpointMonitorStatus: to.Ptr(armtrafficmanager.EndpointMonitorStatusStopped),
		// 							EndpointStatus: to.Ptr(armtrafficmanager.EndpointStatusEnabled),
		// 							MinChildEndpoints: to.Ptr[int64](1),
		// 							Priority: to.Ptr[int64](1),
		// 							Target: to.Ptr("onesdk4285.tmpreview.watmtest.azure-test.net"),
		// 							TargetResourceID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/onesdk9785/providers/Microsoft.Network/trafficManagerProfiles/onesdk1792"),
		// 							Weight: to.Ptr[int64](1),
		// 						},
		// 				}},
		// 				MonitorConfig: &armtrafficmanager.MonitorConfig{
		// 					Path: to.Ptr("/testparent.asp"),
		// 					IntervalInSeconds: to.Ptr[int64](30),
		// 					Port: to.Ptr[int64](111),
		// 					ProfileMonitorStatus: to.Ptr(armtrafficmanager.ProfileMonitorStatusOnline),
		// 					TimeoutInSeconds: to.Ptr[int64](10),
		// 					ToleratedNumberOfFailures: to.Ptr[int64](3),
		// 					Protocol: to.Ptr(armtrafficmanager.MonitorProtocolHTTPS),
		// 				},
		// 				ProfileStatus: to.Ptr(armtrafficmanager.ProfileStatusEnabled),
		// 				TrafficRoutingMethod: to.Ptr(armtrafficmanager.TrafficRoutingMethodPerformance),
		// 			},
		// 	}},
		// }
	}
}
