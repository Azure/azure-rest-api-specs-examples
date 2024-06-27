package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountUpdateWithImmutabilityPolicy.json
func ExampleAccountsClient_Update_storageAccountUpdateWithImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Update(ctx, "res9407", "sto8596", armstorage.AccountUpdateParameters{
		Properties: &armstorage.AccountPropertiesUpdateParameters{
			ImmutableStorageWithVersioning: &armstorage.ImmutableStorageAccount{
				Enabled: to.Ptr(true),
				ImmutabilityPolicy: &armstorage.AccountImmutabilityPolicyProperties{
					AllowProtectedAppendWrites:            to.Ptr(true),
					ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](15),
					State:                                 to.Ptr(armstorage.AccountImmutabilityPolicyStateLocked),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armstorage.Account{
	// 	Name: to.Ptr("sto8596"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9407/providers/Microsoft.Storage/storageAccounts/sto8596"),
	// 	Location: to.Ptr("eastus2(stage)"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Kind: to.Ptr(armstorage.KindStorage),
	// 	Properties: &armstorage.AccountProperties{
	// 		ImmutableStorageWithVersioning: &armstorage.ImmutableStorageAccount{
	// 			Enabled: to.Ptr(true),
	// 			ImmutabilityPolicy: &armstorage.AccountImmutabilityPolicyProperties{
	// 				AllowProtectedAppendWrites: to.Ptr(true),
	// 				ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](15),
	// 				State: to.Ptr(armstorage.AccountImmutabilityPolicyStateLocked),
	// 			},
	// 		},
	// 	},
	// 	SKU: &armstorage.SKU{
	// 		Name: to.Ptr(armstorage.SKUNameStandardGRS),
	// 		Tier: to.Ptr(armstorage.SKUTierStandard),
	// 	},
	// }
}
