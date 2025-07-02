package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86c6306649b02e542117adb46c61e8019dbd78e9/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/BlobContainersList.json
func ExampleBlobContainersClient_NewListPager_listContainers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBlobContainersClient().NewListPager("res9290", "sto1590", &armstorage.BlobContainersClientListOptions{Maxpagesize: nil,
		Filter:  nil,
		Include: nil,
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
		// page.ListContainerItems = armstorage.ListContainerItems{
		// 	Value: []*armstorage.ListContainerItem{
		// 		{
		// 			Name: to.Ptr("container1644"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/containers"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/blobServices/default/containers/container1644"),
		// 			Etag: to.Ptr("\"0x8D589847D51C7DE\""),
		// 			Properties: &armstorage.ContainerProperties{
		// 				HasImmutabilityPolicy: to.Ptr(false),
		// 				HasLegalHold: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T08:20:47.000Z"); return t}()),
		// 				LeaseState: to.Ptr(armstorage.LeaseStateAvailable),
		// 				LeaseStatus: to.Ptr(armstorage.LeaseStatusUnlocked),
		// 				PublicAccess: to.Ptr(armstorage.PublicAccessContainer),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("container4052"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/containers"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/blobServices/default/containers/container4052"),
		// 			Etag: to.Ptr("\"0x8D589847DAB5AF9\""),
		// 			Properties: &armstorage.ContainerProperties{
		// 				HasImmutabilityPolicy: to.Ptr(false),
		// 				HasLegalHold: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T08:20:47.000Z"); return t}()),
		// 				LeaseState: to.Ptr(armstorage.LeaseStateAvailable),
		// 				LeaseStatus: to.Ptr(armstorage.LeaseStatusUnlocked),
		// 				PublicAccess: to.Ptr(armstorage.PublicAccessNone),
		// 			},
		// 	}},
		// }
	}
}
