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

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-03-08/examples/MyWorkbookUpdate.json
func ExampleMyWorkbooksClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewMyWorkbooksClient("<subscription-id>", cred, nil)
	_, err = client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.MyWorkbook{
			Name:     to.StringPtr("<name>"),
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"0": to.StringPtr("TagSample01"),
				"1": to.StringPtr("TagSample02"),
			},
			Kind: armapplicationinsights.Kind("user").ToPtr(),
			Properties: &armapplicationinsights.MyWorkbookProperties{
				Category:       to.StringPtr("<category>"),
				DisplayName:    to.StringPtr("<display-name>"),
				SerializedData: to.StringPtr("<serialized-data>"),
				SourceID:       to.StringPtr("<source-id>"),
				Version:        to.StringPtr("<version>"),
			},
		},
		&armapplicationinsights.MyWorkbooksClientUpdateOptions{SourceID: nil})
	if err != nil {
		log.Fatal(err)
	}
}
```
