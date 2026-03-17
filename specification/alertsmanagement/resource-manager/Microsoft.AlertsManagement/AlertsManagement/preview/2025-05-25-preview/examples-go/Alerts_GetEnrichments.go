package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: 2025-05-25-preview/Alerts_GetEnrichments.json
func ExampleAlertsClient_NewGetEnrichmentsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("subscriptions/72fa99ef-9c84-4a7c-b343-ec62da107d81", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertsClient().NewGetEnrichmentsPager("66114d64-d9d9-478b-95c9-b789d6502101", nil)
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
		// page = armalertsmanagement.AlertsClientGetEnrichmentsResponse{
		// 	AlertEnrichmentsList: armalertsmanagement.AlertEnrichmentsList{
		// 		Value: []*armalertsmanagement.AlertEnrichmentResponse{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.AlertsManagement/alerts/enrichments"),
		// 				ID: to.Ptr("/subscriptions/9e261de7-c804-4b9d-9ebf-6f50fe350a9a/providers/Microsoft.AlertsManagement/alerts/66114d64-d9d9-478b-95c9-b789d6502101/enrichments/default"),
		// 				Properties: &armalertsmanagement.AlertEnrichmentProperties{
		// 					AlertID: to.Ptr("66114d64-d9d9-478b-95c9-b789d6502101"),
		// 					Enrichments: []armalertsmanagement.AlertEnrichmentItemClassification{
		// 						&armalertsmanagement.PrometheusInstantQuery{
		// 							Type: to.Ptr(armalertsmanagement.TypePrometheusInstantQuery),
		// 							Description: to.Ptr("Enrichment description"),
		// 							Datasources: []*string{
		// 								to.Ptr("/subscriptions/72fa99ef-9c84-4a7c-b343-ec62da107d81/resourceGroups/SyntheticRules/providers/microsoft.monitor/accounts/canaryamw"),
		// 							},
		// 							GrafanaExplorePath: to.Ptr("/explore?left=%7B%22datasource%22..."),
		// 							LinkToAPI: to.Ptr("https://test-3sxl.eastus.prometheus.monitor.azure.com/api/v1/query_range?..."),
		// 							Query: to.Ptr("sum by (cluster,container,replicaset,namespace)(label_replace( kube_pod_container_status_last_terminated_reason{reason='OOMKilled', cluster='cluster1', namespace='namespace1'}'}, 'replicaset', '$1', 'pod', '(.*)(-[a-z0-9]{5})$')) > 0"),
		// 							Status: to.Ptr(armalertsmanagement.StatusSucceeded),
		// 							Time: to.Ptr("2015-07-01T20:10:51.781Z"),
		// 							Title: to.Ptr("Number of OOM killed events by container"),
		// 						},
		// 						&armalertsmanagement.PrometheusRangeQuery{
		// 							Type: to.Ptr(armalertsmanagement.TypePrometheusRangeQuery),
		// 							Description: to.Ptr("Enrichment description"),
		// 							Datasources: []*string{
		// 								to.Ptr("/subscriptions/72fa99ef-9c84-4a7c-b343-ec62da107d81/resourceGroups/SyntheticRules/providers/microsoft.monitor/accounts/canaryamw"),
		// 							},
		// 							End: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-07-01T20:20:51.781Z"); return t}()),
		// 							GrafanaExplorePath: to.Ptr("/explore?left=%7B%22datasource%22..."),
		// 							LinkToAPI: to.Ptr("https://test-3sxl.eastus.prometheus.monitor.azure.com/api/v1/query_range?..."),
		// 							Query: to.Ptr("sum by (cluster,container,replicaset,namespace)(label_replace( kube_pod_container_status_last_terminated_reason{reason='OOMKilled', cluster='cluster1', namespace='namespace'}'}, 'replicaset', '$1', 'pod', '(.*)(-[a-z0-9]{5})$')) > 0"),
		// 							Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-07-01T20:10:51.781Z"); return t}()),
		// 							Status: to.Ptr(armalertsmanagement.StatusSucceeded),
		// 							Step: to.Ptr("PT15S"),
		// 							Title: to.Ptr("Number of OOM killed events by container"),
		// 						},
		// 						&armalertsmanagement.PrometheusRangeQuery{
		// 							Type: to.Ptr(armalertsmanagement.TypePrometheusRangeQuery),
		// 							Description: to.Ptr("Enrichment description"),
		// 							Datasources: []*string{
		// 								to.Ptr("/subscriptions/72fa99ef-9c84-4a7c-b343-ec62da107d81/resourceGroups/SyntheticRules/providers/microsoft.monitor/accounts/canaryamw"),
		// 							},
		// 							End: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-07-01T20:20:51.781Z"); return t}()),
		// 							ErrorMessage: to.Ptr("Calling Prometheus query API failed"),
		// 							GrafanaExplorePath: to.Ptr("/explore?left=%7B%22datasource%22..."),
		// 							LinkToAPI: to.Ptr("https://test-3sxl.eastus.prometheus.monitor.azure.com/api/v1/query_range?..."),
		// 							Query: to.Ptr("sum by (cluster,container,replicaset,namespace)(label_replace( kube_pod_container_status_last_terminated_reason{reason='OOMKilled', cluster='cluster1', namespace='namespace'}'}, 'replicaset', '$1', 'pod', '(.*)(-[a-z0-9]{5})$')) > 0"),
		// 							Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-07-01T20:10:51.781Z"); return t}()),
		// 							Status: to.Ptr(armalertsmanagement.StatusFailed),
		// 							Step: to.Ptr("PT15S"),
		// 							Title: to.Ptr("Number of OOM killed events by container"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
