package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace/v2"
)

// Generated from example definition: 2024-04-24/Monitors_CreateOrUpdate_MinimumSet_Gen.json
func ExampleMonitorsClient_BeginCreateOrUpdate_monitorsCreateOrUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitorsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myMonitor", armdynatrace.MonitorResource{
		Location:   to.Ptr("West US 2"),
		Properties: &armdynatrace.MonitorProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdynatrace.MonitorsClientCreateOrUpdateResponse{
	// 	MonitorResource: &armdynatrace.MonitorResource{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Dynatrace.Observability/monitors/myMonitor"),
	// 		Location: to.Ptr("West US 2"),
	// 		Properties: &armdynatrace.MonitorProperties{
	// 		},
	// 	},
	// }
}
