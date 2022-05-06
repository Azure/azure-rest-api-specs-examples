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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookAdd.json
func ExampleWorkbooksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWorkbooksClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.Workbook{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"TagSample01": to.Ptr("sample01"),
				"TagSample02": to.Ptr("sample02"),
			},
			Kind: to.Ptr(armapplicationinsights.KindShared),
			Properties: &armapplicationinsights.WorkbookProperties{
				Description:    to.Ptr("<description>"),
				Category:       to.Ptr("<category>"),
				DisplayName:    to.Ptr("<display-name>"),
				SerializedData: to.Ptr("<serialized-data>"),
			},
		},
		&armapplicationinsights.WorkbooksClientCreateOrUpdateOptions{SourceID: to.Ptr("<source-id>")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
