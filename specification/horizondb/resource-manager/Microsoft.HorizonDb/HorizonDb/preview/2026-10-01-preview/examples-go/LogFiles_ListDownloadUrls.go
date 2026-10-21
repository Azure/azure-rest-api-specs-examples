package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/LogFiles_ListDownloadUrls.json
func ExampleLogFilesClient_ListDownloadUrls() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogFilesClient().ListDownloadUrls(ctx, "example-resource-group", "example-cluster", "22222222-2222-2222-2222-222222222222", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.LogFilesClientListDownloadUrlsResponse{
	// 	LogFileDownloadURL: armhorizondb.LogFileDownloadURL{
	// 		DownloadURI: to.Ptr("https://example.blob.core.windows.net/logs/22222222-2222-2222-2222-222222222222.log?se=2026-09-11T13%3A00%3A00Z&sig=redacted"),
	// 		ExpiresAtUTC: to.Ptr(time.Date(2026, time.September, 11, 13, 0, 0, 0, time.UTC)),
	// 	},
	// }
}
