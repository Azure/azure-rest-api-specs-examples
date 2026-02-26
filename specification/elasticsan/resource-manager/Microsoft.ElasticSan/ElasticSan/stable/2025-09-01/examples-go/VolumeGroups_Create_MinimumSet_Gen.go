package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: 2025-09-01/VolumeGroups_Create_MinimumSet_Gen.json
func ExampleVolumeGroupsClient_BeginCreate_volumeGroupsCreateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeGroupsClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", armelasticsan.VolumeGroup{}, nil)
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
	// res = armelasticsan.VolumeGroupsClientCreateResponse{
	// 	VolumeGroup: &armelasticsan.VolumeGroup{
	// 		Name: to.Ptr("cr"),
	// 		Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"),
	// 		Identity: &armelasticsan.Identity{
	// 			Type: to.Ptr(armelasticsan.IdentityTypeNone),
	// 			PrincipalID: to.Ptr("ihsiwrwdofymkhquaxcrtfmmrsygw"),
	// 			TenantID: to.Ptr("gtkzkjsy"),
	// 			UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
	// 				"key7482": &armelasticsan.UserAssignedIdentity{
	// 					ClientID: to.Ptr("jaczsquolgxwpznljbmdupn"),
	// 					PrincipalID: to.Ptr("vfdzizicxcfcqecgsmshz"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armelasticsan.VolumeGroupProperties{
	// 			Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
	// 			EncryptionProperties: &armelasticsan.EncryptionProperties{
	// 				EncryptionIdentity: &armelasticsan.EncryptionIdentity{
	// 					EncryptionUserAssignedIdentity: to.Ptr("im"),
	// 				},
	// 				KeyVaultProperties: &armelasticsan.KeyVaultProperties{
	// 					CurrentVersionedKeyExpirationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
	// 					CurrentVersionedKeyIdentifier: to.Ptr("rnpxhtzkquzyoepwbwktbwb"),
	// 					KeyName: to.Ptr("sftaiernmrzypnrkpakrrawxcbsqzc"),
	// 					KeyVaultURI: to.Ptr("https://microsoft.com/axmblwp"),
	// 					KeyVersion: to.Ptr("c"),
	// 					LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:11.388Z"); return t}()),
	// 				},
	// 			},
	// 			EnforceDataIntegrityCheckForIscsi: to.Ptr(true),
	// 			NetworkACLs: &armelasticsan.NetworkRuleSet{
	// 				VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
	// 					{
	// 						Action: to.Ptr(armelasticsan.ActionAllow),
	// 						VirtualNetworkResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}"),
	// 					},
	// 				},
	// 			},
	// 			PrivateEndpointConnections: []*armelasticsan.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("gewxykc"),
	// 					Type: to.Ptr("ailymcedgvxbqklmqtlty"),
	// 					ID: to.Ptr("opcjchensdf"),
	// 					Properties: &armelasticsan.PrivateEndpointConnectionProperties{
	// 						GroupIDs: []*string{
	// 							to.Ptr("bolviufgqnyid"),
	// 						},
	// 						PrivateEndpoint: &armelasticsan.PrivateEndpoint{
	// 							ID: to.Ptr("ehxmltubeltzmgcqxocakaansat"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("nahklgxicbqjbbvcdrkljqdhprruys"),
	// 							ActionsRequired: to.Ptr("sairafcqpvucoy"),
	// 							Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
	// 						},
	// 						ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesInvalid),
	// 					},
	// 					SystemData: &armelasticsan.SystemData{
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 						CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 						CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("bcclmbseed"),
	// 						LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 					},
	// 				},
	// 			},
	// 			ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		},
	// 		SystemData: &armelasticsan.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 			CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 			CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-23T12:16:10.057Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("bcclmbseed"),
	// 			LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
