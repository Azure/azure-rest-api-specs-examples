package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streaming-locators-list-paths-streaming-and-download.json
func ExampleStreamingLocatorsClient_ListPaths_listPathsWhichHasStreamingPathsAndDownloadPaths() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingLocatorsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListPaths(ctx, "contoso", "contosomedia", "clearStreamingLocator", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
