Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv0.3.0/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/bookmarks/CreateBookmark.json
func ExampleBookmarksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewBookmarksClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<bookmark-id>",
		armsecurityinsights.Bookmark{
			Etag: to.Ptr("<etag>"),
			Properties: &armsecurityinsights.BookmarkProperties{
				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:15:30Z"); return t }()),
				CreatedBy: &armsecurityinsights.UserInfo{
					ObjectID: to.Ptr("<object-id>"),
				},
				DisplayName: to.Ptr("<display-name>"),
				EntityMappings: []*armsecurityinsights.BookmarkEntityMappings{
					{
						EntityType: to.Ptr("<entity-type>"),
						FieldMappings: []*armsecurityinsights.EntityFieldMapping{
							{
								Identifier: to.Ptr("<identifier>"),
								Value:      to.Ptr("<value>"),
							}},
					}},
				Labels: []*string{
					to.Ptr("Tag1"),
					to.Ptr("Tag2")},
				Notes:       to.Ptr("<notes>"),
				Query:       to.Ptr("<query>"),
				QueryResult: to.Ptr("<query-result>"),
				Tactics: []*armsecurityinsights.AttackTactic{
					to.Ptr(armsecurityinsights.AttackTacticExecution)},
				Techniques: []*string{
					to.Ptr("T1609")},
				Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:15:30Z"); return t }()),
				UpdatedBy: &armsecurityinsights.UserInfo{
					ObjectID: to.Ptr("<object-id>"),
				},
			},
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
