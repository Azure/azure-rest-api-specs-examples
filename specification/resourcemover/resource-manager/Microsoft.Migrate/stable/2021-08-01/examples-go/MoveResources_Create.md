Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresourcemover%2Farmresourcemover%2Fv0.4.0/sdk/resourcemanager/resourcemover/armresourcemover/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2021-08-01/examples/MoveResources_Create.json
func ExampleMoveResourcesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armresourcemover.NewMoveResourcesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<move-collection-name>",
		"<move-resource-name>",
		&armresourcemover.MoveResourcesClientBeginCreateOptions{Body: &armresourcemover.MoveResource{
			Properties: &armresourcemover.MoveResourceProperties{
				DependsOnOverrides: []*armresourcemover.MoveResourceDependencyOverride{
					{
						ID:       to.Ptr("<id>"),
						TargetID: to.Ptr("<target-id>"),
					}},
				ResourceSettings: &armresourcemover.VirtualMachineResourceSettings{
					ResourceType:            to.Ptr("<resource-type>"),
					TargetResourceName:      to.Ptr("<target-resource-name>"),
					TargetAvailabilitySetID: to.Ptr("<target-availability-set-id>"),
					TargetAvailabilityZone:  to.Ptr(armresourcemover.TargetAvailabilityZoneTwo),
					UserManagedIdentities: []*string{
						to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1")},
				},
				SourceID: to.Ptr("<source-id>"),
			},
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
