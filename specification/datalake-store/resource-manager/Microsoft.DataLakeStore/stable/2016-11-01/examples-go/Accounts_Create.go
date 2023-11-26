package armdatalakestore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Create.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakestore.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "contosorg", "contosoadla", armdatalakestore.CreateDataLakeStoreAccountParameters{
		Identity: &armdatalakestore.EncryptionIdentity{
			Type: to.Ptr("SystemAssigned"),
		},
		Location: to.Ptr("eastus2"),
		Properties: &armdatalakestore.CreateDataLakeStoreAccountProperties{
			DefaultGroup: to.Ptr("test_default_group"),
			EncryptionConfig: &armdatalakestore.EncryptionConfig{
				Type: to.Ptr(armdatalakestore.EncryptionConfigTypeUserManaged),
				KeyVaultMetaInfo: &armdatalakestore.KeyVaultMetaInfo{
					EncryptionKeyName:    to.Ptr("test_encryption_key_name"),
					EncryptionKeyVersion: to.Ptr("encryption_key_version"),
					KeyVaultResourceID:   to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
				},
			},
			EncryptionState:       to.Ptr(armdatalakestore.EncryptionStateEnabled),
			FirewallAllowAzureIPs: to.Ptr(armdatalakestore.FirewallAllowAzureIPsStateEnabled),
			FirewallRules: []*armdatalakestore.CreateFirewallRuleWithAccountParameters{
				{
					Name: to.Ptr("test_rule"),
					Properties: &armdatalakestore.CreateOrUpdateFirewallRuleProperties{
						EndIPAddress:   to.Ptr("2.2.2.2"),
						StartIPAddress: to.Ptr("1.1.1.1"),
					},
				}},
			FirewallState:          to.Ptr(armdatalakestore.FirewallStateEnabled),
			NewTier:                to.Ptr(armdatalakestore.TierTypeConsumption),
			TrustedIDProviderState: to.Ptr(armdatalakestore.TrustedIDProviderStateEnabled),
			TrustedIDProviders: []*armdatalakestore.CreateTrustedIDProviderWithAccountParameters{
				{
					Name: to.Ptr("test_trusted_id_provider_name"),
					Properties: &armdatalakestore.CreateOrUpdateTrustedIDProviderProperties{
						IDProvider: to.Ptr("https://sts.windows.net/ea9ec534-a3e3-4e45-ad36-3afc5bb291c1"),
					},
				}},
		},
		Tags: map[string]*string{
			"test_key": to.Ptr("test_value"),
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
	// res.Account = armdatalakestore.Account{
	// 	Name: to.Ptr("contosoadla"),
	// 	Type: to.Ptr("test_type"),
	// 	ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 	Location: to.Ptr("eastus2"),
	// 	Tags: map[string]*string{
	// 		"test_key": to.Ptr("test_value"),
	// 	},
	// 	Identity: &armdatalakestore.EncryptionIdentity{
	// 		Type: to.Ptr("SystemAssigned"),
	// 		PrincipalID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 		TenantID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab346"),
	// 	},
	// 	Properties: &armdatalakestore.AccountProperties{
	// 		AccountID: to.Ptr("94f4bf5d-78a9-4c31-8aa7-b34d07bad898"),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-14T20:21:56.681Z"); return t}()),
	// 		Endpoint: to.Ptr("testadlfs17607.azuredatalakestore.net"),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-14T20:21:56.681Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armdatalakestore.DataLakeStoreAccountStatusSucceeded),
	// 		State: to.Ptr(armdatalakestore.DataLakeStoreAccountStateActive),
	// 		CurrentTier: to.Ptr(armdatalakestore.TierTypeConsumption),
	// 		DefaultGroup: to.Ptr("test_default_group"),
	// 		EncryptionConfig: &armdatalakestore.EncryptionConfig{
	// 			Type: to.Ptr(armdatalakestore.EncryptionConfigTypeUserManaged),
	// 			KeyVaultMetaInfo: &armdatalakestore.KeyVaultMetaInfo{
	// 				EncryptionKeyName: to.Ptr("test_encryption_key_name"),
	// 				EncryptionKeyVersion: to.Ptr("encryption_key_version"),
	// 				KeyVaultResourceID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 			},
	// 		},
	// 		EncryptionProvisioningState: to.Ptr(armdatalakestore.EncryptionProvisioningStateSucceeded),
	// 		EncryptionState: to.Ptr(armdatalakestore.EncryptionStateEnabled),
	// 		FirewallAllowAzureIPs: to.Ptr(armdatalakestore.FirewallAllowAzureIPsStateEnabled),
	// 		FirewallRules: []*armdatalakestore.FirewallRule{
	// 			{
	// 				Name: to.Ptr("test_rule"),
	// 				Type: to.Ptr("test_type"),
	// 				ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 				Properties: &armdatalakestore.FirewallRuleProperties{
	// 					EndIPAddress: to.Ptr("2.2.2.2"),
	// 					StartIPAddress: to.Ptr("1.1.1.1"),
	// 				},
	// 		}},
	// 		FirewallState: to.Ptr(armdatalakestore.FirewallStateEnabled),
	// 		NewTier: to.Ptr(armdatalakestore.TierTypeConsumption),
	// 		TrustedIDProviderState: to.Ptr(armdatalakestore.TrustedIDProviderStateEnabled),
	// 		TrustedIDProviders: []*armdatalakestore.TrustedIDProvider{
	// 			{
	// 				Name: to.Ptr("test_trusted_id_provider_name"),
	// 				Type: to.Ptr("test_type"),
	// 				ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 				Properties: &armdatalakestore.TrustedIDProviderProperties{
	// 					IDProvider: to.Ptr("https://sts.windows.net/ea9ec534-a3e3-4e45-ad36-3afc5bb291c1"),
	// 				},
	// 		}},
	// 	},
	// }
}
