Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.2.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookAdd.json
func ExampleWorkbooksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewWorkbooksClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.Workbook{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"TagSample01": to.StringPtr("sample01"),
				"TagSample02": to.StringPtr("sample02"),
			},
			Kind: armapplicationinsights.Kind("shared").ToPtr(),
			Properties: &armapplicationinsights.WorkbookProperties{
				Description:    to.StringPtr("<description>"),
				Category:       to.StringPtr("<category>"),
				DisplayName:    to.StringPtr("<display-name>"),
				SerializedData: to.StringPtr("<serialized-data>"),
			},
		},
		&armapplicationinsights.WorkbooksClientCreateOrUpdateOptions{SourceID: to.StringPtr("<source-id>")})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkbooksClientCreateOrUpdateResult)
}
```
