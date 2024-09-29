package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/computeschedule/resource-manager/Microsoft.ComputeSchedule/preview/2024-08-15-preview/examples/ScheduledActions_VirtualMachinesGetOperationErrors_MinimumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesGetOperationErrors_csScheduledActionsVirtualMachinesGetOperationErrorsMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesGetOperationErrors(ctx, "ggxoaxzxtdbi", armcomputeschedule.GetOperationErrorsRequest{
		OperationIDs: []*string{
			to.Ptr("qeicik")},
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
	// 			OperationID: to.Ptr("wetjrhx"),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-08T19:00:14.771Z"); return t}()),
	// 	}},
	// }
}
