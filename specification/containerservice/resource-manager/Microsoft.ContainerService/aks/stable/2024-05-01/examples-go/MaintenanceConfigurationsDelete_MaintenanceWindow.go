package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8cbd903b0d037c427def0e8fbf167b26e87bd1a9/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2024-05-01/examples/MaintenanceConfigurationsDelete_MaintenanceWindow.json
func ExampleMaintenanceConfigurationsClient_Delete_deleteMaintenanceConfigurationForNodeOsUpgrade() {
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
