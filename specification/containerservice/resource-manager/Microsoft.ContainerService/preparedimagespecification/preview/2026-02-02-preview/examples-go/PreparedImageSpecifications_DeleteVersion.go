package armcontainerservicepreparedimgspec_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicepreparedimgspec/armcontainerservicepreparedimgspec"
)

// Generated from example definition: 2026-02-02-preview/PreparedImageSpecifications_DeleteVersion.json
func ExamplePreparedImageSpecificationsClient_BeginDeleteVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicepreparedimgspec.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPreparedImageSpecificationsClient().BeginDeleteVersion(ctx, "rg1", "my-prepared-image-specification", "20250101-abcd1234", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
