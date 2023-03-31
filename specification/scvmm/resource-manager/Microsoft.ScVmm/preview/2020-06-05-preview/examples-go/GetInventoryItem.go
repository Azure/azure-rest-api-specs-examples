package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/GetInventoryItem.json
func ExampleInventoryItemsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInventoryItemsClient().Get(ctx, "testrg", "ContosoVMMServer", "12345678-1234-1234-1234-123456789abc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InventoryItem = armscvmm.InventoryItem{
	// 	Name: to.Ptr("12345678-1234-1234-1234-123456789abc"),
	// 	Type: to.Ptr("Microsoft.SCVMM/VMMServers/InventoryItems"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer/InventoryItems/12345678-1234-1234-1234-123456789abc"),
	// 	Properties: &armscvmm.CloudInventoryItem{
	// 		InventoryItemName: to.Ptr("contoso-cloud"),
	// 		InventoryType: to.Ptr(armscvmm.InventoryTypeCloud),
	// 		ManagedResourceID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/Clouds/contoso-cloud"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		UUID: to.Ptr("12345678-1234-1234-1234-123456789abc"),
	// 	},
	// }
}
