package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/bookmarks/CreateBookmark.json
func ExampleBookmarksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBookmarksClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", armsecurityinsights.Bookmark{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Properties: &armsecurityinsights.BookmarkProperties{
			Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:15:30.000Z"); return t }()),
			CreatedBy: &armsecurityinsights.UserInfo{
				ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
			},
			DisplayName: to.Ptr("My bookmark"),
			EntityMappings: []*armsecurityinsights.BookmarkEntityMappings{
				{
					EntityType: to.Ptr("Account"),
					FieldMappings: []*armsecurityinsights.EntityFieldMapping{
						{
							Identifier: to.Ptr("Fullname"),
							Value:      to.Ptr("johndoe@microsoft.com"),
						}},
				}},
			Labels: []*string{
				to.Ptr("Tag1"),
				to.Ptr("Tag2")},
			Notes:       to.Ptr("Found a suspicious activity"),
			Query:       to.Ptr("SecurityEvent | where TimeGenerated > ago(1d) and TimeGenerated < ago(2d)"),
			QueryResult: to.Ptr("Security Event query result"),
			Tactics: []*armsecurityinsights.AttackTactic{
				to.Ptr(armsecurityinsights.AttackTacticExecution)},
			Techniques: []*string{
				to.Ptr("T1609")},
			Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:15:30.000Z"); return t }()),
			UpdatedBy: &armsecurityinsights.UserInfo{
				ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Bookmark = armsecurityinsights.Bookmark{
	// 	Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/bookmarks"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/bookmarks/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 	Properties: &armsecurityinsights.BookmarkProperties{
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:15:30.000Z"); return t}()),
	// 		CreatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		DisplayName: to.Ptr("My bookmark"),
	// 		EntityMappings: []*armsecurityinsights.BookmarkEntityMappings{
	// 			{
	// 				EntityType: to.Ptr("Account"),
	// 				FieldMappings: []*armsecurityinsights.EntityFieldMapping{
	// 					{
	// 						Identifier: to.Ptr("Fullname"),
	// 						Value: to.Ptr("johndoe@microsoft.com"),
	// 				}},
	// 		}},
	// 		Labels: []*string{
	// 			to.Ptr("Tag1"),
	// 			to.Ptr("Tag2")},
	// 			Notes: to.Ptr("Found a suspicious activity"),
	// 			Query: to.Ptr("SecurityEvent | where TimeGenerated > ago(1d) and TimeGenerated < ago(2d)"),
	// 			QueryResult: to.Ptr("Security Event query result"),
	// 			Tactics: []*armsecurityinsights.AttackTactic{
	// 				to.Ptr(armsecurityinsights.AttackTacticExecution)},
	// 				Techniques: []*string{
	// 					to.Ptr("T1609")},
	// 					Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:15:30.000Z"); return t}()),
	// 					UpdatedBy: &armsecurityinsights.UserInfo{
	// 						Name: to.Ptr("john doe"),
	// 						Email: to.Ptr("john@contoso.com"),
	// 						ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 					},
	// 				},
	// 			}
}
