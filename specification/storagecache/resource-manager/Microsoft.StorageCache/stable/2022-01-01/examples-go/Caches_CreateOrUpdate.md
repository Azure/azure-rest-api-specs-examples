```go
package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/Caches_CreateOrUpdate.json
func ExampleCachesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragecache.NewCachesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"scgroup",
		"sc1",
		&armstoragecache.CachesClientBeginCreateOrUpdateOptions{Cache: &armstoragecache.Cache{
			Identity: &armstoragecache.CacheIdentity{
				Type: to.Ptr(armstoragecache.CacheIdentityTypeUserAssigned),
				UserAssignedIdentities: map[string]*armstoragecache.UserAssignedIdentitiesValue{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
				},
			},
			Location: to.Ptr("westus"),
			Properties: &armstoragecache.CacheProperties{
				CacheSizeGB: to.Ptr[int32](3072),
				DirectoryServicesSettings: &armstoragecache.CacheDirectorySettings{
					ActiveDirectory: &armstoragecache.CacheActiveDirectorySettings{
						CacheNetBiosName: to.Ptr("contosoSmb"),
						Credentials: &armstoragecache.CacheActiveDirectorySettingsCredentials{
							Password: to.Ptr("<password>"),
							Username: to.Ptr("consotoAdmin"),
						},
						DomainName:            to.Ptr("contosoAd.contoso.local"),
						DomainNetBiosName:     to.Ptr("contosoAd"),
						PrimaryDNSIPAddress:   to.Ptr("192.0.2.10"),
						SecondaryDNSIPAddress: to.Ptr("192.0.2.11"),
					},
					UsernameDownload: &armstoragecache.CacheUsernameDownloadSettings{
						Credentials: &armstoragecache.CacheUsernameDownloadSettingsCredentials{
							BindDn:       to.Ptr("cn=ldapadmin,dc=contosoad,dc=contoso,dc=local"),
							BindPassword: to.Ptr("<bindPassword>"),
						},
						ExtendedGroups: to.Ptr(true),
						LdapBaseDN:     to.Ptr("dc=contosoad,dc=contoso,dc=local"),
						LdapServer:     to.Ptr("192.0.2.12"),
						UsernameSource: to.Ptr(armstoragecache.UsernameSourceLDAP),
					},
				},
				EncryptionSettings: &armstoragecache.CacheEncryptionSettings{
					KeyEncryptionKey: &armstoragecache.KeyVaultKeyReference{
						KeyURL: to.Ptr("https://keyvault-cmk.vault.azure.net/keys/key2047/test"),
						SourceVault: &armstoragecache.KeyVaultKeyReferenceSourceVault{
							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"),
						},
					},
				},
				SecuritySettings: &armstoragecache.CacheSecuritySettings{
					AccessPolicies: []*armstoragecache.NfsAccessPolicy{
						{
							Name: to.Ptr("default"),
							AccessRules: []*armstoragecache.NfsAccessRule{
								{
									Access:         to.Ptr(armstoragecache.NfsAccessRuleAccessRw),
									RootSquash:     to.Ptr(false),
									Scope:          to.Ptr(armstoragecache.NfsAccessRuleScopeDefault),
									SubmountAccess: to.Ptr(true),
									Suid:           to.Ptr(false),
								}},
						}},
				},
				Subnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/sub1"),
			},
			SKU: &armstoragecache.CacheSKU{
				Name: to.Ptr("Standard_2G"),
			},
			Tags: map[string]*string{
				"Dept": to.Ptr("Contoso"),
			},
		},
		})
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragecache%2Farmstoragecache%2Fv1.0.0/sdk/resourcemanager/storagecache/armstoragecache/README.md) on how to add the SDK to your project and authenticate.
