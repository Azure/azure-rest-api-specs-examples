Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatalake-store%2Farmdatalakestore%2Fv0.5.0/sdk/resourcemanager/datalake-store/armdatalakestore/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatalakestore_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Create.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatalakestore.NewAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armdatalakestore.CreateDataLakeStoreAccountParameters{
			Identity: &armdatalakestore.EncryptionIdentity{
				Type: to.Ptr("<type>"),
			},
			Location: to.Ptr("<location>"),
			Properties: &armdatalakestore.CreateDataLakeStoreAccountProperties{
				DefaultGroup: to.Ptr("<default-group>"),
				EncryptionConfig: &armdatalakestore.EncryptionConfig{
					Type: to.Ptr(armdatalakestore.EncryptionConfigTypeUserManaged),
					KeyVaultMetaInfo: &armdatalakestore.KeyVaultMetaInfo{
						EncryptionKeyName:    to.Ptr("<encryption-key-name>"),
						EncryptionKeyVersion: to.Ptr("<encryption-key-version>"),
						KeyVaultResourceID:   to.Ptr("<key-vault-resource-id>"),
					},
				},
				EncryptionState:       to.Ptr(armdatalakestore.EncryptionStateEnabled),
				FirewallAllowAzureIPs: to.Ptr(armdatalakestore.FirewallAllowAzureIPsStateEnabled),
				FirewallRules: []*armdatalakestore.CreateFirewallRuleWithAccountParameters{
					{
						Name: to.Ptr("<name>"),
						Properties: &armdatalakestore.CreateOrUpdateFirewallRuleProperties{
							EndIPAddress:   to.Ptr("<end-ipaddress>"),
							StartIPAddress: to.Ptr("<start-ipaddress>"),
						},
					}},
				FirewallState:          to.Ptr(armdatalakestore.FirewallStateEnabled),
				NewTier:                to.Ptr(armdatalakestore.TierTypeConsumption),
				TrustedIDProviderState: to.Ptr(armdatalakestore.TrustedIDProviderStateEnabled),
				TrustedIDProviders: []*armdatalakestore.CreateTrustedIDProviderWithAccountParameters{
					{
						Name: to.Ptr("<name>"),
						Properties: &armdatalakestore.CreateOrUpdateTrustedIDProviderProperties{
							IDProvider: to.Ptr("<idprovider>"),
						},
					}},
			},
			Tags: map[string]*string{
				"test_key": to.Ptr("test_value"),
			},
		},
		&armdatalakestore.AccountsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
