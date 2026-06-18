package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/UnifiedResilienceItems_List_MinimumSet_Gen.json
func ExampleUnifiedResilienceItemsClient_NewListPager_unifiedResilienceItemsListMaximumSetGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUnifiedResilienceItemsClient().NewListPager("sampleServiceGroupName", nil)
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
		// page = armresiliencemanagement.UnifiedResilienceItemsClientListResponse{
		// 	UnifiedResilienceItemListResult: armresiliencemanagement.UnifiedResilienceItemListResult{
		// 		Value: []*armresiliencemanagement.UnifiedResilienceItem{
		// 			{
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/unifiedResilienceItems/sample-name-1234"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
