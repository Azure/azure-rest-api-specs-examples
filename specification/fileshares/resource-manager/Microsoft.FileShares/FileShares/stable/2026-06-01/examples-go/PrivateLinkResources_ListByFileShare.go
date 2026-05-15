package armfileshares_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fileshares/armfileshares"
)

// Generated from example definition: 2026-06-01/PrivateLinkResources_ListByFileShare.json
func ExamplePrivateLinkResourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfileshares.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListPager("res4303", "testfileshare01", nil)
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
		// page = armfileshares.PrivateLinkResourcesClientListResponse{
		// 	PrivateLinkResourceListResult: armfileshares.PrivateLinkResourceListResult{
		// 		Value: []*armfileshares.PrivateLinkResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/res4303/providers/Microsoft.FileShares/fileShares/testfileshare01/privateLinkResources/fileshare"),
		// 				Name: to.Ptr("fileshare"),
		// 				Type: to.Ptr("Microsoft.FileShares/privateLinkResources"),
		// 				Properties: &armfileshares.PrivateLinkResourceProperties{
		// 					GroupID: to.Ptr("fileshare"),
		// 					RequiredMembers: []*string{
		// 						to.Ptr("file"),
		// 					},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.file.core.windows.net"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
