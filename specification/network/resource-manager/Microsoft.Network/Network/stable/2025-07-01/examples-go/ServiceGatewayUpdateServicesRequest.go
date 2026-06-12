package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ServiceGatewayUpdateServicesRequest.json
func ExampleServiceGatewaysClient_BeginUpdateServices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/be1"),
							},
						},
						PublicNatGatewayID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/natGateways/test-natGateway"),
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
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
