Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresourcemover%2Farmresourcemover%2Fv0.2.1/sdk/resourcemanager/resourcemover/armresourcemover/README.md) on how to add the SDK to your project and authenticate.

```go
package armresourcemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
)

// x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2021-08-01/examples/MoveCollections_Create.json
func ExampleMoveCollectionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armresourcemover.NewMoveCollectionsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<move-collection-name>",
		&armresourcemover.MoveCollectionsClientCreateOptions{Body: &armresourcemover.MoveCollection{
			Identity: &armresourcemover.Identity{
				Type: armresourcemover.ResourceIdentityType("SystemAssigned").ToPtr(),
			},
			Location: to.StringPtr("<location>"),
			Properties: &armresourcemover.MoveCollectionProperties{
				SourceRegion: to.StringPtr("<source-region>"),
				TargetRegion: to.StringPtr("<target-region>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MoveCollectionsClientCreateResult)
}
```
