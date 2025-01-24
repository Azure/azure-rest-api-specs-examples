package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5759c77eee2d57bdb9e47aa1805d0ffb61704f2d/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/EventHubs/EHEventHubCreate.json
func ExampleEventHubsClient_CreateOrUpdate_ehEventHubCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventHubsClient().CreateOrUpdate(ctx, "Default-NotificationHubs-AustraliaEast", "sdk-Namespace-5357", "sdk-EventHub-6547", armeventhub.Eventhub{
		Properties: &armeventhub.Properties{
			CaptureDescription: &armeventhub.CaptureDescription{
				Destination: &armeventhub.Destination{
					Name: to.Ptr("EventHubArchive.AzureBlockBlob"),
					Identity: &armeventhub.CaptureIdentity{
						Type:                 to.Ptr(armeventhub.CaptureIdentityTypeUserAssigned),
						UserAssignedIdentity: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2"),
					},
					Properties: &armeventhub.DestinationProperties{
						ArchiveNameFormat:        to.Ptr("{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}"),
						BlobContainer:            to.Ptr("container"),
						StorageAccountResourceID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage"),
					},
				},
				Enabled:           to.Ptr(true),
				Encoding:          to.Ptr(armeventhub.EncodingCaptureDescriptionAvro),
				IntervalInSeconds: to.Ptr[int32](120),
				SizeLimitInBytes:  to.Ptr[int32](10485763),
			},
			PartitionCount: to.Ptr[int64](4),
			Status:         to.Ptr(armeventhub.EntityStatusActive),
			UserMetadata:   to.Ptr("key"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Eventhub = armeventhub.Eventhub{
	// 	Name: to.Ptr("sdk-EventHub-10"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/EventHubs"),
	// 	ID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-NotificationHubs-AustraliaEast/providers/Microsoft.EventHub/namespaces/sdk-Namespace-716/eventhubs/sdk-EventHub-10"),
	// 	Properties: &armeventhub.Properties{
	// 		CaptureDescription: &armeventhub.CaptureDescription{
	// 			Destination: &armeventhub.Destination{
	// 				Name: to.Ptr("EventHubArchive.AzureBlockBlob"),
	// 				Identity: &armeventhub.CaptureIdentity{
	// 					Type: to.Ptr(armeventhub.CaptureIdentityTypeUserAssigned),
	// 					UserAssignedIdentity: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2"),
	// 				},
	// 				Properties: &armeventhub.DestinationProperties{
	// 					ArchiveNameFormat: to.Ptr("{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}"),
	// 					BlobContainer: to.Ptr("container"),
	// 					StorageAccountResourceID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage"),
	// 				},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			Encoding: to.Ptr(armeventhub.EncodingCaptureDescriptionAvro),
	// 			IntervalInSeconds: to.Ptr[int32](120),
	// 			SizeLimitInBytes: to.Ptr[int32](10485763),
	// 		},
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T02:45:55.877Z"); return t}()),
	// 		Identifier: to.Ptr("identifierIDGUID"),
	// 		MessageRetentionInDays: to.Ptr[int64](7),
	// 		MessageTimestampDescription: &armeventhub.MessageTimestampDescription{
	// 			TimestampType: to.Ptr(armeventhub.TimestampTypeLogAppend),
	// 		},
	// 		PartitionCount: to.Ptr[int64](4),
	// 		PartitionIDs: []*string{
	// 			to.Ptr("0"),
	// 			to.Ptr("1"),
	// 			to.Ptr("2"),
	// 			to.Ptr("3")},
	// 			RetentionDescription: &armeventhub.RetentionDescription{
	// 				CleanupPolicy: to.Ptr(armeventhub.CleanupPolicyRetentionDescriptionDelete),
	// 				RetentionTimeInHours: to.Ptr[int64](168),
	// 			},
	// 			Status: to.Ptr(armeventhub.EntityStatusActive),
	// 			UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T02:46:05.877Z"); return t}()),
	// 			UserMetadata: to.Ptr("key"),
	// 		},
	// 	}
}
