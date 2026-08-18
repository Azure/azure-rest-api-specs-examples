package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/ServiceGatewayUpdateServicesRequest.json
func ExampleServiceGatewaysClient_UpdateServices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceGatewaysClient().UpdateServices(ctx, "rg1", "sg", armnetwork.ServiceGatewayUpdateServicesRequest{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ServiceGatewaysClientUpdateServicesResponse{
	// 	ServiceGatewayActionOkResponseBody: armnetwork.ServiceGatewayActionOkResponseBody{
	// 	},
	// }
}
