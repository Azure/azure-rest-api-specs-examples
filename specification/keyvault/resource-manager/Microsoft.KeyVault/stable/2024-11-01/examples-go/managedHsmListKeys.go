package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/managedHsmListKeys.json
func ExampleManagedHsmKeysClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedHsmKeysClient().NewListPager("sample-group", "sample-managedhsm-name", nil)
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
		// page.ManagedHsmKeyListResult = armkeyvault.ManagedHsmKeyListResult{
		// 	Value: []*armkeyvault.ManagedHsmKey{
		// 		{
		// 			Name: to.Ptr("sample-key-name-1"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedHSMs/keys"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedHSMs/sample-managedhsm-name/keys/sample-key-name-1"),
		// 			Properties: &armkeyvault.ManagedHsmKeyProperties{
		// 				Attributes: &armkeyvault.ManagedHsmKeyAttributes{
		// 					Created: to.Ptr[int64](1596493796),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1596493796),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name-1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sample-key-name-2"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedHSMs/keys"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedHSMs/sample-managedhsm-name/keys/sample-key-name-2"),
		// 			Properties: &armkeyvault.ManagedHsmKeyProperties{
		// 				Attributes: &armkeyvault.ManagedHsmKeyAttributes{
		// 					Created: to.Ptr[int64](1596493797),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1596493797),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-managedhsm-name.managedhsm.azure.net:443/keys/sample-key-name-2"),
		// 			},
		// 	}},
		// }
	}
}
