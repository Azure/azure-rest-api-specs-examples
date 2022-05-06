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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteAdd.json
func ExampleFavoritesClient_Add() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Add(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<favorite-id>",
		armapplicationinsights.ComponentFavorite{
			Config:                  to.Ptr("<config>"),
			FavoriteID:              to.Ptr("<favorite-id>"),
			FavoriteType:            to.Ptr(armapplicationinsights.FavoriteTypeShared),
			IsGeneratedFromTemplate: to.Ptr(false),
			Name:                    to.Ptr("<name>"),
			Tags: []*string{
				to.Ptr("TagSample01"),
				to.Ptr("TagSample02")},
			Version: to.Ptr("<version>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
