package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9cc7633f842575274f715cc02e37c5769ac2742d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/MaintenanceConfigurationsDelete_MaintenanceWindow.json
func ExampleMaintenanceConfigurationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMaintenanceConfigurationsClient().Delete(ctx, "rg1", "clustername1", "aksManagedNodeOSUpgradeSchedule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
