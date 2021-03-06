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
		&armstoragecache.CachesClientUpdateOptions{Cache: &armstoragecache.Cache{
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
						UsernameSource: armstoragecache.UsernameSource("AD").ToPtr(),
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
									Access:         armstoragecache.NfsAccessRuleAccess("rw").ToPtr(),
									RootSquash:     to.BoolPtr(false),
									Scope:          armstoragecache.NfsAccessRuleScope("default").ToPtr(),
									SubmountAccess: to.BoolPtr(true),
									Suid:           to.BoolPtr(false),
								}},
						},
						{
							Name: to.StringPtr("<name>"),
							AccessRules: []*armstoragecache.NfsAccessRule{
								{
									Access:         armstoragecache.NfsAccessRuleAccess("rw").ToPtr(),
									Filter:         to.StringPtr("<filter>"),
									RootSquash:     to.BoolPtr(false),
									Scope:          armstoragecache.NfsAccessRuleScope("host").ToPtr(),
									SubmountAccess: to.BoolPtr(true),
									Suid:           to.BoolPtr(true),
								},
								{
									Access:         armstoragecache.NfsAccessRuleAccess("rw").ToPtr(),
									Filter:         to.StringPtr("<filter>"),
									RootSquash:     to.BoolPtr(false),
									Scope:          armstoragecache.NfsAccessRuleScope("network").ToPtr(),
									SubmountAccess: to.BoolPtr(true),
									Suid:           to.BoolPtr(true),
								},
								{
									Access:         armstoragecache.NfsAccessRuleAccess("no").ToPtr(),
									AnonymousGID:   to.StringPtr("<anonymous-gid>"),
									AnonymousUID:   to.StringPtr("<anonymous-uid>"),
									RootSquash:     to.BoolPtr(true),
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
	log.Printf("Response result: %#v\n", res.CachesClientUpdateResult)
}
