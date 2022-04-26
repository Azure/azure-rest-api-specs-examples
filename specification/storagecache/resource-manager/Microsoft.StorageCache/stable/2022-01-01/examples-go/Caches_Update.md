Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragecache%2Farmstoragecache%2Fv0.4.0/sdk/resourcemanager/storagecache/armstoragecache/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/Caches_Update.json
func ExampleCachesClient_Update() {
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
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<cache-name>",
		&armstoragecache.CachesClientUpdateOptions{Cache: &armstoragecache.Cache{
			Location: to.Ptr("<location>"),
			Properties: &armstoragecache.CacheProperties{
				CacheSizeGB: to.Ptr[int32](3072),
				DirectoryServicesSettings: &armstoragecache.CacheDirectorySettings{
					ActiveDirectory: &armstoragecache.CacheActiveDirectorySettings{
						CacheNetBiosName:      to.Ptr("<cache-net-bios-name>"),
						DomainName:            to.Ptr("<domain-name>"),
						DomainNetBiosName:     to.Ptr("<domain-net-bios-name>"),
						PrimaryDNSIPAddress:   to.Ptr("<primary-dnsipaddress>"),
						SecondaryDNSIPAddress: to.Ptr("<secondary-dnsipaddress>"),
					},
					UsernameDownload: &armstoragecache.CacheUsernameDownloadSettings{
						ExtendedGroups: to.Ptr(true),
						UsernameSource: to.Ptr(armstoragecache.UsernameSourceAD),
					},
				},
				NetworkSettings: &armstoragecache.CacheNetworkSettings{
					DNSSearchDomain: to.Ptr("<dnssearch-domain>"),
					DNSServers: []*string{
						to.Ptr("10.1.22.33"),
						to.Ptr("10.1.12.33")},
					Mtu:       to.Ptr[int32](1500),
					NtpServer: to.Ptr("<ntp-server>"),
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
						},
						{
							Name: to.Ptr("<name>"),
							AccessRules: []*armstoragecache.NfsAccessRule{
								{
									Access:         to.Ptr(armstoragecache.NfsAccessRuleAccessRw),
									Filter:         to.Ptr("<filter>"),
									RootSquash:     to.Ptr(false),
									Scope:          to.Ptr(armstoragecache.NfsAccessRuleScopeHost),
									SubmountAccess: to.Ptr(true),
									Suid:           to.Ptr(true),
								},
								{
									Access:         to.Ptr(armstoragecache.NfsAccessRuleAccessRw),
									Filter:         to.Ptr("<filter>"),
									RootSquash:     to.Ptr(false),
									Scope:          to.Ptr(armstoragecache.NfsAccessRuleScopeNetwork),
									SubmountAccess: to.Ptr(true),
									Suid:           to.Ptr(true),
								},
								{
									Access:         to.Ptr(armstoragecache.NfsAccessRuleAccessNo),
									AnonymousGID:   to.Ptr("<anonymous-gid>"),
									AnonymousUID:   to.Ptr("<anonymous-uid>"),
									RootSquash:     to.Ptr(true),
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
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
