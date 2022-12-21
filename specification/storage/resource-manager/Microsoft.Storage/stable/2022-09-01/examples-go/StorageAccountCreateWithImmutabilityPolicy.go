package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountCreateWithImmutabilityPolicy.json
func ExampleAccountsClient_BeginCreate_storageAccountCreateWithImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewAccountsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "res9101", "sto4445", armstorage.AccountCreateParameters{
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
	// TODO: use response item
	_ = res
}
