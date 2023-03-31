package armcustomerlockbox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/customerlockbox/resource-manager/Microsoft.CustomerLockbox/preview/2018-02-28-preview/examples/EnableLockbox.json
func ExamplePostClient_EnableLockbox() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerlockbox.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPostClient().EnableLockbox(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
