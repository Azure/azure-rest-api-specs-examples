package armwindowsesu_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/windowsesu/resource-manager/Microsoft.WindowsESU/preview/2019-09-16-preview/examples/DeleteMultipleActivationKey.json
func ExampleMultipleActivationKeysClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armwindowsesu.NewMultipleActivationKeysClient("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"testgr1",
		"server08-key-2019",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
