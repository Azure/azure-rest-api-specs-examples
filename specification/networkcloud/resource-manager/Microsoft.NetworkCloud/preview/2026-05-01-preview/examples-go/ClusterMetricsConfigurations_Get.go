package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: 2026-05-01-preview/ClusterMetricsConfigurations_Get.json
func ExampleMetricsConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("123e4567-e89b-12d3-a456-426655440000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricsConfigurationsClient().Get(ctx, "resourceGroupName", "clusterName", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetworkcloud.MetricsConfigurationsClientGetResponse{
	// 	ClusterMetricsConfiguration: armnetworkcloud.ClusterMetricsConfiguration{
	// 		ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 			Type: to.Ptr("CustomLocation"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName/metricsConfigurations/default"),
	// 		Location: to.Ptr("location"),
	// 		Name: to.Ptr("default"),
	// 		Properties: &armnetworkcloud.ClusterMetricsConfigurationProperties{
	// 			CollectionInterval: to.Ptr[int64](15),
	// 			DetailedStatus: to.Ptr(armnetworkcloud.ClusterMetricsConfigurationDetailedStatusApplied),
	// 			DetailedStatusMessage: to.Ptr("Metrics configuration has been successfully applied to all resources."),
	// 			DisabledMetrics: []*string{
	// 				to.Ptr("metric3"),
	// 				to.Ptr("metric4"),
	// 			},
	// 			EnabledMetrics: []*string{
	// 				to.Ptr("metric1"),
	// 				to.Ptr("metric2"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetworkcloud.ClusterMetricsConfigurationProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armnetworkcloud.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 			CreatedBy: to.Ptr("identityA"),
	// 			CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("identityB"),
	// 			LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("myvalue1"),
	// 			"key2": to.Ptr("myvalue2"),
	// 		},
	// 		Type: to.Ptr("Microsoft.NetworkCloud/clusters/metricsConfigurations"),
	// 	},
	// }
}
