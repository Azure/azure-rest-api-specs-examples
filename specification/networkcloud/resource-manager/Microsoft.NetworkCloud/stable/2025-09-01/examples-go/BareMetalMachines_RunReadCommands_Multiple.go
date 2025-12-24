package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/46c51b03d99b113ecc3b38883e3cb2d395fe94a4/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/BareMetalMachines_RunReadCommands_Multiple.json
func ExampleBareMetalMachinesClient_BeginRunReadCommands_runAndRetrieveOutputFromReadOnlyCommandsOnBareMetalMachine() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBareMetalMachinesClient().BeginRunReadCommands(ctx, "resourceGroupName", "bareMetalMachineName", armnetworkcloud.BareMetalMachineRunReadCommandsParameters{
		LimitTimeSeconds: to.Ptr[int64](60),
		Commands: []*armnetworkcloud.BareMetalMachineCommandSpecification{
			{
				Arguments: []*string{
					to.Ptr("pods"),
					to.Ptr("-A")},
				Command: to.Ptr("kubectl get"),
			},
			{
				Arguments: []*string{
					to.Ptr("192.168.0.99"),
					to.Ptr("-c"),
					to.Ptr("3")},
				Command: to.Ptr("ping"),
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
