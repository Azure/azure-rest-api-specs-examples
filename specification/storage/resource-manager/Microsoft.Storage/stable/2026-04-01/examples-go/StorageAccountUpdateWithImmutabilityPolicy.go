package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-04-01/StorageAccountUpdateWithImmutabilityPolicy.json
func ExampleAccountsClient_Update_storageAccountUpdateWithImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
	// res = armstorage.AccountsClientUpdateResponse{
	// 	Account: armstorage.Account{
	// 		Name: to.Ptr("sto8596"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9407/providers/Microsoft.Storage/storageAccounts/sto8596"),
	// 		Kind: to.Ptr(armstorage.KindStorage),
	// 		Location: to.Ptr("eastus2(stage)"),
	// 		Properties: &armstorage.AccountProperties{
	// 			ImmutableStorageWithVersioning: &armstorage.ImmutableStorageAccount{
	// 				Enabled: to.Ptr(true),
	// 				ImmutabilityPolicy: &armstorage.AccountImmutabilityPolicyProperties{
	// 					AllowProtectedAppendWrites: to.Ptr(true),
	// 					ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](15),
	// 					State: to.Ptr(armstorage.AccountImmutabilityPolicyStateLocked),
	// 				},
	// 			},
	// 		},
	// 		SKU: &armstorage.SKU{
	// 			Name: to.Ptr(armstorage.SKUNameStandardGRS),
	// 			Tier: to.Ptr(armstorage.SKUTierStandard),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
