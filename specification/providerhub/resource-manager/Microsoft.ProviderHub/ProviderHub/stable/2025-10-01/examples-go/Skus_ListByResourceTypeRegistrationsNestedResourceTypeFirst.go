package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/Skus_ListByResourceTypeRegistrationsNestedResourceTypeFirst.json
func ExampleSKUsClient_NewListByResourceTypeRegistrationsNestedResourceTypeFirstPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListByResourceTypeRegistrationsNestedResourceTypeFirstPager("Microsoft.Contoso", "testResourceType", "nestedResourceTypeFirst", nil)
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
		// page = armproviderhub.SKUsClientListByResourceTypeRegistrationsNestedResourceTypeFirstResponse{
		// 	SKUResourceArrayResponseWithContinuation: armproviderhub.SKUResourceArrayResponseWithContinuation{
		// 		Value: []*armproviderhub.SKUResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/resourcetypeRegistrations/employees/resourcetypeRegistrations/nestedEmployee/skus/sku1"),
		// 				Name: to.Ptr("Microsoft.Contoso/employees/nestedEmployee/sku1"),
		// 				Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/resourcetypeRegistrations/resourcetypeRegistrations/skus"),
		// 				Properties: &armproviderhub.SKUResourceProperties{
		// 					SKUSettings: []*armproviderhub.SKUSetting{
		// 						{
		// 							Name: to.Ptr("freeSku"),
		// 							Tier: to.Ptr("Tier1"),
		// 							Kind: to.Ptr("Standard"),
		// 						},
		// 						{
		// 							Name: to.Ptr("premiumSku"),
		// 							Tier: to.Ptr("Tier2"),
		// 							Kind: to.Ptr("Premium"),
		// 							Costs: []*armproviderhub.SKUCost{
		// 								{
		// 									MeterID: to.Ptr("xxx"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armproviderhub.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
