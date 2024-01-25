package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8b5618d760532b69d6f7434f57172ea52e109c79/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/MaintenanceConfigurationsDelete.json
func ExampleMaintenanceConfigurationsClient_Delete_deleteMaintenanceConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMaintenanceConfigurationsClient().Delete(ctx, "rg1", "clustername1", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
