package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/AttachedNetworks_GetByDevCenter.json
func ExampleAttachedNetworksClient_GetByDevCenter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAttachedNetworksClient().GetByDevCenter(ctx, "rg1", "Contoso", "network-uswest3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.AttachedNetworksClientGetByDevCenterResponse{
	// 	AttachedNetworkConnection: armdevcenter.AttachedNetworkConnection{
	// 		Name: to.Ptr("network-uswest3"),
	// 		Type: to.Ptr("Microsoft.DevCenter/devcenters/attachednetworks"),
	// 		ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/attachednetworks/network-uswest3"),
	// 		Properties: &armdevcenter.AttachedNetworkConnectionProperties{
	// 			HealthCheckStatus: to.Ptr(armdevcenter.HealthCheckStatus("Healthy")),
	// 			NetworkConnectionID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/NetworkConnections/network-uswest3"),
	// 			NetworkConnectionLocation: to.Ptr("centralus"),
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateCreated),
	// 		},
	// 		SystemData: &armdevcenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 			CreatedBy: to.Ptr("User1"),
	// 			CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("User1"),
	// 			LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
