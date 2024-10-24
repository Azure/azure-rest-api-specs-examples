package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/VolumeGroups_Get_MaximumSet_Gen.json
func ExampleVolumeGroupsClient_Get_volumeGroupsGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeGroupsClient().Get(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VolumeGroup = armelasticsan.VolumeGroup{
	// 	Name: to.Ptr("qymuqyvdlpshrna"),
	// 	Type: to.Ptr("wwkffcgidqktzuzo"),
	// 	ID: to.Ptr("cgwmakehxvhv"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-09T08:26:45.352Z"); return t}()),
	// 		CreatedBy: to.Ptr("iskcypymrroqhmdwccwqzdclpyutpb"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-09T08:26:45.352Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("ujpptnmvqtzudoiaihbgrvnagxh"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Identity: &armelasticsan.Identity{
	// 		Type: to.Ptr(armelasticsan.IdentityTypeNone),
	// 		PrincipalID: to.Ptr("uqclmeopt"),
	// 		TenantID: to.Ptr("qvhlvbvgodtycsfjhamstlrusktly"),
	// 		UserAssignedIdentities: map[string]*armelasticsan.UserAssignedIdentity{
	// 			"key1006": &armelasticsan.UserAssignedIdentity{
	// 				ClientID: to.Ptr("hsfmzocbkqpxspqwamdqjknd"),
	// 				PrincipalID: to.Ptr("pcbbifvgblsrhzmpdt"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armelasticsan.VolumeGroupProperties{
	// 		Encryption: to.Ptr(armelasticsan.EncryptionTypeEncryptionAtRestWithPlatformKey),
	// 		EncryptionProperties: &armelasticsan.EncryptionProperties{
	// 			EncryptionIdentity: &armelasticsan.EncryptionIdentity{
	// 				EncryptionUserAssignedIdentity: to.Ptr("gfhkfbozahmmwluqndfgxunssafa"),
	// 			},
	// 			KeyVaultProperties: &armelasticsan.KeyVaultProperties{
	// 				CurrentVersionedKeyExpirationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-09T08:26:49.141Z"); return t}()),
	// 				CurrentVersionedKeyIdentifier: to.Ptr("giyyriknaomjipru"),
	// 				KeyName: to.Ptr("lunpapamzeimppgobraxjt"),
	// 				KeyVaultURI: to.Ptr("https://microsoft.com/a"),
	// 				KeyVersion: to.Ptr("oemygbnfmqhijmonkqfqmy"),
	// 				LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-09T08:26:49.141Z"); return t}()),
	// 			},
	// 		},
	// 		EnforceDataIntegrityCheckForIscsi: to.Ptr(true),
	// 		NetworkACLs: &armelasticsan.NetworkRuleSet{
	// 			VirtualNetworkRules: []*armelasticsan.VirtualNetworkRule{
	// 				{
	// 					Action: to.Ptr(armelasticsan.ActionAllow),
	// 					VirtualNetworkResourceID: to.Ptr("bkhwaiqvvaguymsmnzzbzz"),
	// 			}},
	// 		},
	// 		PrivateEndpointConnections: []*armelasticsan.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("duzwhl"),
	// 				Type: to.Ptr("mzwnrqbkzgk"),
	// 				ID: to.Ptr("utnp"),
	// 				SystemData: &armelasticsan.SystemData{
	// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-09T08:26:45.352Z"); return t}()),
	// 					CreatedBy: to.Ptr("iskcypymrroqhmdwccwqzdclpyutpb"),
	// 					CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-09T08:26:45.352Z"); return t}()),
	// 					LastModifiedBy: to.Ptr("ujpptnmvqtzudoiaihbgrvnagxh"),
	// 					LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 				},
	// 				Properties: &armelasticsan.PrivateEndpointConnectionProperties{
	// 					GroupIDs: []*string{
	// 						to.Ptr("ftdxwynrey")},
	// 						PrivateEndpoint: &armelasticsan.PrivateEndpoint{
	// 							ID: to.Ptr("kdvywmjnrvqksyvxgosjorlzjk"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("hqsksvdlhacbmawvhlhhfmcyv"),
	// 							ActionsRequired: to.Ptr("xvnopczgivazrjlzirhtww"),
	// 							Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
	// 						},
	// 						ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 					},
	// 			}},
	// 			ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeIscsi),
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		},
	// 	}
}
