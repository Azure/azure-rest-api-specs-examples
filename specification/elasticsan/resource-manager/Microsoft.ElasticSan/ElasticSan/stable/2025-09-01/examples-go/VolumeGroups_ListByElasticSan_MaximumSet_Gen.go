package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: 2025-09-01/VolumeGroups_ListByElasticSan_MaximumSet_Gen.json
func ExampleVolumeGroupsClient_NewListByElasticSanPager_volumeGroupsListByElasticSanMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeGroupsClient().NewListByElasticSanPager("resourcegroupname", "elasticsanname", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armelasticsan.VolumeGroupsClientListByElasticSanResponse{
		// 	VolumeGroupList: armelasticsan.VolumeGroupList{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/subscriptionid/resourceGroups/resourcegroupname/providers/Microsoft.ElasticSan/elasticSans/elasticsanname/volumegroups?api-version=2024-07-01-preview&%24skiptoken=xyz567abc890"),
		// 		Value: []*armelasticsan.VolumeGroup{
		// 			{
		// 				Name: to.Ptr("dov"),
		// 				Type: to.Ptr("kg"),
		// 				ID: to.Ptr("hoazltxzojzwgzohjnh"),
		// 				Identity: &armelasticsan.Identity{
		// 					Type: to.Ptr(armelasticsan.IdentityTypeNone),
		// 					PrincipalID: to.Ptr("zqobj"),
		// 					TenantID: to.Ptr("douwo"),
		// 					UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
		// 						"key2350": &armelasticsan.UserAssignedIdentity{
		// 							ClientID: to.Ptr("ddhoilirjxushxvxttgqh"),
		// 							PrincipalID: to.Ptr("lmhozfpeu"),
		// 						},
		// 					},
		// 				},
		// 				Properties: &armelasticsan.VolumeGroupProperties{
		// 					Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
		// 					EncryptionProperties: &armelasticsan.EncryptionProperties{
		// 						EncryptionIdentity: &armelasticsan.EncryptionIdentity{
		// 							EncryptionUserAssignedIdentity: to.Ptr("vgbeephfgecgg"),
		// 						},
		// 						KeyVaultProperties: &armelasticsan.KeyVaultProperties{
		// 							CurrentVersionedKeyExpirationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T14:22:25.155Z"); return t}()),
		// 							CurrentVersionedKeyIdentifier: to.Ptr("bqgwaoezxtvwuydxxvsecod"),
		// 							KeyName: to.Ptr("rommjwp"),
		// 							KeyVaultURI: to.Ptr("https://microsoft.com/at"),
		// 							KeyVersion: to.Ptr("ulmxxgzgsuhalwesmhfslq"),
		// 							LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T14:22:25.155Z"); return t}()),
		// 						},
		// 					},
		// 					EnforceDataIntegrityCheckForIscsi: to.Ptr(true),
		// 					NetworkACLs: &armelasticsan.NetworkRuleSet{
		// 						VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
		// 							{
		// 								Action: to.Ptr(armelasticsan.ActionAllow),
		// 								VirtualNetworkResourceID: to.Ptr("fhhawhc"),
		// 							},
		// 						},
		// 					},
		// 					PrivateEndpointConnections: []*armelasticsan.PrivateEndpointConnection{
		// 						{
		// 							Name: to.Ptr("qyvurspdwqwfiurkgjklvpdnoli"),
		// 							Type: to.Ptr("qvpdadjcuxksssnjplnpv"),
		// 							ID: to.Ptr("vdm"),
		// 							Properties: &armelasticsan.PrivateEndpointConnectionProperties{
		// 								GroupIDs: []*string{
		// 									to.Ptr("f"),
		// 								},
		// 								PrivateEndpoint: &armelasticsan.PrivateEndpoint{
		// 									ID: to.Ptr("jzaucqmfvqetawalnsqbisqkfokbnj"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
		// 									Description: to.Ptr("wflnhkypjiarhhobagelhjlcsqdtt"),
		// 									ActionsRequired: to.Ptr("hcqnszybqdmdbtuumpvrgbggj"),
		// 									Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
		// 								},
		// 								ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesInvalid),
		// 							},
		// 							SystemData: &armelasticsan.SystemData{
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T14:22:13.546Z"); return t}()),
		// 								CreatedBy: to.Ptr("bpuxtfzqwdhifevjtucoc"),
		// 								CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T14:22:13.547Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("ourjjlolgugpxnkbiegumkicksibep"),
		// 								LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 							},
		// 						},
		// 					},
		// 					ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
		// 					ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesInvalid),
		// 				},
		// 				SystemData: &armelasticsan.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T14:22:13.546Z"); return t}()),
		// 					CreatedBy: to.Ptr("bpuxtfzqwdhifevjtucoc"),
		// 					CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T14:22:13.547Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("ourjjlolgugpxnkbiegumkicksibep"),
		// 					LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
