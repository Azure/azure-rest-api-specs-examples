package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan/v2"
)

// Generated from example definition: 2026-05-01-preview/VolumeGroups_Create_MaximumSet_Gen.json
func ExampleVolumeGroupsClient_BeginCreate_volumeGroupsCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeGroupsClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", armelasticsan.VolumeGroup{
		Identity: &armelasticsan.Identity{
			Type: to.Ptr(armelasticsan.IdentityTypeNone),
			UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
				"key2350": {},
			},
		},
		Properties: &armelasticsan.VolumeGroupProperties{
			DeleteRetentionPolicy: &armelasticsan.DeleteRetentionPolicy{
				PolicyState:         to.Ptr(armelasticsan.PolicyStateEnabled),
				RetentionPeriodDays: to.Ptr[int32](14),
			},
			Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
			EncryptionProperties: &armelasticsan.EncryptionProperties{
				EncryptionIdentity: &armelasticsan.EncryptionIdentity{
					EncryptionUserAssignedIdentity: to.Ptr("vgbeephfgecgg"),
				},
				KeyVaultProperties: &armelasticsan.KeyVaultProperties{
					KeyName:     to.Ptr("rommjwp"),
					KeyVaultURI: to.Ptr("https://microsoft.com/at"),
					KeyVersion:  to.Ptr("ulmxxgzgsuhalwesmhfslq"),
				},
			},
			EnforceDataIntegrityCheckForIscsi: to.Ptr(true),
			EncryptionInTransit:               to.Ptr(true),
			NetworkACLs: &armelasticsan.NetworkRuleSet{
				VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
					{
						Action:                   to.Ptr(armelasticsan.ActionAllow),
						VirtualNetworkResourceID: to.Ptr("fhhawhc"),
					},
				},
			},
			ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armelasticsan.VolumeGroupsClientCreateResponse{
	// 	VolumeGroup: armelasticsan.VolumeGroup{
	// 		Name: to.Ptr("dov"),
	// 		Type: to.Ptr("kg"),
	// 		ID: to.Ptr("hoazltxzojzwgzohjnh"),
	// 		Identity: &armelasticsan.Identity{
	// 			Type: to.Ptr(armelasticsan.IdentityTypeNone),
	// 			PrincipalID: to.Ptr("zqobj"),
	// 			TenantID: to.Ptr("douwo"),
	// 			UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
	// 				"key2350": &armelasticsan.UserAssignedIdentity{
	// 					ClientID: to.Ptr("ddhoilirjxushxvxttgqh"),
	// 					PrincipalID: to.Ptr("lmhozfpeu"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armelasticsan.VolumeGroupProperties{
	// 			DeleteRetentionPolicy: &armelasticsan.DeleteRetentionPolicy{
	// 				PolicyState: to.Ptr(armelasticsan.PolicyStateEnabled),
	// 				RetentionPeriodDays: to.Ptr[int32](14),
	// 			},
	// 			Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
	// 			EncryptionProperties: &armelasticsan.EncryptionProperties{
	// 				EncryptionIdentity: &armelasticsan.EncryptionIdentity{
	// 					EncryptionUserAssignedIdentity: to.Ptr("vgbeephfgecgg"),
	// 				},
	// 				KeyVaultProperties: &armelasticsan.KeyVaultProperties{
	// 					CurrentVersionedKeyExpirationTimestamp: to.Ptr(time.Date(2024, time.April, 29, 14, 22, 25, 155000000, time.UTC)),
	// 					CurrentVersionedKeyIdentifier: to.Ptr("bqgwaoezxtvwuydxxvsecod"),
	// 					KeyName: to.Ptr("rommjwp"),
	// 					KeyVaultURI: to.Ptr("https://microsoft.com/at"),
	// 					KeyVersion: to.Ptr("ulmxxgzgsuhalwesmhfslq"),
	// 					LastKeyRotationTimestamp: to.Ptr(time.Date(2024, time.April, 29, 14, 22, 25, 155000000, time.UTC)),
	// 				},
	// 			},
	// 			EnforceDataIntegrityCheckForIscsi: to.Ptr(true),
	// 			EncryptionInTransit: to.Ptr(true),
	// 			NetworkACLs: &armelasticsan.NetworkRuleSet{
	// 				VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
	// 					{
	// 						Action: to.Ptr(armelasticsan.ActionAllow),
	// 						VirtualNetworkResourceID: to.Ptr("fhhawhc"),
	// 					},
	// 				},
	// 			},
	// 			PrivateEndpointConnections: []*armelasticsan.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("qyvurspdwqwfiurkgjklvpdnoli"),
	// 					Type: to.Ptr("qvpdadjcuxksssnjplnpv"),
	// 					ID: to.Ptr("vdm"),
	// 					Properties: &armelasticsan.PrivateEndpointConnectionProperties{
	// 						GroupIDs: []*string{
	// 							to.Ptr("f"),
	// 						},
	// 						PrivateEndpoint: &armelasticsan.PrivateEndpoint{
	// 							ID: to.Ptr("jzaucqmfvqetawalnsqbisqkfokbnj"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("wflnhkypjiarhhobagelhjlcsqdtt"),
	// 							ActionsRequired: to.Ptr("hcqnszybqdmdbtuumpvrgbggj"),
	// 							Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
	// 						},
	// 						ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesInvalid),
	// 					},
	// 					SystemData: &armelasticsan.SystemData{
	// 						CreatedAt: to.Ptr(time.Date(2024, time.April, 29, 14, 22, 13, 546000000, time.UTC)),
	// 						CreatedBy: to.Ptr("bpuxtfzqwdhifevjtucoc"),
	// 						CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(time.Date(2024, time.April, 29, 14, 22, 13, 547000000, time.UTC)),
	// 						LastModifiedBy: to.Ptr("ourjjlolgugpxnkbiegumkicksibep"),
	// 						LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 					},
	// 				},
	// 			},
	// 			ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesInvalid),
	// 		},
	// 		SystemData: &armelasticsan.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2024, time.April, 29, 14, 22, 13, 546000000, time.UTC)),
	// 			CreatedBy: to.Ptr("bpuxtfzqwdhifevjtucoc"),
	// 			CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2024, time.April, 29, 14, 22, 13, 547000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("ourjjlolgugpxnkbiegumkicksibep"),
	// 			LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
