package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/AttachedNetworks_ListByDevCenter.json
func ExampleAttachedNetworksClient_NewListByDevCenterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAttachedNetworksClient().NewListByDevCenterPager("rg1", "Contoso", nil)
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
		// page = armdevcenter.AttachedNetworksClientListByDevCenterResponse{
		// 	AttachedNetworkListResult: armdevcenter.AttachedNetworkListResult{
		// 		Value: []*armdevcenter.AttachedNetworkConnection{
		// 			{
		// 				Name: to.Ptr("netmap1"),
		// 				Type: to.Ptr("Microsoft.DevCenter/devcenters/attachedNetworks"),
		// 				ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/attachednetworks/netmap1"),
		// 				Properties: &armdevcenter.AttachedNetworkConnectionProperties{
		// 					HealthCheckStatus: to.Ptr(armdevcenter.HealthCheckStatus("Healthy")),
		// 					NetworkConnectionID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/NetworkConnections/network-uswest3"),
		// 					NetworkConnectionLocation: to.Ptr("centralus"),
		// 					ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
		// 				},
		// 				SystemData: &armdevcenter.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
		// 					CreatedBy: to.Ptr("User1"),
		// 					CreatedByType: to.Ptr(armdevcenter.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("User1"),
		// 					LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeApplication),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("netmap2"),
		// 				Type: to.Ptr("Microsoft.DevCenter/devcenters/attachedNetworks"),
		// 				ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/attachednetworks/netmap2"),
		// 				Properties: &armdevcenter.AttachedNetworkConnectionProperties{
		// 					HealthCheckStatus: to.Ptr(armdevcenter.HealthCheckStatus("Healthy")),
		// 					NetworkConnectionID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/networkconnections/network-uswest3"),
		// 					NetworkConnectionLocation: to.Ptr("centralus"),
		// 					ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
		// 				},
		// 				SystemData: &armdevcenter.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
		// 					CreatedBy: to.Ptr("User1"),
		// 					CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("User1"),
		// 					LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
