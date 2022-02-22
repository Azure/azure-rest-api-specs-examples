Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv0.1.1/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/bookmarks/CreateBookmark.json
func ExampleBookmarksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewBookmarksClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<bookmark-id>",
		armsecurityinsights.Bookmark{
			Etag: to.StringPtr("<etag>"),
			Properties: &armsecurityinsights.BookmarkProperties{
				Created: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:15:30Z"); return t }()),
				CreatedBy: &armsecurityinsights.UserInfo{
					ObjectID: to.StringPtr("<object-id>"),
				},
				DisplayName: to.StringPtr("<display-name>"),
				EntityMappings: []*armsecurityinsights.BookmarkEntityMappings{
					{
						EntityType: to.StringPtr("<entity-type>"),
						FieldMappings: []*armsecurityinsights.EntityFieldMapping{
							{
								Identifier: to.StringPtr("<identifier>"),
								Value:      to.StringPtr("<value>"),
							}},
					}},
				Labels: []*string{
					to.StringPtr("Tag1"),
					to.StringPtr("Tag2")},
				Notes:       to.StringPtr("<notes>"),
				Query:       to.StringPtr("<query>"),
				QueryResult: to.StringPtr("<query-result>"),
				Tactics: []*armsecurityinsights.AttackTactic{
					armsecurityinsights.AttackTactic("Execution").ToPtr()},
				Techniques: []*string{
					to.StringPtr("T1609")},
				Updated: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:15:30Z"); return t }()),
				UpdatedBy: &armsecurityinsights.UserInfo{
					ObjectID: to.StringPtr("<object-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BookmarksClientCreateOrUpdateResult)
}
```
