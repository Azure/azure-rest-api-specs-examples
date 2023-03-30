package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/InventoryItems_ListByVCenter.json
func ExampleInventoryItemsClient_NewListByVCenterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInventoryItemsClient().NewListByVCenterPager("testrg", "ContosoVCenter", nil)
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
		// page.InventoryItemsList = armconnectedvmware.InventoryItemsList{
		// 	Value: []*armconnectedvmware.InventoryItem{
		// 		{
		// 			Name: to.Ptr("testItem"),
		// 			Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VCenters/InventoryItems"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter/InventoryItems/testItem"),
		// 			Properties: &armconnectedvmware.ResourcePoolInventoryItem{
		// 				InventoryType: to.Ptr(armconnectedvmware.InventoryTypeResourcePool),
		// 			},
		// 	}},
		// }
	}
}
