package armresourcegraph_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesHistoryGet.json
func ExampleClient_ResourcesHistory_resourceHistoryQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcegraph.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ResourcesHistory(ctx, armresourcegraph.ResourcesHistoryRequest{
		Options: &armresourcegraph.ResourcesHistoryRequestOptions{
			Interval: &armresourcegraph.DateTimeInterval{
				End:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T01:25:00.000Z"); return t }()),
				Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T01:00:00.000Z"); return t }()),
			},
		},
		Query: to.Ptr("where name =~ 'cpu-utilization' | project id, name, properties"),
		Subscriptions: []*string{
			to.Ptr("a7f33fdb-e646-4f15-89aa-3a360210861e")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Interface = map[string]any{
	// 	"count": float64(2),
	// 	"snapshots":map[string]any{
	// 		"columns":[]any{
	// 			map[string]any{
	// 				"name": "id",
	// 				"type": "string",
	// 			},
	// 			map[string]any{
	// 				"name": "name",
	// 				"type": "string",
	// 			},
	// 			map[string]any{
	// 				"name": "properties",
	// 				"type": "object",
	// 			},
	// 		},
	// 		"rows":[]any{
	// 			[]any{
	// 				"/subscriptions/a7f33fdb-e646-4f15-89aa-3a360210861e/resourceGroups/meya-test-rg/providers/Microsoft.Compute/virtualMachines/meya-win-eus/providers/Microsoft.WorkloadMonitor/monitors/cpu-utilization",
	// 				"cpu-utilization",
	// 				map[string]any{
	// 					"currentStateFirstObservedTimestamp": "",
	// 					"monitorName": "",
	// 				},
	// 			},
	// 			[]any{
	// 				"/subscriptions/a7f33fdb-e646-4f15-89aa-3a360210861e/resourceGroups/meya-test-rg/providers/Microsoft.Compute/virtualMachines/meya-win-eus/providers/Microsoft.WorkloadMonitor/monitors/cpu-utilization",
	// 				"cpu-utilization",
	// 				map[string]any{
	// 					"currentStateFirstObservedTimestamp": "",
	// 					"monitorName": "test",
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
