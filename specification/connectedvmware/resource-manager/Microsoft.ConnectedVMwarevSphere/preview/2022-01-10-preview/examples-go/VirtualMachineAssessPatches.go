package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/VirtualMachineAssessPatches.json
func ExampleVirtualMachinesClient_BeginAssessPatches() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginAssessPatches(ctx, "myResourceGroupName", "myMachineName", nil)
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
	// res.VirtualMachineAssessPatchesResult = armconnectedvmware.VirtualMachineAssessPatchesResult{
	// 	AssessmentActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 	AvailablePatchCountByClassification: &armconnectedvmware.AvailablePatchCountByClassification{
	// 		Critical: to.Ptr[int32](0),
	// 		Definition: to.Ptr[int32](0),
	// 		FeaturePack: to.Ptr[int32](0),
	// 		Security: to.Ptr[int32](0),
	// 		ServicePack: to.Ptr[int32](0),
	// 		Tools: to.Ptr[int32](0),
	// 		UpdateRollup: to.Ptr[int32](1),
	// 		Updates: to.Ptr[int32](1),
	// 	},
	// 	LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-15T02:16:06.9740000Z"); return t}()),
	// 	OSType: to.Ptr(armconnectedvmware.OsTypeUMWindows),
	// 	RebootPending: to.Ptr(true),
	// 	StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-15T02:15:20.9340000Z"); return t}()),
	// 	StartedBy: to.Ptr(armconnectedvmware.PatchOperationStartedByUser),
	// 	Status: to.Ptr(armconnectedvmware.PatchOperationStatusSucceeded),
	// }
}
