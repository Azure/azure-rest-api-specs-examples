Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fadvisor%2Farmadvisor%2Fv1.0.0/sdk/resourcemanager/advisor/armadvisor/README.md) on how to add the SDK to your project and authenticate.

```go
package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/EmptyResponse.json
func ExampleRecommendationsClient_GetGenerateStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armadvisor.NewRecommendationsClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.GetGenerateStatus(ctx,
		"operationGUID",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```
