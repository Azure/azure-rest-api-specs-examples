package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: 2024-04-01/CreateFileForSubscription.json
func ExampleFilesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("132d901f-189d-4381-9214-fe68e27e05a1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFilesClient().Create(ctx, "testworkspace", "test.txt", armsupport.FileDetails{
		Properties: &armsupport.FileDetailsProperties{
			ChunkSize:      to.Ptr[int32](41423),
			FileSize:       to.Ptr[int32](41423),
			NumberOfChunks: to.Ptr[int32](1),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
