package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-PUT-MultiValue.json
func ExampleProfilesClient_CreateOrUpdate_profilePutMultiValue() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrafficmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().CreateOrUpdate(ctx, "azuresdkfornetautoresttrafficmanager1421", "azsmnet6386", armtrafficmanager.Profile{
		Location: to.Ptr("global"),
		Properties: &armtrafficmanager.ProfileProperties{
			DNSConfig: &armtrafficmanager.DNSConfig{
				RelativeName: to.Ptr("azsmnet6386"),
				TTL:          to.Ptr[int64](35),
			},
			MaxReturn: to.Ptr[int64](2),
			MonitorConfig: &armtrafficmanager.MonitorConfig{
				Path:     to.Ptr("/testpath.aspx"),
				Port:     to.Ptr[int64](80),
				Protocol: to.Ptr(armtrafficmanager.MonitorProtocolHTTP),
			},
			ProfileStatus:               to.Ptr(armtrafficmanager.ProfileStatusEnabled),
			TrafficRoutingMethod:        to.Ptr(armtrafficmanager.TrafficRoutingMethodMultiValue),
			TrafficViewEnrollmentStatus: to.Ptr(armtrafficmanager.TrafficViewEnrollmentStatusDisabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Profile = armtrafficmanager.Profile{
	// 	Name: to.Ptr("azsmnet6386"),
	// 	Type: to.Ptr("Microsoft.Network/trafficManagerProfiles"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager1421/providers/Microsoft.Network/trafficManagerProfiles/azsmnet6386"),
	// 	Location: to.Ptr("global"),
	// 	Properties: &armtrafficmanager.ProfileProperties{
	// 		DNSConfig: &armtrafficmanager.DNSConfig{
	// 			Fqdn: to.Ptr("azsmnet6386.tmpreview.watmtest.azure-test.net"),
	// 			RelativeName: to.Ptr("azsmnet6386"),
	// 			TTL: to.Ptr[int64](35),
	// 		},
	// 		Endpoints: []*armtrafficmanager.Endpoint{
	// 		},
	// 		MaxReturn: to.Ptr[int64](2),
	// 		MonitorConfig: &armtrafficmanager.MonitorConfig{
	// 			Path: to.Ptr("/testpath.aspx"),
	// 			IntervalInSeconds: to.Ptr[int64](30),
	// 			Port: to.Ptr[int64](80),
	// 			ProfileMonitorStatus: to.Ptr(armtrafficmanager.ProfileMonitorStatusInactive),
	// 			TimeoutInSeconds: to.Ptr[int64](10),
	// 			ToleratedNumberOfFailures: to.Ptr[int64](3),
	// 			Protocol: to.Ptr(armtrafficmanager.MonitorProtocolHTTP),
	// 		},
	// 		ProfileStatus: to.Ptr(armtrafficmanager.ProfileStatusEnabled),
	// 		TrafficRoutingMethod: to.Ptr(armtrafficmanager.TrafficRoutingMethodMultiValue),
	// 		TrafficViewEnrollmentStatus: to.Ptr(armtrafficmanager.TrafficViewEnrollmentStatusDisabled),
	// 	},
	// }
}
