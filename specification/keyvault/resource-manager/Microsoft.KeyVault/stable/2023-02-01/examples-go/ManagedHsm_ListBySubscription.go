package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/ManagedHsm_ListBySubscription.json
func ExampleManagedHsmsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedHsmsClient().NewListBySubscriptionPager(&armkeyvault.ManagedHsmsClientListBySubscriptionOptions{Top: nil})
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
		// page.ManagedHsmListResult = armkeyvault.ManagedHsmListResult{
		// 	Value: []*armkeyvault.ManagedHsm{
		// 		{
		// 			Name: to.Ptr("hsm1"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedHSMs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm1"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armkeyvault.ManagedHsmSKU{
		// 				Name: to.Ptr(armkeyvault.ManagedHsmSKUNameStandardB1),
		// 				Family: to.Ptr(armkeyvault.ManagedHsmSKUFamilyB),
		// 			},
		// 			Tags: map[string]*string{
		// 				"Dept": to.Ptr("hsm"),
		// 				"Environment": to.Ptr("dogfood"),
		// 			},
		// 			Properties: &armkeyvault.ManagedHsmProperties{
		// 				EnablePurgeProtection: to.Ptr(false),
		// 				EnableSoftDelete: to.Ptr(true),
		// 				HsmURI: to.Ptr("https://westus.hsm1.managedhsm.azure.net"),
		// 				InitialAdminObjectIDs: []*string{
		// 					to.Ptr("00000000-0000-0000-0000-000000000000")},
		// 					ProvisioningState: to.Ptr(armkeyvault.ProvisioningStateSucceeded),
		// 					SoftDeleteRetentionInDays: to.Ptr[int32](90),
		// 					StatusMessage: to.Ptr("ManagedHsm is functional."),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("hsm2"),
		// 				Type: to.Ptr("Microsoft.KeyVault/managedHSMs"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.KeyVault/managedHSMs/hsm2"),
		// 				Location: to.Ptr("westus"),
		// 				SKU: &armkeyvault.ManagedHsmSKU{
		// 					Name: to.Ptr(armkeyvault.ManagedHsmSKUNameStandardB1),
		// 					Family: to.Ptr(armkeyvault.ManagedHsmSKUFamilyB),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("hsm"),
		// 					"Environment": to.Ptr("production"),
		// 				},
		// 				Properties: &armkeyvault.ManagedHsmProperties{
		// 					EnablePurgeProtection: to.Ptr(false),
		// 					EnableSoftDelete: to.Ptr(true),
		// 					HsmURI: to.Ptr("https://westus.hsm2.managedhsm.azure.net"),
		// 					InitialAdminObjectIDs: []*string{
		// 						to.Ptr("00000000-0000-0000-0000-000000000000")},
		// 						ProvisioningState: to.Ptr(armkeyvault.ProvisioningStateSucceeded),
		// 						SoftDeleteRetentionInDays: to.Ptr[int32](90),
		// 						StatusMessage: to.Ptr("ManagedHsm is functional."),
		// 						TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 			}},
		// 		}
	}
}
