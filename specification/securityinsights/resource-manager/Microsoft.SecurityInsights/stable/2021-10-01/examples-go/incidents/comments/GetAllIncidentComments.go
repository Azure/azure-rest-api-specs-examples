package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/comments/GetAllIncidentComments.json
func ExampleIncidentCommentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIncidentCommentsClient().NewListPager("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", &armsecurityinsights.IncidentCommentsClientListOptions{Filter: nil,
		Orderby:   nil,
		Top:       nil,
		SkipToken: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.IncidentCommentList = armsecurityinsights.IncidentCommentList{
		// 	Value: []*armsecurityinsights.IncidentComment{
		// 		{
		// 			Name: to.Ptr("4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/incidents/comments"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/73e01a99-5cd7-4139-a149-9f2736ff2ab5/comments/4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014"),
		// 			Etag: to.Ptr("0300bf09-0000-0000-0000-5c37296e0000"),
		// 			Properties: &armsecurityinsights.IncidentCommentProperties{
		// 				Author: &armsecurityinsights.ClientInfo{
		// 					Name: to.Ptr("john doe"),
		// 					Email: to.Ptr("john.doe@contoso.com"),
		// 					ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 					UserPrincipalName: to.Ptr("john@contoso.com"),
		// 				},
		// 				CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30.000Z"); return t}()),
		// 				LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30.000Z"); return t}()),
		// 				Message: to.Ptr("Some message"),
		// 			},
		// 	}},
		// }
	}
}
