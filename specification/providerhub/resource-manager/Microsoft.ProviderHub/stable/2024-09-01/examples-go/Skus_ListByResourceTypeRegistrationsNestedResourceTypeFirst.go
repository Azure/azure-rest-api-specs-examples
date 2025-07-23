package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/Skus_ListByResourceTypeRegistrationsNestedResourceTypeFirst.json
func ExampleSKUsClient_NewListByResourceTypeRegistrationsNestedResourceTypeFirstPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.SKUResourceArrayResponseWithContinuation = armproviderhub.SKUResourceArrayResponseWithContinuation{
		// 	Value: []*armproviderhub.SKUResource{
		// 		{
		// 			Name: to.Ptr("Microsoft.Contoso/employees/nestedEmployee/sku1"),
		// 			Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/resourcetypeRegistrations/resourcetypeRegistrations/skus"),
		// 			ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/resourcetypeRegistrations/employees/resourcetypeRegistrations/nestedEmployee/skus/sku1"),
		// 			SystemData: &armproviderhub.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 			},
		// 			Properties: &armproviderhub.SKUResourceProperties{
		// 				SKUSettings: []*armproviderhub.SKUSetting{
		// 					{
		// 						Name: to.Ptr("freeSku"),
		// 						Kind: to.Ptr("Standard"),
		// 						Tier: to.Ptr("Tier1"),
		// 					},
		// 					{
		// 						Name: to.Ptr("premiumSku"),
		// 						Costs: []*armproviderhub.SKUCost{
		// 							{
		// 								MeterID: to.Ptr("xxx"),
		// 						}},
		// 						Kind: to.Ptr("Premium"),
		// 						Tier: to.Ptr("Tier2"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
