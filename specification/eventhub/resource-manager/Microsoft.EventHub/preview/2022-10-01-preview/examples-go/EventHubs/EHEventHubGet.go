package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/EventHubs/EHEventHubGet.json
func ExampleEventHubsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventHubsClient().Get(ctx, "Default-NotificationHubs-AustraliaEast", "sdk-Namespace-716", "sdk-EventHub-10", nil)
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
	// 		MessageRetentionInDays: to.Ptr[int64](4),
	// 		PartitionCount: to.Ptr[int64](4),
	// 		PartitionIDs: []*string{
	// 			to.Ptr("0"),
	// 			to.Ptr("1"),
	// 			to.Ptr("2"),
	// 			to.Ptr("3")},
	// 			RetentionDescription: &armeventhub.RetentionDescription{
	// 				CleanupPolicy: to.Ptr(armeventhub.CleanupPolicyRetentionDescriptionCompaction),
	// 				RetentionTimeInHours: to.Ptr[int64](96),
	// 				TombstoneRetentionTimeInHours: to.Ptr[int32](1),
	// 			},
	// 			Status: to.Ptr(armeventhub.EntityStatusActive),
	// 			UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T02:46:05.877Z"); return t}()),
	// 		},
	// 	}
}
