package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge/v2"
)

// Generated from example definition: 2023-12-01/ListMonitoringConfig.json
func ExampleMonitoringConfigClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitoringConfigClient().NewListPager("testedgedevice", "testrole", "GroupForEdgeAutomation", nil)
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
		// page = armdataboxedge.MonitoringConfigClientListResponse{
		// 	MonitoringMetricConfigurationList: armdataboxedge.MonitoringMetricConfigurationList{
		// 		Value: []*armdataboxedge.MonitoringMetricConfiguration{
		// 			{
		// 				ID: to.Ptr("foo/bar/id"),
		// 				Properties: &armdataboxedge.MonitoringMetricConfigurationProperties{
		// 					MetricConfigurations: []*armdataboxedge.MetricConfiguration{
		// 						{
		// 							CounterSets: []*armdataboxedge.MetricCounterSet{
		// 								{
		// 									Counters: []*armdataboxedge.MetricCounter{
		// 										{
		// 											Name: to.Ptr("test"),
		// 										},
		// 									},
		// 								},
		// 							},
		// 							MdmAccount: to.Ptr("test"),
		// 							MetricNameSpace: to.Ptr("test"),
		// 							ResourceID: to.Ptr("test"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
