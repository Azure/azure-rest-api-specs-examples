package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan/v2"
)

// Generated from example definition: 2026-05-01-preview/VolumeGroups_PerformanceCritical_Create_MinimumSet_Gen.json
func ExampleVolumeGroupsClient_BeginCreate_volumeGroupsPerformanceCriticalCreateMinimumSetGen() {
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
		Properties: &armelasticsan.VolumeGroupProperties{
			ProtocolType:     to.Ptr(armelasticsan.StorageTargetTypeDirectAttach),
			QualityOfService: to.Ptr(armelasticsan.QualityOfServicePerformanceCritical),
			ReservedIops:     to.Ptr[int32](10000),
			ReservedMBps:     to.Ptr[int32](800),
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
	// 					CurrentVersionedKeyExpirationTimestamp: to.Ptr(time.Date(2023, time.August, 23, 12, 16, 11, 388000000, time.UTC)),
	// 					CurrentVersionedKeyIdentifier: to.Ptr("rnpxhtzkquzyoepwbwktbwb"),
	// 					KeyName: to.Ptr("sftaiernmrzypnrkpakrrawxcbsqzc"),
	// 					KeyVaultURI: to.Ptr("https://microsoft.com/axmblwp"),
	// 					KeyVersion: to.Ptr("c"),
	// 					LastKeyRotationTimestamp: to.Ptr(time.Date(2023, time.August, 23, 12, 16, 11, 388000000, time.UTC)),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			ProtocolType: to.Ptr(armelasticsan.StorageTargetTypeDirectAttach),
	// 			QualityOfService: to.Ptr(armelasticsan.QualityOfServicePerformanceCritical),
	// 			ReservedIops: to.Ptr[int32](10000),
	// 			ReservedMBps: to.Ptr[int32](800),
	// 		},
	// 		SystemData: &armelasticsan.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2023, time.August, 23, 12, 16, 10, 57000000, time.UTC)),
	// 			CreatedBy: to.Ptr("kakcyehdrphqkilgkhpbdtvpupak"),
	// 			CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2023, time.August, 23, 12, 16, 10, 57000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("bcclmbseed"),
	// 			LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
