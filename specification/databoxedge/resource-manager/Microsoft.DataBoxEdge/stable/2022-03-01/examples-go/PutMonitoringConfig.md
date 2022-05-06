Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataboxedge%2Farmdataboxedge%2Fv0.4.0/sdk/resourcemanager/databoxedge/armdataboxedge/README.md) on how to add the SDK to your project and authenticate.

```go
package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/PutMonitoringConfig.json
func ExampleMonitoringConfigClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdataboxedge.NewMonitoringConfigClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
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
										Name: to.Ptr("<name>"),
									}},
							}},
						MdmAccount:      to.Ptr("<mdm-account>"),
						MetricNameSpace: to.Ptr("<metric-name-space>"),
						ResourceID:      to.Ptr("<resource-id>"),
					}},
			},
		},
		&armdataboxedge.MonitoringConfigClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
