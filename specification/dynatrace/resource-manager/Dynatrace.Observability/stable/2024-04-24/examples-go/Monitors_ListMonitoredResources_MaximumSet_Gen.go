package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace/v2"
)

// Generated from example definition: 2024-04-24/Monitors_ListMonitoredResources_MaximumSet_Gen.json
func ExampleMonitorsClient_NewListMonitoredResourcesPager_monitorsListMonitoredResourcesMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListMonitoredResourcesPager("myResourceGroup", "myMonitor", &armdynatrace.MonitorsClientListMonitoredResourcesOptions{
		Request: &armdynatrace.LogStatusRequest{
			MonitoredResourceIDs: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor/listMonitoredResources"),
			},
		}})
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
		// page = armdynatrace.MonitorsClientListMonitoredResourcesResponse{
		// 	MonitoredResourceListResponse: armdynatrace.MonitoredResourceListResponse{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/providers/Dynatrace.Observability/monitors?api-version=2024-04-24&$skiptoken=abc123"),
		// 		Value: []*armdynatrace.MonitoredResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor/listMonitoredResources"),
		// 				ReasonForLogsStatus: to.Ptr("CapturedByRules"),
		// 				ReasonForMetricsStatus: to.Ptr("CapturedByRules"),
		// 				SendingLogs: to.Ptr(armdynatrace.SendingLogsStatusEnabled),
		// 				SendingMetrics: to.Ptr(armdynatrace.SendingMetricsStatusEnabled),
		// 			},
		// 		},
		// 	},
		// }
	}
}
