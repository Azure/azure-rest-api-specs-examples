package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/listApplicationsByResourceGroup.json
func ExampleApplicationsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationsClient().NewListByResourceGroupPager("rg", nil)
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
		// page.ApplicationListResult = armmanagedapplications.ApplicationListResult{
		// 	Value: []*armmanagedapplications.Application{
		// 		{
		// 			Name: to.Ptr("myManagedApplication"),
		// 			Type: to.Ptr("Microsoft.Solutions/applications"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication"),
		// 			Location: to.Ptr("East US 2"),
		// 			Kind: to.Ptr("ServiceCatalog"),
		// 			Properties: &armmanagedapplications.ApplicationProperties{
		// 				ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
		// 				ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myManagedApplication2"),
		// 			Type: to.Ptr("Microsoft.Solutions/applications"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication2"),
		// 			Location: to.Ptr("West US"),
		// 			Kind: to.Ptr("ServiceCatalog"),
		// 			Properties: &armmanagedapplications.ApplicationProperties{
		// 				ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
		// 				ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
