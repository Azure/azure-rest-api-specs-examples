package armqumulo_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/liftrqumulo/armqumulo/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72de08114673a547de8a017c85ed89a2017a86f7/specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/FileSystems_ListByResourceGroup_MinimumSet_Gen.json
func ExampleFileSystemsClient_NewListByResourceGroupPager_fileSystemsListByResourceGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSystemsClient().NewListByResourceGroupPager("rgQumulo", nil)
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
		// page.FileSystemResourceListResult = armqumulo.FileSystemResourceListResult{
		// 	Value: []*armqumulo.FileSystemResource{
		// 		{
		// 			Name: to.Ptr("aaaaa"),
		// 			ID: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 			Location: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			Properties: &armqumulo.FileSystemResourceProperties{
		// 				DelegatedSubnetID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				MarketplaceDetails: &armqumulo.MarketplaceDetails{
		// 					MarketplaceSubscriptionID: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 					MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 					OfferID: to.Ptr("aaaaaaaaa"),
		// 					PlanID: to.Ptr("aaaaaaa"),
		// 				},
		// 				ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
		// 				StorageSKU: to.Ptr("Standard"),
		// 				UserDetails: &armqumulo.UserDetails{
		// 				},
		// 			},
		// 	}},
		// }
	}
}
