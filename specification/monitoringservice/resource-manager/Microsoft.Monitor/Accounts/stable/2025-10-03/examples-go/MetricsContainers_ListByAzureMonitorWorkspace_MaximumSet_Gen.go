package armmonitorworkspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitorworkspaces"
)

// Generated from example definition: 2025-10-03/MetricsContainers_ListByAzureMonitorWorkspace_MaximumSet_Gen.json
func ExampleMetricsContainersClient_NewListByAzureMonitorWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitorworkspaces.NewClientFactory("703362b3-f278-4e4b-9179-c76eaf41ffc2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetricsContainersClient().NewListByAzureMonitorWorkspacePager("rgazuremonitorworkspace", "myAzureMonitorWorkspace", nil)
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
		// page = armmonitorworkspaces.MetricsContainersClientListByAzureMonitorWorkspaceResponse{
		// 	MetricsContainerResourceListResult: armmonitorworkspaces.MetricsContainerResourceListResult{
		// 		Value: []*armmonitorworkspaces.MetricsContainerResource{
		// 			{
		// 				Properties: &armmonitorworkspaces.MetricsContainer{
		// 					ProvisioningState: to.Ptr(armmonitorworkspaces.ResourceProvisioningStateSucceeded),
		// 					Version: to.Ptr("1.0"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/rgazuremonitorworkspace/providers/Microsoft.Monitor/accounts/myAzureMonitorWorkspace/metricsContainers/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.Monitor/accounts/metricsContainers"),
		// 				SystemData: &armmonitorworkspaces.SystemData{
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.1234567Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
