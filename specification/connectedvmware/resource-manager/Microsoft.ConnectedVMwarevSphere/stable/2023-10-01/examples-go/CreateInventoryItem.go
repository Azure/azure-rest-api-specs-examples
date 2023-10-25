package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/CreateInventoryItem.json
func ExampleInventoryItemsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInventoryItemsClient().Create(ctx, "testrg", "ContosoVCenter", "testItem", armconnectedvmware.InventoryItem{
		Properties: &armconnectedvmware.ResourcePoolInventoryItem{
			InventoryType: to.Ptr(armconnectedvmware.InventoryTypeResourcePool),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InventoryItem = armconnectedvmware.InventoryItem{
	// 	Name: to.Ptr("testItem"),
	// 	Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VCenters/InventoryItems"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter/InventoryItems/testItem"),
	// 	Properties: &armconnectedvmware.ResourcePoolInventoryItem{
	// 		InventoryType: to.Ptr(armconnectedvmware.InventoryTypeResourcePool),
	// 	},
	// }
}
