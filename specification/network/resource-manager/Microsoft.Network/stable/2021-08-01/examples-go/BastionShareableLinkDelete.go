package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/BastionShareableLinkDelete.json
func ExampleManagementClient_BeginDeleteBastionShareableLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewManagementClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDeleteBastionShareableLink(ctx,
		"rg1",
		"bastionhosttenant",
		armnetwork.BastionShareableLinkListRequest{
			VMs: []*armnetwork.BastionShareableLink{
				{
					VM: &armnetwork.VM{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm1"),
					},
				},
				{
					VM: &armnetwork.VM{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm2"),
					},
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
