package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-09-01/StorageAccountCreateWithTurboTier.json
func ExampleAccountsClient_BeginCreate_storageAccountCreateWithTurboTier() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "res9101", "staturbocontoso01", armstorage.AccountCreateParameters{
		Kind:     to.Ptr(armstorage.KindStorageV2),
		Location: to.Ptr("eastus2"),
		Properties: &armstorage.AccountPropertiesCreateParameters{
			AccessTier: to.Ptr(armstorage.AccessTierHot),
			TurboTier: &armstorage.TurboTier{
				Status:        to.Ptr(armstorage.TurboTierStatusEnabled),
				TargetPercent: to.Ptr[int32](10),
			},
		},
		SKU: &armstorage.SKU{
			Name: to.Ptr(armstorage.SKUNameStandardLRS),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.AccountsClientCreateResponse{
	// 	Account: armstorage.Account{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Storage/storageAccounts/staturbocontoso01"),
	// 		Kind: to.Ptr(armstorage.KindStorageV2),
	// 		Location: to.Ptr("eastus2"),
	// 		Name: to.Ptr("staturbocontoso01"),
	// 		Properties: &armstorage.AccountProperties{
	// 			AccessTier: to.Ptr(armstorage.AccessTierHot),
	// 			ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 			TurboTier: &armstorage.TurboTier{
	// 				Status: to.Ptr(armstorage.TurboTierStatusEnabled),
	// 				TargetPercent: to.Ptr[int32](10),
	// 			},
	// 		},
	// 		SKU: &armstorage.SKU{
	// 			Name: to.Ptr(armstorage.SKUNameStandardLRS),
	// 			Tier: to.Ptr(armstorage.SKUTierStandard),
	// 		},
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 	},
	// }
}
