Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationalinsights%2Farmoperationalinsights%2Fv0.3.0/sdk/resourcemanager/operationalinsights/armoperationalinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
)

// x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesSavedSearchesCreateOrUpdate.json
func ExampleSavedSearchesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armoperationalinsights.NewSavedSearchesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<saved-search-id>",
		armoperationalinsights.SavedSearch{
			Properties: &armoperationalinsights.SavedSearchProperties{
				Category:           to.StringPtr("<category>"),
				DisplayName:        to.StringPtr("<display-name>"),
				FunctionAlias:      to.StringPtr("<function-alias>"),
				FunctionParameters: to.StringPtr("<function-parameters>"),
				Query:              to.StringPtr("<query>"),
				Tags: []*armoperationalinsights.Tag{
					{
						Name:  to.StringPtr("<name>"),
						Value: to.StringPtr("<value>"),
					}},
				Version: to.Int64Ptr(2),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SavedSearchesClientCreateOrUpdateResult)
}
```
