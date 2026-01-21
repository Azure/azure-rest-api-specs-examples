package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_GetInstanceView_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_GetInstanceView_virtualMachineScaleSetGetInstanceViewMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetsClient().GetInstanceView(ctx, "rgcompute", "aaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineScaleSetInstanceView = armcompute.VirtualMachineScaleSetInstanceView{
	// 	OrchestrationServices: []*armcompute.OrchestrationServiceSummary{
	// 		{
	// 			ServiceName: to.Ptr(armcompute.OrchestrationServiceNamesAutomaticRepairs),
	// 			ServiceState: to.Ptr(armcompute.OrchestrationServiceStateNotRunning),
	// 		},
	// 		{
	// 			LastStatusChangeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-09T13:26:28.360Z"); return t}()),
	// 			LatestOperationStatus: to.Ptr(armcompute.OrchestrationServiceOperationStatusInProgress),
	// 			ServiceName: to.Ptr(armcompute.OrchestrationServiceNamesAutomaticZoneRebalancing),
	// 			ServiceState: to.Ptr(armcompute.OrchestrationServiceStateRunning),
	// 	}},
	// 	Statuses: []*armcompute.InstanceViewStatus{
	// 		{
	// 			Code: to.Ptr("ProvisioningState/succeeded"),
	// 			DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 			Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			Message: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 			Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.526Z"); return t}()),
	// 	}},
	// 	VirtualMachine: &armcompute.VirtualMachineScaleSetInstanceViewStatusesSummary{
	// 		StatusesSummary: []*armcompute.VirtualMachineStatusCodeCount{
	// 			{
	// 				Code: to.Ptr("aa"),
	// 				Count: to.Ptr[int32](21),
	// 		}},
	// 	},
	// 	Extensions: []*armcompute.VirtualMachineScaleSetVMExtensionsSummary{
	// 		{
	// 			Name: to.Ptr("aaaaaaaaaaa"),
	// 			StatusesSummary: []*armcompute.VirtualMachineStatusCodeCount{
	// 				{
	// 					Code: to.Ptr("aa"),
	// 					Count: to.Ptr[int32](21),
	// 			}},
	// 	}},
	// }
}
