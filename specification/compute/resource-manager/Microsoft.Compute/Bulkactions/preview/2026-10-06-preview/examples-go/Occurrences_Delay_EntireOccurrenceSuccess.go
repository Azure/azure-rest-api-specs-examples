package armbulkactions_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/Occurrences_Delay_EntireOccurrenceSuccess.json
func ExampleOccurrencesClient_BeginDelay_twoDelayAllOperationsInARecurringScheduledActionOccurrence() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOccurrencesClient().BeginDelay(ctx, "example-rg", "weekday-start", "77777777-7777-7777-7777-777777777777", armbulkactions.DelayRequest{
		Delay:       to.Ptr(time.Date(2026, time.September, 15, 9, 0, 0, 0, time.FixedZone("", -25200))),
		ResourceIDs: []*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.OccurrencesClientDelayResponse{
	// 	ResourceOperationResponse: armbulkactions.ResourceOperationResponse{
	// 		TotalResources: to.Ptr[int32](2),
	// 		ResourcesStatuses: []*armbulkactions.ResourceStatus{
	// 		},
	// 	},
	// }
}
