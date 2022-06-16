package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/PutMonitoringConfig.json
func ExampleMonitoringConfigClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewMonitoringConfigClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<device-name>",
		"<role-name>",
		"<resource-group-name>",
		armdataboxedge.MonitoringMetricConfiguration{
			Properties: &armdataboxedge.MonitoringMetricConfigurationProperties{
				MetricConfigurations: []*armdataboxedge.MetricConfiguration{
					{
						CounterSets: []*armdataboxedge.MetricCounterSet{
							{
								Counters: []*armdataboxedge.MetricCounter{
									{
										Name: to.StringPtr("<name>"),
									}},
							}},
						MdmAccount:      to.StringPtr("<mdm-account>"),
						MetricNameSpace: to.StringPtr("<metric-name-space>"),
						ResourceID:      to.StringPtr("<resource-id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MonitoringConfigClientCreateOrUpdateResult)
}
