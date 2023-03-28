package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2f2b633d940230cbbd5bcf1339a2e1c48674e4a2/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/EventHubs/EHEventHubListByNameSpace.json
func ExampleEventHubsClient_NewListByNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEventHubsClient().NewListByNamespacePager("Default-NotificationHubs-AustraliaEast", "sdk-Namespace-5357", &armeventhub.EventHubsClientListByNamespaceOptions{Skip: nil,
		Top: nil,
	})
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
		// page.ListResult = armeventhub.ListResult{
		// 	Value: []*armeventhub.Eventhub{
		// 		{
		// 			Name: to.Ptr("sdk-eventhub-10"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/EventHubs"),
		// 			ID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-NotificationHubs-AustraliaEast/providers/Microsoft.EventHub/namespaces/sdk-Namespace-716/eventhubs/sdk-eventhub-10"),
		// 			Properties: &armeventhub.Properties{
		// 				CaptureDescription: &armeventhub.CaptureDescription{
		// 					Destination: &armeventhub.Destination{
		// 						Name: to.Ptr("EventHubArchive.AzureBlockBlob"),
		// 						Properties: &armeventhub.DestinationProperties{
		// 							ArchiveNameFormat: to.Ptr("{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}"),
		// 							BlobContainer: to.Ptr("container"),
		// 							StorageAccountResourceID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage"),
		// 						},
		// 					},
		// 					Enabled: to.Ptr(true),
		// 					Encoding: to.Ptr(armeventhub.EncodingCaptureDescriptionAvro),
		// 					IntervalInSeconds: to.Ptr[int32](120),
		// 					SizeLimitInBytes: to.Ptr[int32](10485763),
		// 				},
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T02:45:55.877Z"); return t}()),
		// 				MessageRetentionInDays: to.Ptr[int64](4),
		// 				PartitionCount: to.Ptr[int64](4),
		// 				PartitionIDs: []*string{
		// 					to.Ptr("0"),
		// 					to.Ptr("1"),
		// 					to.Ptr("2"),
		// 					to.Ptr("3")},
		// 					Status: to.Ptr(armeventhub.EntityStatusActive),
		// 					UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T02:46:05.877Z"); return t}()),
		// 				},
		// 		}},
		// 	}
	}
}
