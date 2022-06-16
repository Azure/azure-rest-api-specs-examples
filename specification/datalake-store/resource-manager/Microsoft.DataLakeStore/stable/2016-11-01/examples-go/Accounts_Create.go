package armdatalakestore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Create.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakestore.NewAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"contosorg",
		"contosoadla",
		armdatalakestore.CreateDataLakeStoreAccountParameters{
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
		},
		nil)
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
