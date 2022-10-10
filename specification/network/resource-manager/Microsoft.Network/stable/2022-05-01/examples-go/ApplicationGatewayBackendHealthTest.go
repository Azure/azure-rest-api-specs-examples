package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ApplicationGatewayBackendHealthTest.json
func ExampleApplicationGatewaysClient_BeginBackendHealthOnDemand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewApplicationGatewaysClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginBackendHealthOnDemand(ctx, "rg1", "appgw", armnetwork.ApplicationGatewayOnDemandProbe{
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
	// TODO: use response item
	_ = res
}
