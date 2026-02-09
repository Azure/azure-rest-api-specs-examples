package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/BastionSessionsList.json
func ExampleManagementClient_BeginGetActiveSessions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagementClient().BeginGetActiveSessions(ctx, "rg1", "bastionhosttenant", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	for res.More() {
		page, err := res.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.BastionActiveSessionListResult = armnetwork.BastionActiveSessionListResult{
		// 	Value: []*armnetwork.BastionActiveSession{
		// 		{
		// 			ResourceType: to.Ptr("VM"),
		// 			SessionDurationInMins: to.Ptr[float32](0),
		// 			SessionID: to.Ptr("sessionId"),
		// 			StartTime: "2019-1-1T12:00:00.0000Z",
		// 			TargetHostName: to.Ptr("vm01"),
		// 			TargetIPAddress: to.Ptr("1.1.1.1"),
		// 			TargetResourceGroup: to.Ptr("rg1"),
		// 			TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm01"),
		// 			TargetSubscriptionID: to.Ptr("subid"),
		// 			UserName: to.Ptr("user"),
		// 			Protocol: to.Ptr(armnetwork.BastionConnectProtocolSSH),
		// 	}},
		// }
	}
}
