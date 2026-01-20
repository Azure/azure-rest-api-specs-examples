package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: 2025-04-15-preview/Occurrences_Get_MaximumSet_Gen.json
func ExampleOccurrencesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOccurrencesClient().Get(ctx, "rgcomputeschedule", "myScheduledAction", "67b5bada-4772-43fc-8dbb-402476d98a45", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.OccurrencesClientGetResponse{
	// 	Occurrence: &armcomputeschedule.Occurrence{
	// 		Properties: &armcomputeschedule.OccurrenceProperties{
	// 			ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:59.243Z"); return t}()),
	// 			ResultSummary: &armcomputeschedule.OccurrenceResultSummary{
	// 				Total: to.Ptr[int32](25),
	// 				Statuses: []*armcomputeschedule.ResourceResultSummary{
	// 					{
	// 						Code: to.Ptr("ubmwzgcwfopegjtkqyrkedgish"),
	// 						Count: to.Ptr[int32](4),
	// 						ErrorDetails: &armcomputeschedule.Error{
	// 							Code: to.Ptr("baxjmkbhoatqcj"),
	// 							Message: to.Ptr("chapcwfkqymeof"),
	// 							Target: to.Ptr("mkirmorowetsigohjamvk"),
	// 							Details: []*armcomputeschedule.Error{
	// 							},
	// 							Innererror: &armcomputeschedule.InnerError{
	// 								Code: to.Ptr("cgalioufsabcwatbxa"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcomputeschedule.OccurrenceStateCreated),
	// 		},
	// 		ID: to.Ptr("/subscriptions/83C27AB3-A7B9-498B-B165-D9440661474F/resourceGroups/myRg/providers/Microsoft.ComputeSchedule/scheduledActions/myScheduledAction/occurrences/83C27AB3-A7B9-498B-B165-D9440661474F"),
	// 		Name: to.Ptr("edirlrdovp"),
	// 		Type: to.Ptr("hvnnadgtjavnsbilpaipgrtdjy"),
	// 		SystemData: &armcomputeschedule.SystemData{
	// 			CreatedBy: to.Ptr("cvryvreuvvjtiamcwhisrt"),
	// 			CreatedByType: to.Ptr(armcomputeschedule.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:55.288Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("supbnksztdbgulxgvfmqvriqdlpirh"),
	// 			LastModifiedByType: to.Ptr(armcomputeschedule.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-17T00:23:55.288Z"); return t}()),
	// 		},
	// 	},
	// }
}
