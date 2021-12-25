Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.1.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/APIKeysCreate.json
func ExampleAPIKeysClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewAPIKeysClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.APIKeyRequest{
			Name: to.StringPtr("<name>"),
			LinkedReadProperties: []*string{
				to.StringPtr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/api"),
				to.StringPtr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/agentconfig")},
			LinkedWriteProperties: []*string{
				to.StringPtr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/annotations")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ApplicationInsightsComponentAPIKey.ID: %s\n", *res.ID)
}
```
