package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/examples/UpdateStrategies_Get.json
func ExampleFleetUpdateStrategiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.FleetUpdateStrategy = armcontainerservicefleet.FleetUpdateStrategy{
	// 	Name: to.Ptr("strategy1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/fleets/updateStrategies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/updateStrategies/strategy1"),
	// 	SystemData: &armcontainerservicefleet.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
	// 		CreatedBy: to.Ptr("@contoso.com"),
	// 		CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 	},
	// 	ETag: to.Ptr("\"EtagValue\""),
	// 	Properties: &armcontainerservicefleet.FleetUpdateStrategyProperties{
	// 		ProvisioningState: to.Ptr(armcontainerservicefleet.FleetUpdateStrategyProvisioningStateSucceeded),
	// 		Strategy: &armcontainerservicefleet.UpdateRunStrategy{
	// 			Stages: []*armcontainerservicefleet.UpdateStage{
	// 				{
	// 					Name: to.Ptr("stage1"),
	// 					AfterStageWaitInSeconds: to.Ptr[int32](3600),
	// 					Groups: []*armcontainerservicefleet.UpdateGroup{
	// 						{
	// 							Name: to.Ptr("group-a"),
	// 					}},
	// 			}},
	// 		},
	// 	},
	// }
}
