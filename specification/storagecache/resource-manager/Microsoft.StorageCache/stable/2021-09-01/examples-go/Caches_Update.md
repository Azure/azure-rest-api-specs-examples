Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragecache%2Farmstoragecache%2Fv0.1.0/sdk/resourcemanager/storagecache/armstoragecache/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache"
)

// x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/Caches_Update.json
func ExampleCachesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstoragecache.NewCachesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<cache-name>",
		&armstoragecache.CachesUpdateOptions{Cache: &armstoragecache.Cache{
			Location: to.StringPtr("<location>"),
			Properties: &armstoragecache.CacheProperties{
				CacheSizeGB: to.Int32Ptr(3072),
				DirectoryServicesSettings: &armstoragecache.CacheDirectorySettings{
					ActiveDirectory: &armstoragecache.CacheActiveDirectorySettings{
						CacheNetBiosName:      to.StringPtr("<cache-net-bios-name>"),
						DomainName:            to.StringPtr("<domain-name>"),
						DomainNetBiosName:     to.StringPtr("<domain-net-bios-name>"),
						PrimaryDNSIPAddress:   to.StringPtr("<primary-dnsipaddress>"),
						SecondaryDNSIPAddress: to.StringPtr("<secondary-dnsipaddress>"),
					},
					UsernameDownload: &armstoragecache.CacheUsernameDownloadSettings{
						ExtendedGroups: to.BoolPtr(true),
						UsernameSource: armstoragecache.UsernameSourceAD.ToPtr(),
					},
				},
				NetworkSettings: &armstoragecache.CacheNetworkSettings{
					DNSSearchDomain: to.StringPtr("<dnssearch-domain>"),
					DNSServers: []*string{
						to.StringPtr("10.1.22.33"),
						to.StringPtr("10.1.12.33")},
					Mtu:       to.Int32Ptr(1500),
					NtpServer: to.StringPtr("<ntp-server>"),
				},
				SecuritySettings: &armstoragecache.CacheSecuritySettings{
					AccessPolicies: []*armstoragecache.NfsAccessPolicy{
						{
							Name: to.StringPtr("<name>"),
							AccessRules: []*armstoragecache.NfsAccessRule{
								{
									Access:         armstoragecache.NfsAccessRuleAccessRw.ToPtr(),
									RootSquash:     to.BoolPtr(false),
									Scope:          armstoragecache.NfsAccessRuleScopeDefault.ToPtr(),
									SubmountAccess: to.BoolPtr(true),
									Suid:           to.BoolPtr(false),
								}},
						},
						{
							Name: to.StringPtr("<name>"),
							AccessRules: []*armstoragecache.NfsAccessRule{
								{
									Access:         armstoragecache.NfsAccessRuleAccessRw.ToPtr(),
									Filter:         to.StringPtr("<filter>"),
									RootSquash:     to.BoolPtr(false),
									Scope:          armstoragecache.NfsAccessRuleScopeHost.ToPtr(),
									SubmountAccess: to.BoolPtr(true),
									Suid:           to.BoolPtr(true),
								},
								{
									Access:         armstoragecache.NfsAccessRuleAccessRw.ToPtr(),
									Filter:         to.StringPtr("<filter>"),
									RootSquash:     to.BoolPtr(false),
									Scope:          armstoragecache.NfsAccessRuleScopeNetwork.ToPtr(),
									SubmountAccess: to.BoolPtr(true),
									Suid:           to.BoolPtr(true),
								},
								{
									Access:         armstoragecache.NfsAccessRuleAccessNo.ToPtr(),
									AnonymousGID:   to.StringPtr("<anonymous-gid>"),
									AnonymousUID:   to.StringPtr("<anonymous-uid>"),
									RootSquash:     to.BoolPtr(true),
									Scope:          armstoragecache.NfsAccessRuleScopeDefault.ToPtr(),
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
	log.Printf("Cache.ID: %s\n", *res.ID)
}
```
