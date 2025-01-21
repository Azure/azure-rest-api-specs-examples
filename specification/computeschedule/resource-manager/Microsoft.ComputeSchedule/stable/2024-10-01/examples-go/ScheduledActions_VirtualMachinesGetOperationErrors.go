package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a4eca6b060cf70da696963245656fdc440b666b/specification/computeschedule/resource-manager/Microsoft.ComputeSchedule/stable/2024-10-01/examples/ScheduledActions_VirtualMachinesGetOperationErrors.json
func ExampleScheduledActionsClient_VirtualMachinesGetOperationErrors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesGetOperationErrors(ctx, "eastus2euap", armcomputeschedule.GetOperationErrorsRequest{
		OperationIDs: []*string{
			to.Ptr("23480d2f-1dca-4610-afb4-dd25eec1f34r")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GetOperationErrorsResponse = armcomputeschedule.GetOperationErrorsResponse{
	// 	Results: []*armcomputeschedule.OperationErrorsResult{
	// 		{
	// 			OperationID: to.Ptr("23480d2f-1dca-4610-afb4-dd25eec1f34r"),
	// 			ActivationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T17:52:53.999Z"); return t}()),
	// 			CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T17:52:53.999Z"); return t}()),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T17:52:53.999Z"); return t}()),
	// 			OperationErrors: []*armcomputeschedule.OperationErrorDetails{
	// 				{
	// 					AzureOperationName: to.Ptr("23480d2f-1dca-4610-afb4-dd25eec1f34r"),
	// 					CrpOperationID: to.Ptr("23480d2f-1dca-4610-afb4-dd25eec1f34r"),
	// 					ErrorCode: to.Ptr("null"),
	// 					ErrorDetails: to.Ptr("null"),
	// 					TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-27T16:55:03.357Z"); return t}()),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T17:52:53.999Z"); return t}()),
	// 			}},
	// 			RequestErrorCode: to.Ptr("null"),
	// 			RequestErrorDetails: to.Ptr("null"),
	// 	}},
	// }
}
