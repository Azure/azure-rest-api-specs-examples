package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v3"
)

// Generated from example definition: 2025-04-01-preview/UpdateStrategies_ListByFleet.json
func ExampleFleetUpdateStrategiesClient_NewListByFleetPager_listTheFleetUpdateStrategyResourcesByFleet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetUpdateStrategiesClient().NewListByFleetPager("rg1", "fleet1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcontainerservicefleet.FleetUpdateStrategiesClientListByFleetResponse{
		// 	FleetUpdateStrategyListResult: armcontainerservicefleet.FleetUpdateStrategyListResult{
		// 		Value: []*armcontainerservicefleet.FleetUpdateStrategy{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateStrategies/strategy1"),
		// 				Name: to.Ptr("strategy1"),
		// 				Type: to.Ptr("Microsoft.ContainerService/fleets/updateStrategies"),
		// 				SystemData: &armcontainerservicefleet.SystemData{
		// 					CreatedBy: to.Ptr("@contoso.com"),
		// 					CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
		// 				},
		// 				Properties: &armcontainerservicefleet.FleetUpdateStrategyProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservicefleet.FleetUpdateStrategyProvisioningStateSucceeded),
		// 					Strategy: &armcontainerservicefleet.UpdateRunStrategy{
		// 						Stages: []*armcontainerservicefleet.UpdateStage{
		// 							{
		// 								Name: to.Ptr("stage1"),
		// 								Groups: []*armcontainerservicefleet.UpdateGroup{
		// 									{
		// 										Name: to.Ptr("group-a"),
		// 										BeforeGates: []*armcontainerservicefleet.GateConfiguration{
		// 											{
		// 												DisplayName: to.Ptr("gate before group-a"),
		// 												Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
		// 											},
		// 										},
		// 										AfterGates: []*armcontainerservicefleet.GateConfiguration{
		// 											{
		// 												DisplayName: to.Ptr("gate after group-a"),
		// 												Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
		// 											},
		// 										},
		// 									},
		// 								},
		// 								BeforeGates: []*armcontainerservicefleet.GateConfiguration{
		// 									{
		// 										DisplayName: to.Ptr("gate before stage1"),
		// 										Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
		// 									},
		// 								},
		// 								AfterGates: []*armcontainerservicefleet.GateConfiguration{
		// 									{
		// 										DisplayName: to.Ptr("gate after stage1"),
		// 										Type: to.Ptr(armcontainerservicefleet.GateTypeApproval),
		// 									},
		// 								},
		// 								AfterStageWaitInSeconds: to.Ptr[int32](3600),
		// 							},
		// 						},
		// 					},
		// 				},
		// 				ETag: to.Ptr("\"EtagValue\""),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://nextlink.contoso.com"),
		// 	},
		// }
	}
}
