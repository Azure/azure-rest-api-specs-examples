package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/hunts/GetHuntComments.json
func ExampleHuntCommentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("bd794837-4d29-4647-9105-6339bfdb4e6a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHuntCommentsClient().NewListPager("myRg", "myWorkspace", "163e7b2a-a2ec-4041-aaba-d878a38f265f", nil)
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
		// page = armsecurityinsights.HuntCommentsClientListResponse{
		// 	HuntCommentList: armsecurityinsights.HuntCommentList{
		// 		Value: []*armsecurityinsights.HuntComment{
		// 			{
		// 				Name: to.Ptr("2216d0e1-91e3-4902-89fd-d2df8c123456"),
		// 				Type: to.Ptr("Microsoft.SecurityInsights/hunts/comments"),
		// 				Etag: to.Ptr("\"3102f74d-0000-0c00-0000-629e6e050000\""),
		// 				ID: to.Ptr("/subscriptions/bd794837-4d29-4647-9105-6339bfdb4e6a/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/workspaces/avdvirint/providers/Microsoft.SecurityInsights/hunts/163e7b2a-a2ec-4041-aaba-d878a38f265f/comments/2216d0e1-91e3-4902-89fd-d2df8c123456"),
		// 				Properties: &armsecurityinsights.HuntCommentProperties{
		// 					Message: to.Ptr("This is a test comment."),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
