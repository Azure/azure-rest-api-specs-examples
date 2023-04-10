package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Profile-GET-WithTrafficViewDisabled.json
func ExampleProfilesClient_Get_profileGetWithTrafficViewDisabled() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrafficmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().Get(ctx, "azuresdkfornetautoresttrafficmanager1323", "azuresdkfornetautoresttrafficmanager3880", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Profile = armtrafficmanager.Profile{
	// 	Name: to.Ptr("azuresdkfornetautoresttrafficmanager3880"),
	// 	Type: to.Ptr("Microsoft.Network/trafficManagerProfiles"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager1323/providers/Microsoft.Network/trafficManagerProfiles/azuresdkfornetautoresttrafficmanager3880"),
	// 	Location: to.Ptr("global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armtrafficmanager.ProfileProperties{
	// 		DNSConfig: &armtrafficmanager.DNSConfig{
	// 			Fqdn: to.Ptr("azuresdkfornetautoresttrafficmanager3880.tmpreview.watmtest.azure-test.net"),
	// 			RelativeName: to.Ptr("azuresdkfornetautoresttrafficmanager3880"),
	// 			TTL: to.Ptr[int64](35),
	// 		},
	// 		Endpoints: []*armtrafficmanager.Endpoint{
	// 			{
	// 				Name: to.Ptr("My external endpoint"),
	// 				Type: to.Ptr("Microsoft.Network/trafficManagerProfiles/externalEndpoints"),
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager1323/providers/Microsoft.Network/trafficManagerProfiles/azuresdkfornetautoresttrafficmanager3880/externalEndpoints/My external endpoint"),
	// 				Properties: &armtrafficmanager.EndpointProperties{
	// 					EndpointLocation: to.Ptr("North Europe"),
	// 					EndpointMonitorStatus: to.Ptr(armtrafficmanager.EndpointMonitorStatusCheckingEndpoint),
	// 					EndpointStatus: to.Ptr(armtrafficmanager.EndpointStatusEnabled),
	// 					Priority: to.Ptr[int64](1),
	// 					Target: to.Ptr("foobar.contoso.com"),
	// 					Weight: to.Ptr[int64](1),
	// 				},
	// 		}},
	// 		MonitorConfig: &armtrafficmanager.MonitorConfig{
	// 			Path: to.Ptr("/testpath.aspx"),
	// 			IntervalInSeconds: to.Ptr[int64](30),
	// 			Port: to.Ptr[int64](80),
	// 			ProfileMonitorStatus: to.Ptr(armtrafficmanager.ProfileMonitorStatusCheckingEndpoints),
	// 			TimeoutInSeconds: to.Ptr[int64](10),
	// 			ToleratedNumberOfFailures: to.Ptr[int64](3),
	// 			Protocol: to.Ptr(armtrafficmanager.MonitorProtocolHTTP),
	// 		},
	// 		ProfileStatus: to.Ptr(armtrafficmanager.ProfileStatusEnabled),
	// 		TrafficRoutingMethod: to.Ptr(armtrafficmanager.TrafficRoutingMethodPerformance),
	// 		TrafficViewEnrollmentStatus: to.Ptr(armtrafficmanager.TrafficViewEnrollmentStatusDisabled),
	// 	},
	// }
}
