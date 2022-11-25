package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/preview/2022-04-01-preview/examples/Profile-PUT-MultiValue.json
func ExampleProfilesClient_CreateOrUpdate_profilePutMultiValue() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewProfilesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "azuresdkfornetautoresttrafficmanager1421", "azsmnet6386", armtrafficmanager.Profile{
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
	// TODO: use response item
	_ = res
}
