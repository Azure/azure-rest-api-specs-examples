package armworkloadmonitor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadmonitor/armworkloadmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/workloadmonitor/resource-manager/Microsoft.WorkloadMonitor/preview/2020-01-13-preview/examples/MonitorHistory_GetDefault.json
func ExampleHealthMonitorsClient_NewListStateChangesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armworkloadmonitor.NewHealthMonitorsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListStateChangesPager("bc27da3b-3ba2-4e00-a6ec-1fde64aa1e21",
		"tugamidiAlerts",
		"Microsoft.Compute",
		"virtualMachines",
		"linuxEUS",
		"logical-disks|C@3A",
		&armworkloadmonitor.HealthMonitorsClientListStateChangesOptions{Filter: nil,
			Expand:            nil,
			StartTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-19T19:24:14Z"); return t }()),
			EndTimestampUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-20T01:24:14Z"); return t }()),
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
