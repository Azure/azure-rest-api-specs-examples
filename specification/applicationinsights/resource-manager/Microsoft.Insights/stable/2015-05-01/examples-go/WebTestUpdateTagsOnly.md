Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.5.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestUpdateTagsOnly.json
func ExampleWebTestsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWebTestsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateTags(ctx,
		"my-resource-group",
		"my-webtest-my-component",
		armapplicationinsights.TagsResource{
			Tags: map[string]*string{
				"Color":          to.Ptr("AzureBlue"),
				"CustomField-01": to.Ptr("This is a random value"),
				"SystemType":     to.Ptr("A08"),
				"hidden-link:/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component": to.Ptr("Resource"),
				"hidden-link:/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Web/sites/mytestwebapp":           to.Ptr("Resource"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
