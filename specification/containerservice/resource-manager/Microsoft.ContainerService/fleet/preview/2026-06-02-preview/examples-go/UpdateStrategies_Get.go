package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v3"
)

// Generated from example definition: 2026-06-02-preview/UpdateStrategies_Get.json
func ExampleFleetUpdateStrategiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetUpdateStrategiesClient().Get(ctx, "rg1", "fleet1", "strategy1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservicefleet.FleetUpdateStrategiesClientGetResponse{
	// 	FleetUpdateStrategy: armcontainerservicefleet.FleetUpdateStrategy{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateStrategies/strategy1"),
	// 		Name: to.Ptr("strategy1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/fleets/updateStrategies"),
	// 		SystemData: &armcontainerservicefleet.SystemData{
	// 			CreatedBy: to.Ptr("@contoso.com"),
	// 			CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2023, time.March, 1, 1, 10, 8, 395000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2023, time.March, 1, 1, 10, 8, 395000000, time.UTC)),
	// 		},
	// 		Properties: &armcontainerservicefleet.FleetUpdateStrategyProperties{
	// 			ProvisioningState: to.Ptr(armcontainerservicefleet.FleetUpdateStrategyProvisioningStateSucceeded),
	// 			Strategy: &armcontainerservicefleet.UpdateRunStrategy{
	// 				Stages: []*armcontainerservicefleet.UpdateStage{
	// 					{
	// 						Name: to.Ptr("stage1"),
	// 						MaxConcurrency: to.Ptr("20%"),
	// 						Groups: []*armcontainerservicefleet.UpdateGroup{
	// 							{
	// 								Name: to.Ptr("group-a"),
	// 								MaxConcurrency: to.Ptr("5"),
	// 								MemberSelector: &armcontainerservicefleet.MemberSelector{
	// 									ByLabel: to.Ptr("tier=frontend"),
	// 								},
	// 								BeforeGates: []*armcontainerservicefleet.GateConfiguration{
	// 									{
	// 										DisplayName: to.Ptr("gate before group-a"),
	// 										Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
	// 									},
	// 								},
	// 								AfterGates: []*armcontainerservicefleet.GateConfiguration{
	// 									{
	// 										DisplayName: to.Ptr("gate after group-a"),
	// 										Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						BeforeGates: []*armcontainerservicefleet.GateConfiguration{
	// 							{
	// 								DisplayName: to.Ptr("gate before stage1"),
	// 								Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
	// 							},
	// 						},
	// 						AfterGates: []*armcontainerservicefleet.GateConfiguration{
	// 							{
	// 								DisplayName: to.Ptr("gate after stage1"),
	// 								Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
	// 							},
	// 						},
	// 						AfterStageWaitInSeconds: to.Ptr[int32](3600),
	// 					},
	// 					{
	// 						Name: to.Ptr("stage2"),
	// 						MaxConcurrency: to.Ptr("50%"),
	// 						MemberSelector: &armcontainerservicefleet.MemberSelector{
	// 							ByLabel: to.Ptr("env=production"),
	// 						},
	// 						AfterStageWaitInSeconds: to.Ptr[int32](600),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ETag: to.Ptr("\"EtagValue\""),
	// 	},
	// }
}
