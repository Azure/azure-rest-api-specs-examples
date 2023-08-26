package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/BareMetalMachines_RunReadCommands.json
func ExampleBareMetalMachinesClient_BeginRunReadCommands() {
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
