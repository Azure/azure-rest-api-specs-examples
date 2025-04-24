package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/29bc17223449e7a865a6f38552eacfd213d812fd/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-02-02-preview/examples/SnapshotsDelete.json
func ExampleSnapshotsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSnapshotsClient().Delete(ctx, "rg1", "snapshot1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
