package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/machine/Machine_AssessPatches.json
func ExampleMachinesClient_BeginAssessPatches() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMachinesClient().BeginAssessPatches(ctx, "myResourceGroupName", "myMachineName", nil)
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
	// res.MachineAssessPatchesResult = armhybridcompute.MachineAssessPatchesResult{
	// 	AssessmentActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 	AvailablePatchCountByClassification: &armhybridcompute.AvailablePatchCountByClassification{
	// 		Critical: to.Ptr[int32](0),
	// 		Definition: to.Ptr[int32](0),
	// 		FeaturePack: to.Ptr[int32](0),
	// 		Security: to.Ptr[int32](0),
	// 		ServicePack: to.Ptr[int32](0),
	// 		Tools: to.Ptr[int32](0),
	// 		UpdateRollup: to.Ptr[int32](1),
	// 		Updates: to.Ptr[int32](1),
	// 	},
	// 	LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-22T02:16:06.974Z"); return t}()),
	// 	OSType: to.Ptr(armhybridcompute.OsTypeWindows),
	// 	RebootPending: to.Ptr(true),
	// 	StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-22T02:15:20.934Z"); return t}()),
	// 	StartedBy: to.Ptr(armhybridcompute.PatchOperationStartedByUser),
	// 	Status: to.Ptr(armhybridcompute.PatchOperationStatusSucceeded),
	// }
}
