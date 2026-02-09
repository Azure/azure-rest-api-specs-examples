package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ApplicationGatewayBackendHealthTest.json
func ExampleApplicationGatewaysClient_BeginBackendHealthOnDemand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationGatewaysClient().BeginBackendHealthOnDemand(ctx, "rg1", "appgw", armnetwork.ApplicationGatewayOnDemandProbe{
		Path: to.Ptr("/"),
		BackendAddressPool: &armnetwork.SubResource{
			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendaddressPools/MFAnalyticsPool"),
		},
		BackendHTTPSettings: &armnetwork.SubResource{
			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/MFPoolSettings"),
		},
		PickHostNameFromBackendHTTPSettings: to.Ptr(true),
		Timeout:                             to.Ptr[int32](30),
		Protocol:                            to.Ptr(armnetwork.ApplicationGatewayProtocolHTTP),
	}, &armnetwork.ApplicationGatewaysClientBeginBackendHealthOnDemandOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationGatewayBackendHealthOnDemand = armnetwork.ApplicationGatewayBackendHealthOnDemand{
	// 	BackendAddressPool: &armnetwork.ApplicationGatewayBackendAddressPool{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendaddressPools/MFAnalyticsPool"),
	// 	},
	// 	BackendHealthHTTPSettings: &armnetwork.ApplicationGatewayBackendHealthHTTPSettings{
	// 		BackendHTTPSettings: &armnetwork.ApplicationGatewayBackendHTTPSettings{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/MFPoolSettings"),
	// 		},
	// 		Servers: []*armnetwork.ApplicationGatewayBackendHealthServer{
	// 			{
	// 				Address: to.Ptr("10.220.1.4"),
	// 				Health: to.Ptr(armnetwork.ApplicationGatewayBackendHealthServerHealthUp),
	// 			},
	// 			{
	// 				Address: to.Ptr("10.220.1.5"),
	// 				Health: to.Ptr(armnetwork.ApplicationGatewayBackendHealthServerHealthUp),
	// 		}},
	// 	},
	// }
}
