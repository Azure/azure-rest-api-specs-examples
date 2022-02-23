Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresourcemover%2Farmresourcemover%2Fv0.2.1/sdk/resourcemanager/resourcemover/armresourcemover/README.md) on how to add the SDK to your project and authenticate.

```go
package armresourcemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
)

// x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2021-08-01/examples/RequiredFor_Get.json
func ExampleMoveCollectionsClient_ListRequiredFor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armresourcemover.NewMoveCollectionsClient("<subscription-id>", cred, nil)
	res, err := client.ListRequiredFor(ctx,
		"<resource-group-name>",
		"<move-collection-name>",
		"<source-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MoveCollectionsClientListRequiredForResult)
}
```
