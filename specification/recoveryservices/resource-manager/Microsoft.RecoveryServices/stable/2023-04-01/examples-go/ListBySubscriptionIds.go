package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/ListBySubscriptionIds.json
func ExampleVaultsClient_NewListBySubscriptionIDPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVaultsClient().NewListBySubscriptionIDPager(nil)
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
		// page.VaultList = armrecoveryservices.VaultList{
		// 	Value: []*armrecoveryservices.Vault{
		// 		{
		// 			Name: to.Ptr("patchtest"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults"),
		// 			Etag: to.Ptr("W/\"datetime'2017-11-22T11%3A05%3A19.907Z'\""),
		// 			ID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/resourceGroups/Default-RecoveryServices-ResourceGroup/providers/Microsoft.RecoveryServices/vaults/patchtest"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"Love": to.Ptr("India"),
		// 			},
		// 			Properties: &armrecoveryservices.VaultProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 			SKU: &armrecoveryservices.SKU{
		// 				Name: to.Ptr(armrecoveryservices.SKUNameStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("swaggerExample"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults"),
		// 			Etag: to.Ptr("W/\"datetime'2017-12-15T12%3A36%3A51.68Z'\""),
		// 			ID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/resourceGroups/HelloWorld/providers/Microsoft.RecoveryServices/vaults/swaggerExample"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"TestUpdatedKey": to.Ptr("TestUpdatedValue"),
		// 			},
		// 			Properties: &armrecoveryservices.VaultProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 			SKU: &armrecoveryservices.SKU{
		// 				Name: to.Ptr(armrecoveryservices.SKUNameStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("today1"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults"),
		// 			Etag: to.Ptr("W/\"datetime'2017-11-21T10%3A52%3A19.633Z'\""),
		// 			ID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/resourceGroups/HelloWorld/providers/Microsoft.RecoveryServices/vaults/today1"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"TestUpdatedKey": to.Ptr("TestUpdatedValue"),
		// 			},
		// 			Properties: &armrecoveryservices.VaultProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 			SKU: &armrecoveryservices.SKU{
		// 				Name: to.Ptr(armrecoveryservices.SKUNameStandard),
		// 			},
		// 	}},
		// }
	}
}
