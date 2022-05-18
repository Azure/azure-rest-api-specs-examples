Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragecache%2Farmstoragecache%2Fv1.0.0/sdk/resourcemanager/storagecache/armstoragecache/README.md) on how to add the SDK to your project and authenticate.

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
	}
	ctx := context.Background()
	client, err := armstoragecache.NewCachesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"scgroup",
		"sc1",
		&armstoragecache.CachesClientUpdateOptions{Cache: &armstoragecache.Cache{
			Location: to.Ptr("westus"),
			Properties: &armstoragecache.CacheProperties{
				CacheSizeGB: to.Ptr[int32](3072),
				DirectoryServicesSettings: &armstoragecache.CacheDirectorySettings{
					ActiveDirectory: &armstoragecache.CacheActiveDirectorySettings{
						CacheNetBiosName:      to.Ptr("contosoSmb"),
						DomainName:            to.Ptr("contosoAd.contoso.local"),
						DomainNetBiosName:     to.Ptr("contosoAd"),
						PrimaryDNSIPAddress:   to.Ptr("192.0.2.10"),
						SecondaryDNSIPAddress: to.Ptr("192.0.2.11"),
					},
					UsernameDownload: &armstoragecache.CacheUsernameDownloadSettings{
						ExtendedGroups: to.Ptr(true),
						UsernameSource: to.Ptr(armstoragecache.UsernameSourceAD),
					},
				},
				NetworkSettings: &armstoragecache.CacheNetworkSettings{
					DNSSearchDomain: to.Ptr("contoso.com"),
					DNSServers: []*string{
						to.Ptr("10.1.22.33"),
						to.Ptr("10.1.12.33")},
					Mtu:       to.Ptr[int32](1500),
					NtpServer: to.Ptr("time.contoso.com"),
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
						},
						{
							Name: to.Ptr("restrictive"),
							AccessRules: []*armstoragecache.NfsAccessRule{
								{
									Access:         to.Ptr(armstoragecache.NfsAccessRuleAccessRw),
									Filter:         to.Ptr("10.99.3.145"),
									RootSquash:     to.Ptr(false),
									Scope:          to.Ptr(armstoragecache.NfsAccessRuleScopeHost),
									SubmountAccess: to.Ptr(true),
									Suid:           to.Ptr(true),
								},
								{
									Access:         to.Ptr(armstoragecache.NfsAccessRuleAccessRw),
									Filter:         to.Ptr("10.99.1.0/24"),
									RootSquash:     to.Ptr(false),
									Scope:          to.Ptr(armstoragecache.NfsAccessRuleScopeNetwork),
									SubmountAccess: to.Ptr(true),
									Suid:           to.Ptr(true),
								},
								{
									Access:         to.Ptr(armstoragecache.NfsAccessRuleAccessNo),
									AnonymousGID:   to.Ptr("65534"),
									AnonymousUID:   to.Ptr("65534"),
									RootSquash:     to.Ptr(true),
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
	// TODO: use response item
	_ = res
}
```
