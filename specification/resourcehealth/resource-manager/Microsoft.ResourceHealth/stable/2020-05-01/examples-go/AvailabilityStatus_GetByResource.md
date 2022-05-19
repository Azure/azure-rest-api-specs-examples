Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresourcehealth%2Farmresourcehealth%2Fv1.0.0/sdk/resourcemanager/resourcehealth/armresourcehealth/README.md) on how to add the SDK to your project and authenticate.

```go
package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2020-05-01/examples/AvailabilityStatus_GetByResource.json
func ExampleAvailabilityStatusesClient_GetByResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresourcehealth.NewAvailabilityStatusesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetByResource(ctx,
		"resourceUri",
		&armresourcehealth.AvailabilityStatusesClientGetByResourceOptions{Filter: nil,
			Expand: to.Ptr("recommendedactions"),
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
