```go
package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-11-16-preview/examples/PatchAccessReviewInstanceMyDecisionById.json
func ExampleAccessReviewInstanceMyDecisionsClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewAccessReviewInstanceMyDecisionsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Patch(ctx,
		"488a6d0e-0a63-4946-86e3-1f5bbc934661",
		"4135f961-be78-4005-8101-c72a5af307a2",
		"fa73e90b-5bf1-45fd-a182-35ce5fc0674d",
		armauthorization.AccessReviewDecisionProperties{},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fauthorization%2Farmauthorization%2Fv2.0.0-beta.1/sdk/resourcemanager/authorization/armauthorization/README.md) on how to add the SDK to your project and authenticate.
