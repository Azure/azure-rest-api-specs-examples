```go
package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-11-16-preview/examples/PutAccessReviewHistoryDefinition.json
func ExampleAccessReviewHistoryDefinitionClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewAccessReviewHistoryDefinitionClient("129a304b-4aea-4b86-a9f7-ba7e2b23737a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"44724910-d7a5-4c29-b28f-db73e717165a",
		armauthorization.AccessReviewHistoryDefinitionProperties{},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fauthorization%2Farmauthorization%2Fv2.0.0-beta.1/sdk/resourcemanager/authorization/armauthorization/README.md) on how to add the SDK to your project and authenticate.
