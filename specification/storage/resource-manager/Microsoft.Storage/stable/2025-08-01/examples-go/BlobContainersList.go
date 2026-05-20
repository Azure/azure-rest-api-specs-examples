package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2025-08-01/BlobContainersList.json
func ExampleBlobContainersClient_NewListPager_listContainers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBlobContainersClient().NewListPager("res9290", "sto1590", nil)
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
		// page = armstorage.BlobContainersClientListResponse{
		// 	ListContainerItems: armstorage.ListContainerItems{
		// 		NextLink: to.Ptr("https://sto1590endpoint/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/blobServices/default/containers?api-version=2022-09-01&$maxpagesize=2&$skipToken=/sto1590/container5103"),
		// 		Value: []*armstorage.ListContainerItem{
		// 			{
		// 				Name: to.Ptr("container1644"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/containers"),
		// 				Etag: to.Ptr("\"0x8D589847D51C7DE\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/blobServices/default/containers/container1644"),
		// 				Properties: &armstorage.ContainerProperties{
		// 					HasImmutabilityPolicy: to.Ptr(false),
		// 					HasLegalHold: to.Ptr(false),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T08:20:47Z"); return t}()),
		// 					LeaseState: to.Ptr(armstorage.LeaseStateAvailable),
		// 					LeaseStatus: to.Ptr(armstorage.LeaseStatusUnlocked),
		// 					PublicAccess: to.Ptr(armstorage.PublicAccessContainer),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("container4052"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/containers"),
		// 				Etag: to.Ptr("\"0x8D589847DAB5AF9\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/blobServices/default/containers/container4052"),
		// 				Properties: &armstorage.ContainerProperties{
		// 					HasImmutabilityPolicy: to.Ptr(false),
		// 					HasLegalHold: to.Ptr(false),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T08:20:47Z"); return t}()),
		// 					LeaseState: to.Ptr(armstorage.LeaseStateAvailable),
		// 					LeaseStatus: to.Ptr(armstorage.LeaseStatusUnlocked),
		// 					PublicAccess: to.Ptr(armstorage.PublicAccessNone),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
