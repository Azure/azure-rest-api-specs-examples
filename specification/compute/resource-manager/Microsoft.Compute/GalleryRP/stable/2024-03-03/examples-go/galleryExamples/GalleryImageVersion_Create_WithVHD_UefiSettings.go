package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/GalleryImageVersion_Create_WithVHD_UefiSettings.json
func ExampleGalleryImageVersionsClient_BeginCreateOrUpdate_createOrUpdateASimpleGalleryImageVersionUsingVhdAsASourceWithCustomUefiKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryImageVersionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0", armcompute.GalleryImageVersion{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryImageVersionProperties{
			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
				TargetRegions: []*armcompute.TargetRegion{
					{
						Name: to.Ptr("West US"),
						Encryption: &armcompute.EncryptionImages{
							DataDiskImages: []*armcompute.DataDiskImageEncryption{
								{
									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet"),
									Lun:                 to.Ptr[int32](1),
								}},
							OSDiskImage: &armcompute.OSDiskImageEncryption{
								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
							},
						},
						ExcludeFromLatest:    to.Ptr(false),
						RegionalReplicaCount: to.Ptr[int32](1),
					},
					{
						Name:                 to.Ptr("East US"),
						ExcludeFromLatest:    to.Ptr(false),
						RegionalReplicaCount: to.Ptr[int32](2),
						StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardZRS),
					}},
			},
			SafetyProfile: &armcompute.GalleryImageVersionSafetyProfile{
				AllowDeletionOfReplicatedLocations: to.Ptr(false),
				BlockDeletionBeforeEndOfLife:       to.Ptr(false),
			},
			SecurityProfile: &armcompute.ImageVersionSecurityProfile{
				UefiSettings: &armcompute.GalleryImageVersionUefiSettings{
					AdditionalSignatures: &armcompute.UefiKeySignatures{
						Db: []*armcompute.UefiKey{
							{
								Type: to.Ptr(armcompute.UefiKeyTypeX509),
								Value: []*string{
									to.Ptr("<x509 value>")},
							}},
						Dbx: []*armcompute.UefiKey{
							{
								Type: to.Ptr(armcompute.UefiKeyTypeX509),
								Value: []*string{
									to.Ptr("<x509 value>")},
							}},
						Kek: []*armcompute.UefiKey{
							{
								Type: to.Ptr(armcompute.UefiKeyTypeSHA256),
								Value: []*string{
									to.Ptr("<sha256 value>")},
							}},
					},
					SignatureTemplateNames: []*armcompute.UefiSignatureTemplateName{
						to.Ptr(armcompute.UefiSignatureTemplateNameMicrosoftUefiCertificateAuthorityTemplate)},
				},
			},
			StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
				DataDiskImages: []*armcompute.GalleryDataDiskImage{
					{
						HostCaching: to.Ptr(armcompute.HostCachingNone),
						Source: &armcompute.GalleryDiskImageSource{
							StorageAccountID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"),
							URI:              to.Ptr("https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd"),
						},
						Lun: to.Ptr[int32](1),
					}},
				OSDiskImage: &armcompute.GalleryOSDiskImage{
					HostCaching: to.Ptr(armcompute.HostCachingReadOnly),
					Source: &armcompute.GalleryDiskImageSource{
						StorageAccountID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"),
						URI:              to.Ptr("https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd"),
					},
				},
			},
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
	// res.GalleryImageVersion = armcompute.GalleryImageVersion{
	// 	Name: to.Ptr("1.0.0"),
	// 	ID: to.Ptr("/providers/Microsoft.Compute/locations/westus/Galleries/myGalleryName/Images/myGalleryImageName/Versions/1.0.0"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryImageVersionProperties{
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
	// 			PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.000Z"); return t}()),
	// 			ReplicaCount: to.Ptr[int32](1),
	// 			StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			TargetRegions: []*armcompute.TargetRegion{
	// 				{
	// 					Name: to.Ptr("West US"),
	// 					Encryption: &armcompute.EncryptionImages{
	// 						DataDiskImages: []*armcompute.DataDiskImageEncryption{
	// 							{
	// 								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet"),
	// 								Lun: to.Ptr[int32](1),
	// 						}},
	// 						OSDiskImage: &armcompute.OSDiskImageEncryption{
	// 							DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
	// 						},
	// 					},
	// 					ExcludeFromLatest: to.Ptr(false),
	// 					RegionalReplicaCount: to.Ptr[int32](1),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 				},
	// 				{
	// 					Name: to.Ptr("East US"),
	// 					ExcludeFromLatest: to.Ptr(false),
	// 					RegionalReplicaCount: to.Ptr[int32](2),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardZRS),
	// 			}},
	// 		},
	// 		SafetyProfile: &armcompute.GalleryImageVersionSafetyProfile{
	// 			AllowDeletionOfReplicatedLocations: to.Ptr(false),
	// 			BlockDeletionBeforeEndOfLife: to.Ptr(false),
	// 			PolicyViolations: []*armcompute.PolicyViolation{
	// 				{
	// 					Category: to.Ptr(armcompute.PolicyViolationCategoryImageFlaggedUnsafe),
	// 					Details: to.Ptr("This is the policy violation details."),
	// 			}},
	// 			ReportedForPolicyViolation: to.Ptr(true),
	// 		},
	// 		SecurityProfile: &armcompute.ImageVersionSecurityProfile{
	// 			UefiSettings: &armcompute.GalleryImageVersionUefiSettings{
	// 				AdditionalSignatures: &armcompute.UefiKeySignatures{
	// 					Db: []*armcompute.UefiKey{
	// 						{
	// 							Type: to.Ptr(armcompute.UefiKeyTypeX509),
	// 							Value: []*string{
	// 								to.Ptr("<x509 value>")},
	// 						}},
	// 						Dbx: []*armcompute.UefiKey{
	// 							{
	// 								Type: to.Ptr(armcompute.UefiKeyTypeX509),
	// 								Value: []*string{
	// 									to.Ptr("<x509 value>")},
	// 							}},
	// 							Kek: []*armcompute.UefiKey{
	// 								{
	// 									Type: to.Ptr(armcompute.UefiKeyTypeSHA256),
	// 									Value: []*string{
	// 										to.Ptr("<sha256 value>")},
	// 								}},
	// 							},
	// 							SignatureTemplateNames: []*armcompute.UefiSignatureTemplateName{
	// 								to.Ptr(armcompute.UefiSignatureTemplateNameMicrosoftUefiCertificateAuthorityTemplate)},
	// 							},
	// 						},
	// 						StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
	// 							DataDiskImages: []*armcompute.GalleryDataDiskImage{
	// 								{
	// 									HostCaching: to.Ptr(armcompute.HostCachingNone),
	// 									Source: &armcompute.GalleryDiskImageSource{
	// 										StorageAccountID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"),
	// 										URI: to.Ptr("https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd"),
	// 									},
	// 									Lun: to.Ptr[int32](1),
	// 							}},
	// 							OSDiskImage: &armcompute.GalleryOSDiskImage{
	// 								HostCaching: to.Ptr(armcompute.HostCachingReadOnly),
	// 								Source: &armcompute.GalleryDiskImageSource{
	// 									StorageAccountID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"),
	// 									URI: to.Ptr("https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				}
}
