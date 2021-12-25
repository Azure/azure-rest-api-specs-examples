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

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteAdd.json
func ExampleFavoritesClient_Add() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewFavoritesClient("<subscription-id>", cred, nil)
	_, err = client.Add(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<favorite-id>",
		armapplicationinsights.ApplicationInsightsComponentFavorite{
			Category:                to.StringPtr("<category>"),
			Config:                  to.StringPtr("<config>"),
			FavoriteID:              to.StringPtr("<favorite-id>"),
			FavoriteType:            armapplicationinsights.FavoriteTypeShared.ToPtr(),
			IsGeneratedFromTemplate: to.BoolPtr(false),
			Name:                    to.StringPtr("<name>"),
			SourceType:              to.StringPtr("<source-type>"),
			Tags: []*string{
				to.StringPtr("TagSample01"),
				to.StringPtr("TagSample02")},
			TimeModified: to.StringPtr("<time-modified>"),
			Version:      to.StringPtr("<version>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
