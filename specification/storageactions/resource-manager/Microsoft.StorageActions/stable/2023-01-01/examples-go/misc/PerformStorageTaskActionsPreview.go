package armstorageactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageactions/armstorageactions"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/misc/PerformStorageTaskActionsPreview.json
func ExampleStorageTasksClient_PreviewActions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorageactions.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStorageTasksClient().PreviewActions(ctx, "eastus", armstorageactions.StorageTaskPreviewAction{
		Properties: &armstorageactions.StorageTaskPreviewActionProperties{
			Action: &armstorageactions.StorageTaskPreviewActionCondition{
				ElseBlockExists: to.Ptr(true),
				If: &armstorageactions.StorageTaskPreviewActionIfCondition{
					Condition: to.Ptr("[[equals(AccessTier, 'Hot')]]"),
				},
			},
			Blobs: []*armstorageactions.StorageTaskPreviewBlobProperties{
				{
					Name: to.Ptr("folder1/file1.txt"),
					Metadata: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
						{
							Key:   to.Ptr("mKey1"),
							Value: to.Ptr("mValue1"),
						}},
					Properties: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
						{
							Key:   to.Ptr("Creation-Time"),
							Value: to.Ptr("Wed, 07 Jun 2023 05:23:29 GMT"),
						},
						{
							Key:   to.Ptr("Last-Modified"),
							Value: to.Ptr("Wed, 07 Jun 2023 05:23:29 GMT"),
						},
						{
							Key:   to.Ptr("Etag"),
							Value: to.Ptr("0x8DB67175454D36D"),
						},
						{
							Key:   to.Ptr("Content-Length"),
							Value: to.Ptr("38619"),
						},
						{
							Key:   to.Ptr("Content-Type"),
							Value: to.Ptr("text/xml"),
						},
						{
							Key:   to.Ptr("Content-Encoding"),
							Value: to.Ptr(""),
						},
						{
							Key:   to.Ptr("Content-Language"),
							Value: to.Ptr(""),
						},
						{
							Key:   to.Ptr("Content-CRC64"),
							Value: to.Ptr(""),
						},
						{
							Key:   to.Ptr("Content-MD5"),
							Value: to.Ptr("njr6iDrmU9+FC89WMK22EA=="),
						},
						{
							Key:   to.Ptr("Cache-Control"),
							Value: to.Ptr(""),
						},
						{
							Key:   to.Ptr("Content-Disposition"),
							Value: to.Ptr(""),
						},
						{
							Key:   to.Ptr("BlobType"),
							Value: to.Ptr("BlockBlob"),
						},
						{
							Key:   to.Ptr("AccessTier"),
							Value: to.Ptr("Hot"),
						},
						{
							Key:   to.Ptr("AccessTierInferred"),
							Value: to.Ptr("true"),
						},
						{
							Key:   to.Ptr("LeaseStatus"),
							Value: to.Ptr("unlocked"),
						},
						{
							Key:   to.Ptr("LeaseState"),
							Value: to.Ptr("available"),
						},
						{
							Key:   to.Ptr("ServerEncrypted"),
							Value: to.Ptr("true"),
						},
						{
							Key:   to.Ptr("TagCount"),
							Value: to.Ptr("1"),
						}},
					Tags: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
						{
							Key:   to.Ptr("tKey1"),
							Value: to.Ptr("tValue1"),
						}},
				},
				{
					Name: to.Ptr("folder2/file1.txt"),
					Metadata: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
						{
							Key:   to.Ptr("mKey2"),
							Value: to.Ptr("mValue2"),
						}},
					Properties: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
						{
							Key:   to.Ptr("Creation-Time"),
							Value: to.Ptr("Wed, 06 Jun 2023 05:23:29 GMT"),
						},
						{
							Key:   to.Ptr("Last-Modified"),
							Value: to.Ptr("Wed, 06 Jun 2023 05:23:29 GMT"),
						},
						{
							Key:   to.Ptr("Etag"),
							Value: to.Ptr("0x6FB67175454D36D"),
						}},
					Tags: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
						{
							Key:   to.Ptr("tKey2"),
							Value: to.Ptr("tValue2"),
						}},
				}},
			Container: &armstorageactions.StorageTaskPreviewContainerProperties{
				Name: to.Ptr("firstContainer"),
				Metadata: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
					{
						Key:   to.Ptr("mContainerKey1"),
						Value: to.Ptr("mContainerValue1"),
					}},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageTaskPreviewAction = armstorageactions.StorageTaskPreviewAction{
	// 	Properties: &armstorageactions.StorageTaskPreviewActionProperties{
	// 		Action: &armstorageactions.StorageTaskPreviewActionCondition{
	// 			ElseBlockExists: to.Ptr(true),
	// 			If: &armstorageactions.StorageTaskPreviewActionIfCondition{
	// 				Condition: to.Ptr("[[equals(AccessTier, 'Hot')]]"),
	// 			},
	// 		},
	// 		Blobs: []*armstorageactions.StorageTaskPreviewBlobProperties{
	// 			{
	// 				Name: to.Ptr("folder1/file1.txt"),
	// 				MatchedBlock: to.Ptr(armstorageactions.MatchedBlockNameIf),
	// 				Metadata: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
	// 					{
	// 						Key: to.Ptr("mKey1"),
	// 						Value: to.Ptr("mValue1"),
	// 				}},
	// 				Properties: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
	// 					{
	// 						Key: to.Ptr("Creation-Time"),
	// 						Value: to.Ptr("Wed, 07 Jun 2023 05:23:29 GMT"),
	// 					},
	// 					{
	// 						Key: to.Ptr("Last-Modified"),
	// 						Value: to.Ptr("Wed, 07 Jun 2023 05:23:29 GMT"),
	// 					},
	// 					{
	// 						Key: to.Ptr("Etag"),
	// 						Value: to.Ptr("0x8DB67175454D36D"),
	// 					},
	// 					{
	// 						Key: to.Ptr("Content-Length"),
	// 						Value: to.Ptr("38619"),
	// 					},
	// 					{
	// 						Key: to.Ptr("Content-Type"),
	// 						Value: to.Ptr("text/xml"),
	// 					},
	// 					{
	// 						Key: to.Ptr("Content-Encoding"),
	// 						Value: to.Ptr(""),
	// 					},
	// 					{
	// 						Key: to.Ptr("Content-Language"),
	// 						Value: to.Ptr(""),
	// 					},
	// 					{
	// 						Key: to.Ptr("Content-CRC64"),
	// 						Value: to.Ptr(""),
	// 					},
	// 					{
	// 						Key: to.Ptr("Content-MD5"),
	// 						Value: to.Ptr("njr6iDrmU9+FC89WMK22EA=="),
	// 					},
	// 					{
	// 						Key: to.Ptr("Cache-Control"),
	// 						Value: to.Ptr(""),
	// 					},
	// 					{
	// 						Key: to.Ptr("Content-Disposition"),
	// 						Value: to.Ptr(""),
	// 					},
	// 					{
	// 						Key: to.Ptr("BlobType"),
	// 						Value: to.Ptr("BlockBlob"),
	// 					},
	// 					{
	// 						Key: to.Ptr("AccessTier"),
	// 						Value: to.Ptr("Hot"),
	// 					},
	// 					{
	// 						Key: to.Ptr("AccessTierInferred"),
	// 						Value: to.Ptr("true"),
	// 					},
	// 					{
	// 						Key: to.Ptr("LeaseStatus"),
	// 						Value: to.Ptr("unlocked"),
	// 					},
	// 					{
	// 						Key: to.Ptr("LeaseState"),
	// 						Value: to.Ptr("available"),
	// 					},
	// 					{
	// 						Key: to.Ptr("ServerEncrypted"),
	// 						Value: to.Ptr("true"),
	// 					},
	// 					{
	// 						Key: to.Ptr("TagCount"),
	// 						Value: to.Ptr("1"),
	// 				}},
	// 				Tags: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
	// 					{
	// 						Key: to.Ptr("tKey1"),
	// 						Value: to.Ptr("tValue1"),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("folder2/file1.txt"),
	// 				MatchedBlock: to.Ptr(armstorageactions.MatchedBlockNameElse),
	// 				Metadata: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
	// 					{
	// 						Key: to.Ptr("mKey2"),
	// 						Value: to.Ptr("mValue2"),
	// 				}},
	// 				Properties: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
	// 					{
	// 						Key: to.Ptr("Creation-Time"),
	// 						Value: to.Ptr("Wed, 06 Jun 2023 05:23:29 GMT"),
	// 					},
	// 					{
	// 						Key: to.Ptr("Last-Modified"),
	// 						Value: to.Ptr("Wed, 06 Jun 2023 05:23:29 GMT"),
	// 					},
	// 					{
	// 						Key: to.Ptr("Etag"),
	// 						Value: to.Ptr("0x6FB67175454D36D"),
	// 				}},
	// 				Tags: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
	// 					{
	// 						Key: to.Ptr("tKey2"),
	// 						Value: to.Ptr("tValue2"),
	// 				}},
	// 		}},
	// 		Container: &armstorageactions.StorageTaskPreviewContainerProperties{
	// 			Name: to.Ptr("firstContainer"),
	// 			Metadata: []*armstorageactions.StorageTaskPreviewKeyValueProperties{
	// 				{
	// 					Key: to.Ptr("mContainerKey1"),
	// 					Value: to.Ptr("mContainerValue1"),
	// 			}},
	// 		},
	// 	},
	// }
}
