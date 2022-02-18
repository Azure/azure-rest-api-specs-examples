Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv0.1.0/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/bookmarks/expand/PostExpandBookmark.json
func ExampleBookmarkClient_Expand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewBookmarkClient("<subscription-id>", cred, nil)
	res, err := client.Expand(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<bookmark-id>",
		armsecurityinsights.BookmarkExpandParameters{
			EndTime:     to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-24T17:21:00.000Z"); return t }()),
			ExpansionID: to.StringPtr("<expansion-id>"),
			StartTime:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-25T17:21:00.000Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BookmarkClientExpandResult)
}
```
