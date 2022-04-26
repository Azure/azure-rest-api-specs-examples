Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragecache%2Farmstoragecache%2Fv0.4.0/sdk/resourcemanager/storagecache/armstoragecache/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragecache_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/Caches_CreateOrUpdate.json
func ExampleCachesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstoragecache.NewCachesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cache-name>",
		&armstoragecache.CachesClientBeginCreateOrUpdateOptions{Cache: &armstoragecache.Cache{
			Identity: &armstoragecache.CacheIdentity{
				Type: to.Ptr(armstoragecache.CacheIdentityTypeUserAssigned),
				UserAssignedIdentities: map[string]*armstoragecache.UserAssignedIdentitiesValue{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
				},
			},
			Location: to.Ptr("<location>"),
			Properties: &armstoragecache.CacheProperties{
				CacheSizeGB: to.Ptr[int32](3072),
				DirectoryServicesSettings: &armstoragecache.CacheDirectorySettings{
					ActiveDirectory: &armstoragecache.CacheActiveDirectorySettings{
						CacheNetBiosName: to.Ptr("<cache-net-bios-name>"),
						Credentials: &armstoragecache.CacheActiveDirectorySettingsCredentials{
							Password: to.Ptr("<password>"),
							Username: to.Ptr("<username>"),
						},
						DomainName:            to.Ptr("<domain-name>"),
						DomainNetBiosName:     to.Ptr("<domain-net-bios-name>"),
						PrimaryDNSIPAddress:   to.Ptr("<primary-dnsipaddress>"),
						SecondaryDNSIPAddress: to.Ptr("<secondary-dnsipaddress>"),
					},
					UsernameDownload: &armstoragecache.CacheUsernameDownloadSettings{
						Credentials: &armstoragecache.CacheUsernameDownloadSettingsCredentials{
							BindDn:       to.Ptr("<bind-dn>"),
							BindPassword: to.Ptr("<bind-password>"),
						},
						ExtendedGroups: to.Ptr(true),
						LdapBaseDN:     to.Ptr("<ldap-base-dn>"),
						LdapServer:     to.Ptr("<ldap-server>"),
						UsernameSource: to.Ptr(armstoragecache.UsernameSourceLDAP),
					},
				},
				EncryptionSettings: &armstoragecache.CacheEncryptionSettings{
					KeyEncryptionKey: &armstoragecache.KeyVaultKeyReference{
						KeyURL: to.Ptr("<key-url>"),
						SourceVault: &armstoragecache.KeyVaultKeyReferenceSourceVault{
							ID: to.Ptr("<id>"),
						},
					},
				},
				SecuritySettings: &armstoragecache.CacheSecuritySettings{
					AccessPolicies: []*armstoragecache.NfsAccessPolicy{
						{
							Name: to.Ptr("<name>"),
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
				Subnet: to.Ptr("<subnet>"),
			},
			SKU: &armstoragecache.CacheSKU{
				Name: to.Ptr("<name>"),
			},
			Tags: map[string]*string{
				"Dept": to.Ptr("Contoso"),
			},
		},
			ResumeToken: "",
		})
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
