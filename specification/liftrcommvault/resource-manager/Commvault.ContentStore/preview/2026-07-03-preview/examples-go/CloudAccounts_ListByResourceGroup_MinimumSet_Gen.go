package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/CloudAccounts_ListByResourceGroup_MinimumSet_Gen.json
func ExampleCloudAccountsClient_NewListByResourceGroupPager_cloudAccountsListByResourceGroupMaximumSetGeneratedByMaximumSetRuleGeneratedByMaximumSetRuleGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("65D4E6D7-7063-4C4B-BAC5-13C45474009E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudAccountsClient().NewListByResourceGroupPager("rgcommvault", nil)
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
		// page = armcommvaultcontentstore.CloudAccountsClientListByResourceGroupResponse{
		// 	CloudAccountListResult: armcommvaultcontentstore.CloudAccountListResult{
		// 		Value: []*armcommvaultcontentstore.CloudAccount{
		// 			{
		// 				ID: to.Ptr("/subscriptions/65D4E6D7-7063-4C4B-BAC5-13C45474009E/resourceGroups/rgcommvault/providers/Commvault.ContentStore/cloudAccounts/sample-cloudAccountName"),
		// 				Location: to.Ptr("sxzmmidsfbba"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
