package armstoragecache_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/Caches_CreateOrUpdate.json
func ExampleCachesClient_BeginCreateOrUpdate_cachesCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCachesClient().BeginCreateOrUpdate(ctx, "scgroup", "sc1", armstoragecache.Cache{
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
			UpgradeSettings: &armstoragecache.CacheUpgradeSettings{
				ScheduledTime:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-26T18:25:43.511Z"); return t }()),
				UpgradeScheduleEnabled: to.Ptr(true),
			},
		},
		SKU: &armstoragecache.CacheSKU{
			Name: to.Ptr("Standard_2G"),
		},
		Tags: map[string]*string{
			"Dept": to.Ptr("Contoso"),
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
	// res.Cache = armstoragecache.Cache{
	// 	Name: to.Ptr("sc1"),
	// 	Type: to.Ptr("Microsoft.StorageCache/Cache"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armstoragecache.CacheProperties{
	// 		CacheSizeGB: to.Ptr[int32](3072),
	// 		DirectoryServicesSettings: &armstoragecache.CacheDirectorySettings{
	// 			ActiveDirectory: &armstoragecache.CacheActiveDirectorySettings{
	// 				CacheNetBiosName: to.Ptr("contosoSmb"),
	// 				DomainJoined: to.Ptr(armstoragecache.DomainJoinedTypeYes),
	// 				DomainName: to.Ptr("contosoAd.contoso.local"),
	// 				DomainNetBiosName: to.Ptr("contosoAd"),
	// 				PrimaryDNSIPAddress: to.Ptr("192.0.2.10"),
	// 				SecondaryDNSIPAddress: to.Ptr("192.0.2.11"),
	// 			},
	// 			UsernameDownload: &armstoragecache.CacheUsernameDownloadSettings{
	// 				AutoDownloadCertificate: to.Ptr(false),
	// 				CaCertificateURI: to.Ptr("http://contoso.net/cacert.pem"),
	// 				EncryptLdapConnection: to.Ptr(false),
	// 				ExtendedGroups: to.Ptr(true),
	// 				GroupFileURI: to.Ptr("http://contoso.net/group.file"),
	// 				LdapBaseDN: to.Ptr("dc=contosoad,dc=contoso,dc=local"),
	// 				LdapServer: to.Ptr("192.0.2.12"),
	// 				RequireValidCertificate: to.Ptr(false),
	// 				UserFileURI: to.Ptr("http://contoso.net/passwd.file"),
	// 				UsernameDownloaded: to.Ptr(armstoragecache.UsernameDownloadedTypeYes),
	// 				UsernameSource: to.Ptr(armstoragecache.UsernameSourceLDAP),
	// 			},
	// 		},
	// 		EncryptionSettings: &armstoragecache.CacheEncryptionSettings{
	// 			KeyEncryptionKey: &armstoragecache.KeyVaultKeyReference{
	// 				KeyURL: to.Ptr("https://keyvault-cmk.vault.azure.net/keys/key2048/test"),
	// 				SourceVault: &armstoragecache.KeyVaultKeyReferenceSourceVault{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"),
	// 				},
	// 			},
	// 		},
	// 		Health: &armstoragecache.CacheHealth{
	// 			Conditions: []*armstoragecache.Condition{
	// 				{
	// 					Message: to.Ptr("Cannot contact DNS server"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-21T18:25:43.511Z"); return t}()),
	// 			}},
	// 			State: to.Ptr(armstoragecache.HealthStateTypeTransitioning),
	// 			StatusDescription: to.Ptr("Cache is being created."),
	// 		},
	// 		MountAddresses: []*string{
	// 			to.Ptr("192.168.1.1"),
	// 			to.Ptr("192.168.1.2")},
	// 			NetworkSettings: &armstoragecache.CacheNetworkSettings{
	// 				DNSSearchDomain: to.Ptr("contoso.com"),
	// 				DNSServers: []*string{
	// 					to.Ptr("10.1.22.33"),
	// 					to.Ptr("10.1.12.33")},
	// 					Mtu: to.Ptr[int32](1500),
	// 					NtpServer: to.Ptr("time.contoso.com"),
	// 				},
	// 				PrimingJobs: []*armstoragecache.PrimingJob{
	// 					{
	// 						PrimingJobDetails: to.Ptr("Files: Cached=635, Failed=0, Excluded=80, Data=346030 bytes, Directories: Cached=1003, Failed=0, Excluded=0"),
	// 						PrimingJobID: to.Ptr("00000000000_0000000000"),
	// 						PrimingJobName: to.Ptr("contosoJob1"),
	// 						PrimingJobPercentComplete: to.Ptr[float64](100),
	// 						PrimingJobState: to.Ptr(armstoragecache.PrimingJobStateComplete),
	// 						PrimingJobStatus: to.Ptr("success"),
	// 					},
	// 					{
	// 						PrimingJobDetails: to.Ptr(""),
	// 						PrimingJobID: to.Ptr("11111111111_1111111111"),
	// 						PrimingJobName: to.Ptr("contosoJob2"),
	// 						PrimingJobPercentComplete: to.Ptr[float64](0),
	// 						PrimingJobState: to.Ptr(armstoragecache.PrimingJobStateQueued),
	// 						PrimingJobStatus: to.Ptr(""),
	// 				}},
	// 				ProvisioningState: to.Ptr(armstoragecache.ProvisioningStateTypeSucceeded),
	// 				SecuritySettings: &armstoragecache.CacheSecuritySettings{
	// 					AccessPolicies: []*armstoragecache.NfsAccessPolicy{
	// 						{
	// 							Name: to.Ptr("default"),
	// 							AccessRules: []*armstoragecache.NfsAccessRule{
	// 								{
	// 									Access: to.Ptr(armstoragecache.NfsAccessRuleAccessRw),
	// 									RootSquash: to.Ptr(false),
	// 									Scope: to.Ptr(armstoragecache.NfsAccessRuleScopeDefault),
	// 									SubmountAccess: to.Ptr(true),
	// 									Suid: to.Ptr(false),
	// 							}},
	// 					}},
	// 				},
	// 				SpaceAllocation: []*armstoragecache.StorageTargetSpaceAllocation{
	// 				},
	// 				Subnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/sub1"),
	// 				UpgradeSettings: &armstoragecache.CacheUpgradeSettings{
	// 					ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-26T18:25:43.511Z"); return t}()),
	// 					UpgradeScheduleEnabled: to.Ptr(true),
	// 				},
	// 				UpgradeStatus: &armstoragecache.CacheUpgradeStatus{
	// 					CurrentFirmwareVersion: to.Ptr("2022.08.1"),
	// 					FirmwareUpdateDeadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T18:25:43.511Z"); return t}()),
	// 					FirmwareUpdateStatus: to.Ptr(armstoragecache.FirmwareStatusTypeAvailable),
	// 					LastFirmwareUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-21T18:25:43.511Z"); return t}()),
	// 					PendingFirmwareVersion: to.Ptr("2022.08.1"),
	// 				},
	// 				Zones: []*string{
	// 					to.Ptr("1")},
	// 				},
	// 				SKU: &armstoragecache.CacheSKU{
	// 					Name: to.Ptr("Standard_2G"),
	// 				},
	// 				SystemData: &armstoragecache.SystemData{
	// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 					CreatedBy: to.Ptr("user1"),
	// 					CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 					LastModifiedBy: to.Ptr("user2"),
	// 					LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 				},
	// 				Tags: map[string]*string{
	// 					"Dept": to.Ptr("Contoso"),
	// 				},
	// 			}
}
