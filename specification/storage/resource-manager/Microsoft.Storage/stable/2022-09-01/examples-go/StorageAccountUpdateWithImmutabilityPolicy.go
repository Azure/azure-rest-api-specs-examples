package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountUpdateWithImmutabilityPolicy.json
func ExampleAccountsClient_Update_storageAccountUpdateWithImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewAccountsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "res9407", "sto8596", armstorage.AccountUpdateParameters{
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
	// TODO: use response item
	_ = res
}
