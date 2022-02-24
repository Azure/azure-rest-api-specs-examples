Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagementgroups%2Farmmanagementgroups%2Fv0.3.1/sdk/resourcemanager/managementgroups/armmanagementgroups/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagementgroups_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
)

// x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PutManagementGroup.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagementgroups.NewClient(cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<group-id>",
		armmanagementgroups.CreateManagementGroupRequest{
			Properties: &armmanagementgroups.CreateManagementGroupProperties{
				DisplayName: to.StringPtr("<display-name>"),
				Details: &armmanagementgroups.CreateManagementGroupDetails{
					Parent: &armmanagementgroups.CreateParentGroupInfo{
						ID: to.StringPtr("<id>"),
					},
				},
			},
		},
		&armmanagementgroups.ClientBeginCreateOrUpdateOptions{CacheControl: to.StringPtr("<cache-control>")})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClientCreateOrUpdateResult)
}
```
