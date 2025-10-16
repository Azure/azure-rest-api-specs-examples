package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2025-07-01/ResourceGuardCRUD/PatchResourceGuard.json
func ExampleResourceGuardsClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("0b352192-dcac-4cc7-992e-a96190ccc68c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceGuardsClient().Patch(ctx, "SampleResourceGroup", "swaggerExample", armdataprotection.PatchResourceGuardInput{
		Tags: map[string]*string{
			"newKey": to.Ptr("newVal"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataprotection.ResourceGuardsClientPatchResponse{
	// 	ResourceGuardResource: &armdataprotection.ResourceGuardResource{
	// 		Name: to.Ptr("VaultGuardTestNew"),
	// 		Type: to.Ptr("Microsoft.DataProtection/resourceGuards"),
	// 		ID: to.Ptr("/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armdataprotection.ResourceGuard{
	// 			Description: to.Ptr("Please take JIT access before performing any of the critical operation"),
	// 			AllowAutoApprovals: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 			ResourceGuardOperations: []*armdataprotection.ResourceGuardOperation{
	// 				{
	// 					RequestResourceType: to.Ptr("Microsoft.DataProtection/resourceGuards/deleteResourceGuardProxyRequests"),
	// 					VaultCriticalOperation: to.Ptr("Microsoft.RecoveryServices/vaults/backupResourceGuardProxies/delete"),
	// 				},
	// 				{
	// 					RequestResourceType: to.Ptr("Microsoft.DataProtection/resourceGuards/disableSoftDeleteRequests"),
	// 					VaultCriticalOperation: to.Ptr("Microsoft.RecoveryServices/vaults/backupconfig/write"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"TestKey": to.Ptr("TestValue"),
	// 		},
	// 	},
	// }
}
