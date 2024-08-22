package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedEnvironmentsStorages_Delete.json
func ExampleManagedEnvironmentsStoragesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewManagedEnvironmentsStoragesClient().Delete(ctx, "examplerg", "managedEnv", "jlaw-demo1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
