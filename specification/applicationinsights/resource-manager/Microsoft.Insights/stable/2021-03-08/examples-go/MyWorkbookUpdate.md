Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.4.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-03-08/examples/MyWorkbookUpdate.json
func ExampleMyWorkbooksClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewMyWorkbooksClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.MyWorkbook{
			Name:     to.Ptr("<name>"),
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"0": to.Ptr("TagSample01"),
				"1": to.Ptr("TagSample02"),
			},
			Kind: to.Ptr(armapplicationinsights.KindUser),
			Properties: &armapplicationinsights.MyWorkbookProperties{
				Category:       to.Ptr("<category>"),
				DisplayName:    to.Ptr("<display-name>"),
				SerializedData: to.Ptr("<serialized-data>"),
				SourceID:       to.Ptr("<source-id>"),
				Version:        to.Ptr("<version>"),
			},
		},
		&armapplicationinsights.MyWorkbooksClientUpdateOptions{SourceID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
```
