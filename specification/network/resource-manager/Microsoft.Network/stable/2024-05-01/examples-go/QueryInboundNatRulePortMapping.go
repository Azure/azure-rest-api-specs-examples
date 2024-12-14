package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/QueryInboundNatRulePortMapping.json
func ExampleLoadBalancersClient_BeginListInboundNatRulePortMappings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLoadBalancersClient().BeginListInboundNatRulePortMappings(ctx, "rg1", "lb1", "bp1", armnetwork.QueryInboundNatRulePortMappingRequest{
		IPAddress: to.Ptr("10.0.0.4"),
	}, nil)
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
	// res.BackendAddressInboundNatRulePortMappings = armnetwork.BackendAddressInboundNatRulePortMappings{
	// 	InboundNatRulePortMappings: []*armnetwork.InboundNatRulePortMapping{
	// 		{
	// 			BackendPort: to.Ptr[int32](3389),
	// 			FrontendPort: to.Ptr[int32](3389),
	// 			InboundNatRuleName: to.Ptr("natRule"),
	// 			Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
	// 	}},
	// }
}
