package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/hunts/GetHuntRelations.json
func ExampleHuntRelationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("bd794837-4d29-4647-9105-6339bfdb4e6a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHuntRelationsClient().NewListPager("myRg", "myWorkspace", "163e7b2a-a2ec-4041-aaba-d878a38f265f", nil)
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
		// page = armsecurityinsights.HuntRelationsClientListResponse{
		// 	HuntRelationList: armsecurityinsights.HuntRelationList{
		// 		Value: []*armsecurityinsights.HuntRelation{
		// 			{
		// 				Name: to.Ptr("2216d0e1-91e3-4902-89fd-d2df8c535096"),
		// 				Type: to.Ptr("Microsoft.SecurityInsights/hunts/relations"),
		// 				ID: to.Ptr("/subscriptions/bd794837-4d29-4647-9105-6339bfdb4e6a/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/workspaces/avdvirint/providers/Microsoft.SecurityInsights/hunts/163e7b2a-a2ec-4041-aaba-d878a38f265f/relations/2216d0e1-91e3-4902-89fd-d2df8c535096"),
		// 				Properties: &armsecurityinsights.HuntRelationProperties{
		// 					Labels: []*string{
		// 						to.Ptr("label1"),
		// 					},
		// 					RelatedResourceID: to.Ptr("/subscriptions/bd794837-4d29-4647-9105-6339bfdb4e6a/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/workspaces/avdvirint/providers/Microsoft.SecurityInsights/Bookmarks/2216d0e1-91e3-4902-89fd-d2df8c535096"),
		// 					RelatedResourceName: to.Ptr("2216d0e1-91e3-4902-89fd-d2df8c535096"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
