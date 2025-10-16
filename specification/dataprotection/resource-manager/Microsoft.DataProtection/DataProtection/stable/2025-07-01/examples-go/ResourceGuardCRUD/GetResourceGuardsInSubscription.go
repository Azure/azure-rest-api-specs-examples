package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2025-07-01/ResourceGuardCRUD/GetResourceGuardsInSubscription.json
func ExampleResourceGuardsClient_NewGetResourcesInSubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("0b352192-dcac-4cc7-992e-a96190ccc68c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceGuardsClient().NewGetResourcesInSubscriptionPager(nil)
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
		// page = armdataprotection.ResourceGuardsClientGetResourcesInSubscriptionResponse{
		// 	ResourceGuardResourceList: armdataprotection.ResourceGuardResourceList{
		// 		Value: []*armdataprotection.ResourceGuardResource{
		// 			{
		// 				Name: to.Ptr("VaultGuardTestNew"),
		// 				Type: to.Ptr("Microsoft.DataProtection/resourceGuards"),
		// 				ID: to.Ptr("/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armdataprotection.ResourceGuard{
		// 					Description: to.Ptr("Please take JIT access before performing any of the critical operation"),
		// 					AllowAutoApprovals: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
		// 					ResourceGuardOperations: []*armdataprotection.ResourceGuardOperation{
		// 						{
		// 							RequestResourceType: to.Ptr("Microsoft.DataProtection/resourceGuards/deleteResourceGuardProxyRequests"),
		// 							VaultCriticalOperation: to.Ptr("Microsoft.RecoveryServices/vaults/backupResourceGuardProxies/delete"),
		// 						},
		// 						{
		// 							RequestResourceType: to.Ptr("Microsoft.DataProtection/resourceGuards/disableSoftDeleteRequests"),
		// 							VaultCriticalOperation: to.Ptr("Microsoft.RecoveryServices/vaults/backupconfig/write"),
		// 						},
		// 					},
		// 					VaultCriticalOperationExclusionList: []*string{
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"TestKey": to.Ptr("TestValue"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
