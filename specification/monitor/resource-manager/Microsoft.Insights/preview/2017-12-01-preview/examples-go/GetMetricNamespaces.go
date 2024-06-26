package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2017-12-01-preview/examples/GetMetricNamespaces.json
func ExampleMetricNamespacesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetricNamespacesClient().NewListPager("subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill", &armmonitor.MetricNamespacesClientListOptions{StartTime: to.Ptr("2020-08-31T15:53:00Z")})
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
		// page.MetricNamespaceCollection = armmonitor.MetricNamespaceCollection{
		// 	Value: []*armmonitor.MetricNamespace{
		// 		{
		// 			Name: to.Ptr("Azure.ApplicationInsights"),
		// 			Type: to.Ptr("Microsoft.Insights/metricNamespaces"),
		// 			Classification: to.Ptr(armmonitor.NamespaceClassificationCustom),
		// 			ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricNamespaces/Azure.ApplicationInsights"),
		// 			Properties: &armmonitor.MetricNamespaceName{
		// 				MetricNamespaceName: to.Ptr("Azure.ApplicationInsights"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.insights-components"),
		// 			Type: to.Ptr("Microsoft.Insights/metricNamespaces"),
		// 			Classification: to.Ptr(armmonitor.NamespaceClassificationPlatform),
		// 			ID: to.Ptr("/subscriptions/182c901a-129a-4f5d-86e4-cc6b294590a2/resourceGroups/hyr-log/providers/microsoft.insights/components/f1-bill/providers/microsoft.insights/metricNamespaces/microsoft.insights-components"),
		// 			Properties: &armmonitor.MetricNamespaceName{
		// 				MetricNamespaceName: to.Ptr("microsoft.insights/components"),
		// 			},
		// 	}},
		// }
	}
}
