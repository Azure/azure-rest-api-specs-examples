package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/EnforceAdaptiveNetworkHardeningRules_example.json
func ExampleAdaptiveNetworkHardeningsClient_BeginEnforce() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAdaptiveNetworkHardeningsClient().BeginEnforce(ctx, "rg1", "Microsoft.Compute", "virtualMachines", "vm1", "default", armsecurity.AdaptiveNetworkHardeningEnforceRequest{
		NetworkSecurityGroups: []*string{
			to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
			to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg2/providers/Microsoft.Network/networkSecurityGroups/nsg2")},
		Rules: []*armsecurity.Rule{
			{
				Name:            to.Ptr("rule1"),
				DestinationPort: to.Ptr[int32](3389),
				Direction:       to.Ptr(armsecurity.DirectionInbound),
				IPAddresses: []*string{
					to.Ptr("100.10.1.1"),
					to.Ptr("200.20.2.2"),
					to.Ptr("81.199.3.0/24")},
				Protocols: []*armsecurity.TransportProtocol{
					to.Ptr(armsecurity.TransportProtocolTCP)},
			},
			{
				Name:            to.Ptr("rule2"),
				DestinationPort: to.Ptr[int32](22),
				Direction:       to.Ptr(armsecurity.DirectionInbound),
				IPAddresses:     []*string{},
				Protocols: []*armsecurity.TransportProtocol{
					to.Ptr(armsecurity.TransportProtocolTCP)},
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
