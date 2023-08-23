package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountCreateWithImmutabilityPolicy.json
func ExampleAccountsClient_BeginCreate_storageAccountCreateWithImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "res9101", "sto4445", armstorage.AccountCreateParameters{
		ExtendedLocation: &armstorage.ExtendedLocation{
			Name: to.Ptr("losangeles001"),
			Type: to.Ptr(armstorage.ExtendedLocationTypesEdgeZone),
		},
		Kind:     to.Ptr(armstorage.KindStorage),
		Location: to.Ptr("eastus"),
		Properties: &armstorage.AccountPropertiesCreateParameters{
			ImmutableStorageWithVersioning: &armstorage.ImmutableStorageAccount{
				Enabled: to.Ptr(true),
				ImmutabilityPolicy: &armstorage.AccountImmutabilityPolicyProperties{
					AllowProtectedAppendWrites:            to.Ptr(true),
					ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](15),
					State:                                 to.Ptr(armstorage.AccountImmutabilityPolicyStateUnlocked),
				},
			},
		},
		SKU: &armstorage.SKU{
			Name: to.Ptr(armstorage.SKUNameStandardGRS),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armstorage.Account{
	// 	Name: to.Ptr("sto4445"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Storage/storageAccounts/sto4445"),
	// 	Location: to.Ptr("eastus"),
	// 	ExtendedLocation: &armstorage.ExtendedLocation{
	// 		Name: to.Ptr("losangeles001"),
	// 		Type: to.Ptr(armstorage.ExtendedLocationTypesEdgeZone),
	// 	},
	// 	Kind: to.Ptr(armstorage.KindStorage),
	// 	Properties: &armstorage.AccountProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T13:25:33.4863236Z"); return t}()),
	// 		ImmutableStorageWithVersioning: &armstorage.ImmutableStorageAccount{
	// 			Enabled: to.Ptr(true),
	// 			ImmutabilityPolicy: &armstorage.AccountImmutabilityPolicyProperties{
	// 				AllowProtectedAppendWrites: to.Ptr(true),
	// 				ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](15),
	// 				State: to.Ptr(armstorage.AccountImmutabilityPolicyStateUnlocked),
	// 			},
	// 		},
	// 		PrimaryEndpoints: &armstorage.Endpoints{
	// 			Blob: to.Ptr("https://sto4445.blob.core.windows.net/"),
	// 			Dfs: to.Ptr("https://sto4445.dfs.core.windows.net/"),
	// 			File: to.Ptr("https://sto4445.file.core.windows.net/"),
	// 			Queue: to.Ptr("https://sto4445.queue.core.windows.net/"),
	// 			Table: to.Ptr("https://sto4445.table.core.windows.net/"),
	// 			Web: to.Ptr("https://sto4445.web.core.windows.net/"),
	// 		},
	// 		PrimaryLocation: to.Ptr("eastus2euap"),
	// 		ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armstorage.SKU{
	// 		Name: to.Ptr(armstorage.SKUNameStandardGRS),
	// 		Tier: to.Ptr(armstorage.SKUTierStandard),
	// 	},
	// }
}
