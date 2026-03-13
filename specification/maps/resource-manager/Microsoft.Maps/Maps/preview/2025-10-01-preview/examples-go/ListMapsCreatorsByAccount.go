package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps/v2"
)

// Generated from example definition: 2025-10-01-preview/ListMapsCreatorsByAccount.json
func ExampleCreatorsClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("21a9967a-e8a9-4656-a70b-96ff1c4d05a0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCreatorsClient().NewListByAccountPager("myResourceGroup", "myMapsAccount", nil)
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
		// page = armmaps.CreatorsClientListByAccountResponse{
		// 	CreatorList: armmaps.CreatorList{
		// 		Value: []*armmaps.Creator{
		// 			{
		// 				Name: to.Ptr("myCreator"),
		// 				Type: to.Ptr("Microsoft.Maps/accounts/creators"),
		// 				ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount/creators/myCreator"),
		// 				Location: to.Ptr("eastus2"),
		// 				Properties: &armmaps.CreatorProperties{
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					StorageUnits: to.Ptr[int32](5),
		// 				},
		// 				Tags: map[string]*string{
		// 					"test": to.Ptr("true"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
