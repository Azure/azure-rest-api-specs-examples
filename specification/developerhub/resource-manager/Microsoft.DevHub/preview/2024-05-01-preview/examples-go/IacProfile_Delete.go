package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/34f5146bc945549d087d38a8a593c0a5f475ad7f/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2024-05-01-preview/examples/IacProfile_Delete.json
func ExampleIacProfilesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewIacProfilesClient().Delete(ctx, "resourceGroup1", "iacprofile", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
