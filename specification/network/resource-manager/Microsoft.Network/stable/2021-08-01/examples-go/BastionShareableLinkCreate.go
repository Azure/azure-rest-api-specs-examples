package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/BastionShareableLinkCreate.json
func ExampleManagementClient_BeginPutBastionShareableLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewManagementClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPutBastionShareableLink(ctx,
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
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	for res.More() {
		nextResult, err := res.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
