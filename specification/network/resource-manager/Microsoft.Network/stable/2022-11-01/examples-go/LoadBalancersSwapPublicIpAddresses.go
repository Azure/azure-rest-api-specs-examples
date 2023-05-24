package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/LoadBalancersSwapPublicIpAddresses.json
func ExampleLoadBalancersClient_BeginSwapPublicIPAddresses() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLoadBalancersClient().BeginSwapPublicIPAddresses(ctx, "westus", armnetwork.LoadBalancerVipSwapRequest{
		FrontendIPConfigurations: []*armnetwork.LoadBalancerVipSwapRequestFrontendIPConfiguration{
			{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/lbfe1"),
				Properties: &armnetwork.LoadBalancerVipSwapRequestFrontendIPConfigurationProperties{
					PublicIPAddress: &armnetwork.SubResource{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/publicIPAddresses/pip2"),
					},
				},
			},
			{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/loadBalancers/lb2/frontendIPConfigurations/lbfe2"),
				Properties: &armnetwork.LoadBalancerVipSwapRequestFrontendIPConfigurationProperties{
					PublicIPAddress: &armnetwork.SubResource{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pip1"),
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
