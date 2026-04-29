package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineExamples/VirtualMachine_AssessPatches.json
func ExampleVirtualMachinesClient_BeginAssessPatches() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginAssessPatches(ctx, "myResourceGroupName", "myVMName", nil)
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
	// res = armcompute.VirtualMachinesClientAssessPatchesResponse{
	// 	VirtualMachineAssessPatchesResult: &armcompute.VirtualMachineAssessPatchesResult{
	// 		Status: to.Ptr(armcompute.PatchOperationStatusSucceeded),
	// 		AssessmentActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 		RebootPending: to.Ptr(true),
	// 		CriticalAndSecurityPatchCount: to.Ptr[int32](1),
	// 		OtherPatchCount: to.Ptr[int32](2),
	// 		StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T21:02:04.2556154Z"); return t}()),
	// 		AvailablePatches: []*armcompute.VirtualMachineSoftwarePatchProperties{
	// 			{
	// 				PatchID: to.Ptr("35428702-5784-4ba4-a6e0-5222258b5411"),
	// 				Name: to.Ptr("Definition Update for Windows Defender Antivirus - KB2267602 (Definition 1.279.1373.0)"),
	// 				Version: to.Ptr(""),
	// 				KbID: to.Ptr("2267602"),
	// 				Classifications: []*string{
	// 					to.Ptr("Definition Updates"),
	// 				},
	// 				RebootBehavior: to.Ptr(armcompute.VMGuestPatchRebootBehaviorNeverReboots),
	// 				ActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 				PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-07T00:00:00Z"); return t}()),
	// 				LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T21:18:45.2830263Z"); return t}()),
	// 				AssessmentState: to.Ptr(armcompute.PatchAssessmentStateAvailable),
	// 			},
	// 			{
	// 				PatchID: to.Ptr("39f9cdd1-795c-4d0e-8c0a-73ab3f31746d"),
	// 				Name: to.Ptr("Windows Malicious Software Removal Tool x64 - October 2018 (KB890830)"),
	// 				Version: to.Ptr(""),
	// 				KbID: to.Ptr("890830"),
	// 				Classifications: []*string{
	// 					to.Ptr("Update Rollups"),
	// 				},
	// 				RebootBehavior: to.Ptr(armcompute.VMGuestPatchRebootBehaviorCanRequestReboot),
	// 				ActivityID: to.Ptr("68f8b292-dfc2-4646-9781-33cc88631968"),
	// 				PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-07T00:00:00Z"); return t}()),
	// 				LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T21:18:45.2830263Z"); return t}()),
	// 				AssessmentState: to.Ptr(armcompute.PatchAssessmentStateAvailable),
	// 			},
	// 		},
	// 	},
	// }
}
