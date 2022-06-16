package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/EventHubs/EHEventHubCreate.json
func ExampleEventHubsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventhub.NewEventHubsClient("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"Default-NotificationHubs-AustraliaEast",
		"sdk-Namespace-5357",
		"sdk-EventHub-6547",
		armeventhub.Eventhub{
			Properties: &armeventhub.Properties{
				CaptureDescription: &armeventhub.CaptureDescription{
					Destination: &armeventhub.Destination{
						Name: to.Ptr("EventHubArchive.AzureBlockBlob"),
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
				MessageRetentionInDays: to.Ptr[int64](4),
				PartitionCount:         to.Ptr[int64](4),
				Status:                 to.Ptr(armeventhub.EntityStatusActive),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
