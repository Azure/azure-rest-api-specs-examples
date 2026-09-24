package armstorage_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-06-01/BlobRangesRestore.json
func ExampleAccountsClient_BeginRestoreBlobRanges() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
			},
		},
		TimeToRestore: to.Ptr(time.Date(2019, time.April, 20, 15, 30, 0, 0, time.UTC)),
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
	// res = armstorage.AccountsClientRestoreBlobRangesResponse{
	// 	BlobRestoreStatus: armstorage.BlobRestoreStatus{
	// 		Parameters: &armstorage.BlobRestoreParameters{
	// 			BlobRanges: []*armstorage.BlobRestoreRange{
	// 				{
	// 					EndRange: to.Ptr("container/blobpath2"),
	// 					StartRange: to.Ptr("container/blobpath1"),
	// 				},
	// 				{
	// 					EndRange: to.Ptr(""),
	// 					StartRange: to.Ptr("container2/blobpath3"),
	// 				},
	// 			},
	// 			TimeToRestore: to.Ptr(time.Date(2019, time.April, 20, 15, 30, 0, 0, time.UTC)),
	// 		},
	// 		RestoreID: to.Ptr("{restore_id}"),
	// 		Status: to.Ptr(armstorage.BlobRestoreProgressStatus("Succeeded")),
	// 	},
	// }
}
