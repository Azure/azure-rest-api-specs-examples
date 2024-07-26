package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/ResourceGuardCRUD/PutResourceGuard.json
func ExampleResourceGuardsClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceGuardsClient().Put(ctx, "SampleResourceGroup", "swaggerExample", armdataprotection.ResourceGuardResource{
		Location: to.Ptr("WestUS"),
		Tags: map[string]*string{
			"key1": to.Ptr("val1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceGuardResource = armdataprotection.ResourceGuardResource{
	// 	Name: to.Ptr("VaultGuardTestNew"),
	// 	Type: to.Ptr("Microsoft.DataProtection/resourceGuards"),
	// 	ID: to.Ptr("/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"TestKey": to.Ptr("TestValue"),
	// 	},
	// 	Properties: &armdataprotection.ResourceGuard{
	// 		Description: to.Ptr("Please take JIT access before performing any of the critical operation"),
	// 		AllowAutoApprovals: to.Ptr(true),
	// 		ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 		ResourceGuardOperations: []*armdataprotection.ResourceGuardOperation{
	// 			{
	// 				RequestResourceType: to.Ptr("Microsoft.DataProtection/resourceGuards/deleteResourceGuardProxyRequests"),
	// 				VaultCriticalOperation: to.Ptr("Microsoft.RecoveryServices/vaults/backupResourceGuardProxies/delete"),
	// 			},
	// 			{
	// 				RequestResourceType: to.Ptr("Microsoft.DataProtection/resourceGuards/disableSoftDeleteRequests"),
	// 				VaultCriticalOperation: to.Ptr("Microsoft.RecoveryServices/vaults/backupconfig/write"),
	// 		}},
	// 		VaultCriticalOperationExclusionList: []*string{
	// 		},
	// 	},
	// }
}
