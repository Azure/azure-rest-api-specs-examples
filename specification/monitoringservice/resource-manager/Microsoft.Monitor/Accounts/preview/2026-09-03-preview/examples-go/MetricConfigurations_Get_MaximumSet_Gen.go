package armmonitorworkspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitorworkspaces"
)

// Generated from example definition: 2026-09-03-preview/MetricConfigurations_Get_MaximumSet_Gen.json
func ExampleMetricConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitorworkspaces.NewClientFactory("703362b3-f278-4e4b-9179-c76eaf41ffc2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricConfigurationsClient().Get(ctx, "rgazuremonitorworkspace", "myAzureMonitorWorkspace", "default", "customdefault", "aggregatedRequestCount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitorworkspaces.MetricConfigurationsClientGetResponse{
	// 	MetricConfigurationResource: armmonitorworkspaces.MetricConfigurationResource{
	// 		Properties: &armmonitorworkspaces.MetricConfigurationProperties{
	// 			ProvisioningState: to.Ptr(armmonitorworkspaces.ResourceProvisioningStateSucceeded),
	// 			Namespace: to.Ptr("customdefault"),
	// 			MetricName: to.Ptr("aggregatedRequestCount"),
	// 			MetricType: to.Ptr(armmonitorworkspaces.MetricConfigurationTypeAggregated),
	// 			Dimensions: []*string{
	// 				to.Ptr("microsoft.resourceId"),
	// 			},
	// 			AggregationConfigurations: []*armmonitorworkspaces.MetricAggregationConfiguration{
	// 				{
	// 					StoreAggregatedData: to.Ptr(true),
	// 					Dimensions: []*string{
	// 						to.Ptr("microsoft.resourceId"),
	// 					},
	// 					AggregationFunctions: &armmonitorworkspaces.MetricAggregationFunctions{
	// 						EnableMinMax: to.Ptr(true),
	// 						EnablePercentiles: to.Ptr(false),
	// 					},
	// 				},
	// 			},
	// 			SourceMetricResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/rgazuremonitorworkspace/providers/Microsoft.Monitor/accounts/myAzureMonitorWorkspace/metricsContainers/default/namespaces/customdefault/metrics/requestCount"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/rgazuremonitorworkspace/providers/Microsoft.Monitor/accounts/myAzureMonitorWorkspace/metricsContainers/default/namespaces/customdefault/metrics/aggregatedRequestCount"),
	// 		Name: to.Ptr("aggregatedRequestCount"),
	// 		Type: to.Ptr("Microsoft.Monitor/accounts/metricsContainers/namespaces/metrics"),
	// 		SystemData: &armmonitorworkspaces.SystemData{
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.September, 1, 12, 34, 56, 123456700, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armmonitorworkspaces.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.September, 2, 12, 34, 56, 123456700, time.UTC)),
	// 		},
	// 	},
	// }
}
