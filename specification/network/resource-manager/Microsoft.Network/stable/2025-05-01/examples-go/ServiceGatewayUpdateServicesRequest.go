package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ServiceGatewayUpdateServicesRequest.json
func ExampleServiceGatewaysClient_BeginUpdateServices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceGatewaysClient().BeginUpdateServices(ctx, "rg1", "sg", armnetwork.ServiceGatewayUpdateServicesRequest{
		Action: to.Ptr(armnetwork.ServiceUpdateActionFullUpdate),
		ServiceRequests: []*armnetwork.ServiceGatewayServiceRequest{
			{
				Service: &armnetwork.ServiceGatewayService{
					Name: to.Ptr("Service1"),
					Properties: &armnetwork.ServiceGatewayServicePropertiesFormat{
						IsDefault: to.Ptr(true),
						LoadBalancerBackendPools: []*armnetwork.BackendAddressPool{
							{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/be1"),
							}},
						PublicNatGatewayID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/natGateways/test-natGateway"),
						ServiceType:        to.Ptr(armnetwork.ServiceTypeInbound),
					},
				},
			},
			{
				IsDelete: to.Ptr(true),
				Service: &armnetwork.ServiceGatewayService{
					Name: to.Ptr("Service2"),
					Properties: &armnetwork.ServiceGatewayServicePropertiesFormat{
						IsDefault:   to.Ptr(false),
						ServiceType: to.Ptr(armnetwork.ServiceTypeOutbound),
					},
				},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
