Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresourcemover%2Farmresourcemover%2Fv0.2.1/sdk/resourcemanager/resourcemover/armresourcemover/README.md) on how to add the SDK to your project and authenticate.

```go
package armresourcemover_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
)

// x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2021-08-01/examples/MoveResources_Create.json
func ExampleMoveResourcesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armresourcemover.NewMoveResourcesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<move-collection-name>",
		"<move-resource-name>",
		&armresourcemover.MoveResourcesClientBeginCreateOptions{Body: &armresourcemover.MoveResource{
			Properties: &armresourcemover.MoveResourceProperties{
				DependsOnOverrides: []*armresourcemover.MoveResourceDependencyOverride{
					{
						ID:       to.StringPtr("<id>"),
						TargetID: to.StringPtr("<target-id>"),
					}},
				ResourceSettings: &armresourcemover.VirtualMachineResourceSettings{
					ResourceType:            to.StringPtr("<resource-type>"),
					TargetResourceName:      to.StringPtr("<target-resource-name>"),
					TargetAvailabilitySetID: to.StringPtr("<target-availability-set-id>"),
					TargetAvailabilityZone:  armresourcemover.TargetAvailabilityZone("2").ToPtr(),
					UserManagedIdentities: []*string{
						to.StringPtr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1")},
				},
				SourceID: to.StringPtr("<source-id>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MoveResourcesClientCreateResult)
}
```
