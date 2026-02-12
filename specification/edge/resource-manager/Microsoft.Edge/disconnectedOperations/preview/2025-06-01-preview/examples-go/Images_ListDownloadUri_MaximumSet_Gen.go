package armdisconnectedoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/disconnectedoperations/armdisconnectedoperations"
)

// Generated from example definition: 2025-06-01-preview/Images_ListDownloadUri_MaximumSet_Gen.json
func ExampleImagesClient_ListDownloadURI() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdisconnectedoperations.NewClientFactory("51DB5DE7-A66C-4789-BFFF-9F75C95A0201", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewImagesClient().ListDownloadURI(ctx, "rgdisconnectedOperations", "g_-5-160", "1Q6lGV4V65j-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdisconnectedoperations.ImagesClientListDownloadURIResponse{
	// 	ImageDownloadResult: &armdisconnectedoperations.ImageDownloadResult{
	// 		ReleaseVersion: to.Ptr("vbccjpcosofti"),
	// 		ReleaseDisplayName: to.Ptr("thttwzm"),
	// 		ReleaseNotes: to.Ptr("jswqauqfnwxk"),
	// 		ReleaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.DateOnly, "2024-12-05"); return t}()),
	// 		CompatibleVersions: []*string{
	// 			to.Ptr("czxghshqcn"),
	// 		},
	// 		TransactionID: to.Ptr("337b8e3a-dd7c-4872-a270-5d57632a8aea"),
	// 		DownloadLink: to.Ptr("https://microsoft.com/akmzb"),
	// 		ReleaseType: to.Ptr(armdisconnectedoperations.ReleaseTypeInstall),
	// 		LinkExpiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-05T20:03:13.362Z"); return t}()),
	// 	},
	// }
}
