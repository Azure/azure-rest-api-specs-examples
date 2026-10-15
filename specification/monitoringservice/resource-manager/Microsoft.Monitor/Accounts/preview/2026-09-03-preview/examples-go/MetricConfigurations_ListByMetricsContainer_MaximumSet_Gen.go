package armmonitorworkspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitorworkspaces"
)

// Generated from example definition: 2026-09-03-preview/MetricConfigurations_ListByMetricsContainer_MaximumSet_Gen.json
func ExampleMetricConfigurationsClient_NewListByMetricsContainerPager_metricConfigurationsListByMetricsContainerMaximumSetGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitorworkspaces.NewClientFactory("703362b3-f278-4e4b-9179-c76eaf41ffc2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetricConfigurationsClient().NewListByMetricsContainerPager("rgazuremonitorworkspace", "myAzureMonitorWorkspace", "default", &armmonitorworkspaces.MetricConfigurationsClientListByMetricsContainerOptions{
		Filter: to.Ptr("Properties.sourceMetricResourceId eq null")})
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
		// page = armmonitorworkspaces.MetricConfigurationsClientListByMetricsContainerResponse{
		// 	PagedMetricConfigurationResource: armmonitorworkspaces.PagedMetricConfigurationResource{
		// 		Value: []*armmonitorworkspaces.MetricConfigurationResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/rgazuremonitorworkspace/providers/Microsoft.Monitor/accounts/myAzureMonitorWorkspace/metricsContainers/default/namespaces/customdefault/metrics/requestCount"),
		// 				Name: to.Ptr("requestCount"),
		// 				Type: to.Ptr("Microsoft.Monitor/accounts/metricsContainers/namespaces/metrics"),
		// 				SystemData: &armmonitorworkspaces.SystemData{
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2026, time.September, 1, 12, 34, 56, 123456700, time.UTC)),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2026, time.September, 2, 12, 34, 56, 123456700, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
