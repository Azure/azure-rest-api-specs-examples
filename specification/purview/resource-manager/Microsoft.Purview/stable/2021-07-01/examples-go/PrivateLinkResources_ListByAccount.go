package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateLinkResources_ListByAccount.json
func ExamplePrivateLinkResourcesClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByAccountPager("SampleResourceGroup", "account1", nil)
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
		// page.PrivateLinkResourceList = armpurview.PrivateLinkResourceList{
		// 	Value: []*armpurview.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("plr1"),
		// 			Type: to.Ptr("Microsoft.Purview/accounts/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1/privateLinkResources/plr1"),
		// 			Properties: &armpurview.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("account1"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("account1")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelinkzone1.service.azure.com")},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("plr2"),
		// 					Type: to.Ptr("Microsoft.Purview/accounts/privateLinkResources"),
		// 					ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1/privateLinkResources/plr2"),
		// 					Properties: &armpurview.PrivateLinkResourceProperties{
		// 						GroupID: to.Ptr("account2"),
		// 						RequiredMembers: []*string{
		// 							to.Ptr("account2")},
		// 							RequiredZoneNames: []*string{
		// 								to.Ptr("privatelinkzone2.service.azure.com")},
		// 							},
		// 					}},
		// 				}
	}
}
