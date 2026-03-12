package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v3"
)

// Generated from example definition: 2025-08-01-preview/Gates_ListByFleet_MaximumSet_Gen.json
func ExampleGatesClient_NewListByFleetPager_gatesListByFleetMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("A5DFED4F-5511-4753-B6C8-3ACC54B370E3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGatesClient().NewListByFleetPager("rgfleets", "fleet-1", nil)
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
		// page = armcontainerservicefleet.GatesClientListByFleetResponse{
		// 	GateListResult: armcontainerservicefleet.GateListResult{
		// 		Value: []*armcontainerservicefleet.Gate{
		// 			{
		// 				Properties: &armcontainerservicefleet.GateProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservicefleet.GateProvisioningStateSucceeded),
		// 					DisplayName: to.Ptr("ohoncqhjnqlnpgqhidtmlwqitzujqvyvidoibym"),
		// 					GateType: to.Ptr(armcontainerservicefleet.GateTypeApproval),
		// 					Target: &armcontainerservicefleet.GateTarget{
		// 						ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1/gates/12345678-910a-bcde-f000-000000000000"),
		// 						UpdateRunProperties: &armcontainerservicefleet.UpdateRunGateTargetProperties{
		// 							Name: to.Ptr("run1"),
		// 							Stage: to.Ptr("stage1"),
		// 							Group: to.Ptr("group1"),
		// 							Timing: to.Ptr(armcontainerservicefleet.TimingAfter),
		// 						},
		// 					},
		// 					State: to.Ptr(armcontainerservicefleet.GateStatePending),
		// 				},
		// 				ETag: to.Ptr("zk"),
		// 				ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1/updateRuns/run1"),
		// 				Name: to.Ptr("tvpzzhsokcrigcrjnqyjhhdmsrm"),
		// 				Type: to.Ptr("xttejjiwuk"),
		// 				SystemData: &armcontainerservicefleet.SystemData{
		// 					CreatedBy: to.Ptr("afnyxgutytkmyjqwvebhdh"),
		// 					CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-21T20:36:49.685Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("ecpoq"),
		// 					LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-21T20:36:49.685Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
