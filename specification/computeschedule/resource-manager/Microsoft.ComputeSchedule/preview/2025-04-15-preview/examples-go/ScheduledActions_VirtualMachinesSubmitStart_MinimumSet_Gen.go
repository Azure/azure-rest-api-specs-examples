package armcomputeschedule_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: 2025-04-15-preview/ScheduledActions_VirtualMachinesSubmitStart_MinimumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesSubmitStart_scheduledActionsVirtualMachinesSubmitStartMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesSubmitStart(ctx, "ufrcsuw", armcomputeschedule.SubmitStartContent{
		Schedule: &armcomputeschedule.Schedule{
			DeadLine:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:56.803Z"); return t }()),
			TimeZone:     to.Ptr("aigbjdnldtzkteqi"),
			DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeUnknown),
		},
		ExecutionParameters: &armcomputeschedule.ExecutionParameters{},
		Resources: &armcomputeschedule.Resources{
			IDs: []*string{
				to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4"),
			},
		},
		Correlationid: to.Ptr("b211f086-4b91-4686-a453-2f5c012e4d80"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientVirtualMachinesSubmitStartResponse{
	// 	StartResourceOperationResponse: &armcomputeschedule.StartResourceOperationResponse{
	// 		Description: to.Ptr("tlphodyrecv"),
	// 		Type: to.Ptr("qpmru"),
	// 		Location: to.Ptr("ktsumrgdaifwbpkxurfdfa"),
	// 	},
	// }
}
