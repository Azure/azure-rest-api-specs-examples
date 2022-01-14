Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragecache%2Farmstoragecache%2Fv0.2.0/sdk/resourcemanager/storagecache/armstoragecache/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/Caches_CreateOrUpdate.json
func ExampleCachesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstoragecache.NewCachesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cache-name>",
		&armstoragecache.CachesClientBeginCreateOrUpdateOptions{Cache: &armstoragecache.Cache{
			Identity: &armstoragecache.CacheIdentity{
				Type: armstoragecache.CacheIdentityTypeUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armstoragecache.UserAssignedIdentitiesValue{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
				},
			},
			Location: to.StringPtr("<location>"),
			Properties: &armstoragecache.CacheProperties{
				CacheSizeGB: to.Int32Ptr(3072),
				DirectoryServicesSettings: &armstoragecache.CacheDirectorySettings{
					ActiveDirectory: &armstoragecache.CacheActiveDirectorySettings{
						CacheNetBiosName: to.StringPtr("<cache-net-bios-name>"),
						Credentials: &armstoragecache.CacheActiveDirectorySettingsCredentials{
							Password: to.StringPtr("<password>"),
							Username: to.StringPtr("<username>"),
						},
						DomainName:            to.StringPtr("<domain-name>"),
						DomainNetBiosName:     to.StringPtr("<domain-net-bios-name>"),
						PrimaryDNSIPAddress:   to.StringPtr("<primary-dnsipaddress>"),
						SecondaryDNSIPAddress: to.StringPtr("<secondary-dnsipaddress>"),
					},
					UsernameDownload: &armstoragecache.CacheUsernameDownloadSettings{
						Credentials: &armstoragecache.CacheUsernameDownloadSettingsCredentials{
							BindDn:       to.StringPtr("<bind-dn>"),
							BindPassword: to.StringPtr("<bind-password>"),
						},
						ExtendedGroups: to.BoolPtr(true),
						LdapBaseDN:     to.StringPtr("<ldap-base-dn>"),
						LdapServer:     to.StringPtr("<ldap-server>"),
						UsernameSource: armstoragecache.UsernameSource("LDAP").ToPtr(),
					},
				},
				EncryptionSettings: &armstoragecache.CacheEncryptionSettings{
					KeyEncryptionKey: &armstoragecache.KeyVaultKeyReference{
						KeyURL: to.StringPtr("<key-url>"),
						SourceVault: &armstoragecache.KeyVaultKeyReferenceSourceVault{
							ID: to.StringPtr("<id>"),
						},
					},
				},
				SecuritySettings: &armstoragecache.CacheSecuritySettings{
					AccessPolicies: []*armstoragecache.NfsAccessPolicy{
						{
							Name: to.StringPtr("<name>"),
							AccessRules: []*armstoragecache.NfsAccessRule{
								{
									Access:         armstoragecache.NfsAccessRuleAccess("rw").ToPtr(),
									RootSquash:     to.BoolPtr(false),
									Scope:          armstoragecache.NfsAccessRuleScope("default").ToPtr(),
									SubmountAccess: to.BoolPtr(true),
									Suid:           to.BoolPtr(false),
								}},
						}},
				},
				Subnet: to.StringPtr("<subnet>"),
			},
			SKU: &armstoragecache.CacheSKU{
				Name: to.StringPtr("<name>"),
			},
			Tags: map[string]*string{
				"Dept": to.StringPtr("Contoso"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CachesClientCreateOrUpdateResult)
}
```
