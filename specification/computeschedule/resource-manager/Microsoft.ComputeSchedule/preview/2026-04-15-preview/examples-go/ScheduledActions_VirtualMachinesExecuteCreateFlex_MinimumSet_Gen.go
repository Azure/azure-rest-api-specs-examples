package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/ScheduledActions_VirtualMachinesExecuteCreateFlex_MinimumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesExecuteCreateFlex_scheduledActionsVirtualMachinesExecuteCreateFlexMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("732116BD-AF31-4E74-9283-B387C44B4A44", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesExecuteCreateFlex(ctx, "eastus2", armcomputeschedule.ExecuteCreateFlexRequest{
		ResourceConfigParameters: &armcomputeschedule.ResourceProvisionFlexPayload{
			ResourceCount: to.Ptr[int32](24),
			FlexProperties: &armcomputeschedule.FlexProperties{
				VMSizeProfiles: []*armcomputeschedule.VMSizeProfile{
					{
						Name: to.Ptr("Standard_D2s_v3"),
						Rank: to.Ptr[int32](24),
					},
					{
						Name: to.Ptr("Standard_D2s_v3"),
						Rank: to.Ptr[int32](24),
					},
				},
				OSType:          to.Ptr(armcomputeschedule.OsTypeWindows),
				PriorityProfile: &armcomputeschedule.PriorityProfile{},
			},
		},
		ExecutionParameters: &armcomputeschedule.ExecutionParameters{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientVirtualMachinesExecuteCreateFlexResponse{
	// 	CreateFlexResourceOperationResponse: armcomputeschedule.CreateFlexResourceOperationResponse{
	// 		Type: to.Ptr("VirtualMachine"),
	// 		Location: to.Ptr("eastus2"),
	// 		Description: to.Ptr("Flex create operation completed successfully"),
	// 	},
	// }
}
