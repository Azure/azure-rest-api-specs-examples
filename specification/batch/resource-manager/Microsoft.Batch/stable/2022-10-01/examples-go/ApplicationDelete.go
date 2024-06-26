package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ef6f2f06858cdbec7684968e1a54c7610da97dbb/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/ApplicationDelete.json
func ExampleApplicationClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewApplicationClient().Delete(ctx, "default-azurebatch-japaneast", "sampleacct", "app1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
