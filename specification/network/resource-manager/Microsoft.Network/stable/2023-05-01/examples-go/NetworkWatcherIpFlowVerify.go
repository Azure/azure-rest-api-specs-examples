package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/NetworkWatcherIpFlowVerify.json
func ExampleWatchersClient_BeginVerifyIPFlow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginVerifyIPFlow(ctx, "rg1", "nw1", armnetwork.VerificationIPFlowParameters{
		Direction:        to.Ptr(armnetwork.DirectionOutbound),
		LocalIPAddress:   to.Ptr("10.2.0.4"),
		LocalPort:        to.Ptr("80"),
		RemoteIPAddress:  to.Ptr("121.10.1.1"),
		RemotePort:       to.Ptr("80"),
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
		Protocol:         to.Ptr(armnetwork.IPFlowProtocolTCP),
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
	// res.VerificationIPFlowResult = armnetwork.VerificationIPFlowResult{
	// 	Access: to.Ptr(armnetwork.AccessAllow),
	// 	RuleName: to.Ptr("Rule1"),
	// }
}
