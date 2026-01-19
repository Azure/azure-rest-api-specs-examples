package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: 2025-04-15-preview/ScheduledActions_VirtualMachinesGetOperationErrors_MaximumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesGetOperationErrors_scheduledActionsVirtualMachinesGetOperationErrorsMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("0505D8E4-D41A-48FB-9CA5-4AF8D93BE75F", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesGetOperationErrors(ctx, "ennweqswbghorrgzbet", armcomputeschedule.GetOperationErrorsContent{
		OperationIDs: []*string{
			to.Ptr("ksufjznokhsbowdupyt"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientVirtualMachinesGetOperationErrorsResponse{
	// 	GetOperationErrorsResponse: &armcomputeschedule.GetOperationErrorsResponse{
	// 		Results: []*armcomputeschedule.OperationErrorsResult{
	// 			{
	// 				OperationID: to.Ptr("emftjglfbsxaboxqzxlpbjian"),
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:04.130Z"); return t}()),
	// 				ActivationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:04.130Z"); return t}()),
	// 				CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:04.130Z"); return t}()),
	// 				OperationErrors: []*armcomputeschedule.OperationErrorDetails{
	// 					{
	// 						ErrorCode: to.Ptr("awrovoihpnqsotznapyrrb"),
	// 						ErrorDetails: to.Ptr("af"),
	// 						Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:04.131Z"); return t}()),
	// 						TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-15T19:47:04.131Z"); return t}()),
	// 						AzureOperationName: to.Ptr("ipsoybbslitmwsfkygfjhb"),
	// 						CrpOperationID: to.Ptr("cqkmbfb"),
	// 					},
	// 				},
	// 				RequestErrorCode: to.Ptr("vzixjphgeygyhzpbbkgzcjol"),
	// 				RequestErrorDetails: to.Ptr("ficjafazcvbmlbnqhffwtevkla"),
	// 			},
	// 		},
	// 	},
	// }
}
