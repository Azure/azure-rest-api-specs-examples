package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/bookmarks/CreateBookmark.json
func ExampleBookmarksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewBookmarksClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myRg",
		"myWorkspace",
		"73e01a99-5cd7-4139-a149-9f2736ff2ab5",
		armsecurityinsights.Bookmark{
			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
			Properties: &armsecurityinsights.BookmarkProperties{
				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30Z"); return t }()),
				CreatedBy: &armsecurityinsights.UserInfo{
					ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
				},
				DisplayName: to.Ptr("My bookmark"),
				Labels: []*string{
					to.Ptr("Tag1"),
					to.Ptr("Tag2")},
				Notes:       to.Ptr("Found a suspicious activity"),
				Query:       to.Ptr("SecurityEvent | where TimeGenerated > ago(1d) and TimeGenerated < ago(2d)"),
				QueryResult: to.Ptr("Security Event query result"),
				Updated:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30Z"); return t }()),
				UpdatedBy: &armsecurityinsights.UserInfo{
					ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
