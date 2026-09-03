package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-06-02-preview/AlertConfigurations_ListByManagedCluster.json
func ExampleAlertConfigurationsClient_NewListByManagedClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertConfigurationsClient().NewListByManagedClusterPager("rg1", "clustername1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcontainerservice.AlertConfigurationsClientListByManagedClusterResponse{
		// 	AlertConfigurationListResult: armcontainerservice.AlertConfigurationListResult{
		// 		Value: []*armcontainerservice.AlertConfiguration{
		// 			{
		// 				Name: to.Ptr("alertconfig1"),
		// 				Type: to.Ptr("Microsoft.ContainerService/managedClusters/alertConfigurations"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/alertConfigurations/alertconfig1"),
		// 				Properties: &armcontainerservice.AlertConfigurationProperties{
		// 					Mode: to.Ptr(armcontainerservice.AlertConfigurationModeManaged),
		// 					Notification: &armcontainerservice.AlertNotification{
		// 						ActionGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Insights/actionGroups/actiongroup1"),
		// 					},
		// 					ProvisioningState: to.Ptr(armcontainerservice.AlertConfigurationProvisioningStateSucceeded),
		// 				},
		// 				SystemData: &armcontainerservice.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("user@example.com"),
		// 					CreatedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("user@example.com"),
		// 					LastModifiedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
