package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4e9df3afd38a1cfa00a5d49419dce51bd014601f/specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/BlobRangesRestore.json
func ExampleAccountsClient_BeginRestoreBlobRanges() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginRestoreBlobRanges(ctx, "res9101", "sto4445", armstorage.BlobRestoreParameters{
		BlobRanges: []*armstorage.BlobRestoreRange{
			{
				EndRange:   to.Ptr("container/blobpath2"),
				StartRange: to.Ptr("container/blobpath1"),
			},
			{
				EndRange:   to.Ptr(""),
				StartRange: to.Ptr("container2/blobpath3"),
			}},
		TimeToRestore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-20T15:30:00.000Z"); return t }()),
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
	// res.BlobRestoreStatus = armstorage.BlobRestoreStatus{
	// 	Parameters: &armstorage.BlobRestoreParameters{
	// 		BlobRanges: []*armstorage.BlobRestoreRange{
	// 			{
	// 				EndRange: to.Ptr("container/blobpath2"),
	// 				StartRange: to.Ptr("container/blobpath1"),
	// 			},
	// 			{
	// 				EndRange: to.Ptr(""),
	// 				StartRange: to.Ptr("container2/blobpath3"),
	// 		}},
	// 		TimeToRestore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-20T15:30:00.000Z"); return t}()),
	// 	},
	// 	RestoreID: to.Ptr("{restore_id}"),
	// 	Status: to.Ptr(armstorage.BlobRestoreProgressStatus("Succeeded")),
	// }
}
