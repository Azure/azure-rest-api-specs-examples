Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv1.0.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/BastionShareableLinkGet.json
func ExampleManagementClient_NewGetBastionShareableLinkPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewManagementClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewGetBastionShareableLinkPager("rg1",
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
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
