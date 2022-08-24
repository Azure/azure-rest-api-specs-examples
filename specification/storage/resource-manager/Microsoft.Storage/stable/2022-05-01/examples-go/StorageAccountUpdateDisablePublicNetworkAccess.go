package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountUpdateDisablePublicNetworkAccess.json
func ExampleAccountsClient_Update_storageAccountUpdateDisablePublicNetworkAccess() {
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
			AllowBlobPublicAccess: to.Ptr(false),
			AllowSharedKeyAccess:  to.Ptr(true),
			Encryption: &armstorage.Encryption{
				KeySource: to.Ptr(armstorage.KeySourceMicrosoftStorage),
				Services: &armstorage.EncryptionServices{
					Blob: &armstorage.EncryptionService{
						Enabled: to.Ptr(true),
						KeyType: to.Ptr(armstorage.KeyTypeAccount),
					},
					File: &armstorage.EncryptionService{
						Enabled: to.Ptr(true),
						KeyType: to.Ptr(armstorage.KeyTypeAccount),
					},
				},
			},
			KeyPolicy: &armstorage.KeyPolicy{
				KeyExpirationPeriodInDays: to.Ptr[int32](20),
			},
			MinimumTLSVersion: to.Ptr(armstorage.MinimumTLSVersionTLS12),
			NetworkRuleSet: &armstorage.NetworkRuleSet{
				DefaultAction: to.Ptr(armstorage.DefaultActionAllow),
				ResourceAccessRules: []*armstorage.ResourceAccessRule{
					{
						ResourceID: to.Ptr("/subscriptions/a7e99807-abbf-4642-bdec-2c809a96a8bc/resourceGroups/res9407/providers/Microsoft.Synapse/workspaces/testworkspace"),
						TenantID:   to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
					}},
			},
			PublicNetworkAccess: to.Ptr(armstorage.PublicNetworkAccessDisabled),
			RoutingPreference: &armstorage.RoutingPreference{
				PublishInternetEndpoints:  to.Ptr(true),
				PublishMicrosoftEndpoints: to.Ptr(true),
				RoutingChoice:             to.Ptr(armstorage.RoutingChoiceMicrosoftRouting),
			},
			SasPolicy: &armstorage.SasPolicy{
				ExpirationAction:    to.Ptr(armstorage.ExpirationActionLog),
				SasExpirationPeriod: to.Ptr("1.15:59:59"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
