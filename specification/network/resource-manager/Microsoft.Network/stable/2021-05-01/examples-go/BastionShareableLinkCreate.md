Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/BastionShareableLinkCreate.json
func ExampleManagementClient_BeginPutBastionShareableLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewManagementClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPutBastionShareableLink(ctx,
		"<resource-group-name>",
		"<bastion-host-name>",
		armnetwork.BastionShareableLinkListRequest{
			VMs: []*armnetwork.BastionShareableLink{
				{
					VM: &armnetwork.VM{
						ID: to.StringPtr("<id>"),
					},
				},
				{
					VM: &armnetwork.VM{
						ID: to.StringPtr("<id>"),
					},
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	for {
		nextResult := res.NextPage(ctx)
		if err := res.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range res.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
```
