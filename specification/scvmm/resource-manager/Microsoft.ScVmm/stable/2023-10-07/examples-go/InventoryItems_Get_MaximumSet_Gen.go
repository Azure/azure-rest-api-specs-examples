package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/InventoryItems_Get_MaximumSet_Gen.json
func ExampleInventoryItemsClient_Get_inventoryItemsGetMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInventoryItemsClient().Get(ctx, "rgscvmm", "1", "2bFBede6-EEf8-becB-dBbd-B96DbBFdB3f3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InventoryItem = armscvmm.InventoryItem{
	// 	Name: to.Ptr("oimmcgxagnhmasgsmhdaigznub"),
	// 	Type: to.Ptr("lfhuayaplzxdqzubmjvtgcan"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/{vmmServerName}/inventoryItems/inventoryItemResourceName"),
	// 	SystemData: &armscvmm.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.094Z"); return t}()),
	// 		CreatedBy: to.Ptr("p"),
	// 		CreatedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.095Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("goxcwpyyqlxndquly"),
	// 		LastModifiedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 	},
	// 	Kind: to.Ptr("M\\d_,V."),
	// 	Properties: &armscvmm.InventoryItemProperties{
	// 		InventoryItemName: to.Ptr("kspgdhmlmycalwrepfmshoaoumna"),
	// 		InventoryType: to.Ptr(armscvmm.InventoryType("InventoryItemProperties")),
	// 		ManagedResourceID: to.Ptr("ictxvjzvurnkdgwabqyyfyckkkdx"),
	// 		ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 		UUID: to.Ptr("jolmoxfopwfoje"),
	// 	},
	// }
}
